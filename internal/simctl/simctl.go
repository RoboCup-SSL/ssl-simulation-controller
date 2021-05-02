package simctl

import (
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/sslnet"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/tracker"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/vision"
	"github.com/golang/protobuf/proto"
	"log"
	"math"
	"net"
	"sync"
)

type SimulationController struct {
	visionServer     *sslnet.MulticastServer
	refereeServer    *sslnet.MulticastServer
	trackerServer    *sslnet.MulticastServer
	simControlClient *sslnet.UdpClient
	mutex            sync.Mutex

	simControlPort     string
	simulatorRestarted bool

	lastTrackedFrame   *tracker.TrackerWrapperPacket
	lastRefereeMsg     *referee.Referee
	fieldSize          *vision.SSL_GeometryFieldSize
	lastVisionFrameIds map[uint32]uint32

	ballReplaceHandler *BallReplaceHandler
	robotCountHandler  *RobotCountHandler
	robotSpecsHandler  *RobotSpecHandler
	geometryHandler    *GeometryHandler
}

func NewSimulationController(visionAddress, refereeAddress, trackerAddress, simControlPort, robotSpecConfig string) (c *SimulationController) {
	c = new(SimulationController)
	c.visionServer = sslnet.NewMulticastServer(visionAddress, c.onNewVisionData)
	c.refereeServer = sslnet.NewMulticastServer(refereeAddress, c.onNewRefereeData)
	c.trackerServer = sslnet.NewMulticastServer(trackerAddress, c.onNewTrackerData)
	c.simControlPort = simControlPort
	c.simulatorRestarted = true
	c.lastVisionFrameIds = map[uint32]uint32{}

	c.ballReplaceHandler = NewBallReplaceHandler(c)
	c.robotCountHandler = NewRobotCountHandler(c)
	c.robotSpecsHandler = NewRobotSpecHandler(c, robotSpecConfig)
	c.geometryHandler = NewGeometryHandler(c)

	c.ballReplaceHandler.c = c
	c.robotCountHandler.c = c
	c.robotSpecsHandler.c = c
	c.geometryHandler.c = c
	return
}

func (c *SimulationController) onNewVisionData(data []byte, remoteAddr *net.UDPAddr) {
	wrapper := vision.SSL_WrapperPacket{}
	if err := proto.Unmarshal(data, &wrapper); err != nil {
		log.Println("Could not unmarshal vision wrapper packet", err)
		return
	}

	if wrapper.Geometry != nil && wrapper.Geometry.Field != nil {
		c.mutex.Lock()
		c.fieldSize = wrapper.Geometry.Field
		c.mutex.Unlock()
	}

	if wrapper.Detection != nil {
		c.mutex.Lock()
		frameId := *wrapper.Detection.FrameNumber
		lastFrameId, lastFrameIdPresent := c.lastVisionFrameIds[*wrapper.Detection.CameraId]
		if lastFrameIdPresent && math.Abs(float64(frameId-lastFrameId)) > 100 {
			// large frame id change: Simulator probably restarted
			c.simulatorRestarted = true
			c.lastVisionFrameIds = map[uint32]uint32{}
			log.Printf("Simulator restart detected due to high frame id change (%d -> %d)",
				lastFrameId, frameId)
		}
		c.lastVisionFrameIds[*wrapper.Detection.CameraId] = frameId
		c.mutex.Unlock()
	}

	if c.simControlClient == nil {
		address := remoteAddr.IP.String() + ":" + c.simControlPort
		c.simControlClient = sslnet.NewUdpClient(address)
		c.simControlClient.Consumer = c.onNewSimControlResponseData
		c.simControlClient.Start()
	}
}

func (c *SimulationController) onNewRefereeData(data []byte, _ *net.UDPAddr) {
	refereeMsg := new(referee.Referee)
	if err := proto.Unmarshal(data, refereeMsg); err != nil {
		log.Println("Could not unmarshal referee packet", err)
		return
	}
	c.mutex.Lock()
	c.lastRefereeMsg = refereeMsg
	c.mutex.Unlock()
}

func (c *SimulationController) onNewTrackerData(data []byte, _ *net.UDPAddr) {
	frame := new(tracker.TrackerWrapperPacket)
	if err := proto.Unmarshal(data, frame); err != nil {
		log.Println("Could not unmarshal tracker packet", err)
		return
	}
	if c.lastTrackedFrame == nil || // very first frame
		*c.lastTrackedFrame.Uuid == *frame.Uuid || // frame from same origin
		// new frame is significantly newer than last frame
		(*frame.TrackedFrame.Timestamp-*c.lastTrackedFrame.TrackedFrame.Timestamp) > 5 {
		c.mutex.Lock()
		c.lastTrackedFrame = frame
		c.handle()
		c.mutex.Unlock()
	}
}

func (c *SimulationController) onNewSimControlResponseData(data []byte) {
	response := new(SimulatorResponse)
	if err := proto.Unmarshal(data, response); err != nil {
		log.Println("Could not unmarshal tracker packet", err)
		return
	}
	for _, responseError := range response.Errors {
		log.Printf("SimControl Error: %v", responseError)
	}
}

func (c *SimulationController) handle() {
	if c.lastTrackedFrame == nil ||
		c.fieldSize == nil ||
		c.lastRefereeMsg == nil ||
		c.simControlClient == nil {
		return
	}

	if c.simulatorRestarted {
		c.ballReplaceHandler.Reset()
		c.robotCountHandler.Reset()
		c.robotSpecsHandler.Reset()
		c.geometryHandler.Reset()
		c.simulatorRestarted = false
	}

	c.ballReplaceHandler.handleReplaceBall()
	c.robotCountHandler.handleRobotCount()
	c.robotSpecsHandler.handleRobotSpecs()
	c.geometryHandler.handleGeometry()
}

func (c *SimulationController) Start() {
	c.visionServer.Start()
	c.refereeServer.Start()
	c.trackerServer.Start()
}

func (c *SimulationController) Stop() {
	c.visionServer.Stop()
	c.refereeServer.Stop()
	c.trackerServer.Stop()
	if c.simControlClient != nil {
		c.simControlClient.Stop()
		c.simControlClient = nil
	}
}
