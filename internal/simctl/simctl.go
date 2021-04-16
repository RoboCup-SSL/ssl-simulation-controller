package simctl

import (
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/sslnet"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/tracker"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/vision"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
)

type SimulationController struct {
	visionServer     *sslnet.MulticastServer
	refereeServer    *sslnet.MulticastServer
	trackerServer    *sslnet.MulticastServer
	simControlClient *sslnet.UdpClient

	simControlPort string

	lastTrackedFrame *tracker.TrackerWrapperPacket
	lastRefereeMsg   *referee.Referee
	fieldSize        *vision.SSL_GeometryFieldSize

	ballReplacer         BallReplacer
	robotCountMaintainer RobotCountMaintainer
	robotSpecsSetter     RobotSpecSetter
}

func NewSimulationController(visionAddress, refereeAddress, trackerAddress, simControlPort, robotSpecConfig string) (c *SimulationController) {
	c = new(SimulationController)
	c.visionServer = sslnet.NewMulticastServer(visionAddress, c.onNewVisionData)
	c.refereeServer = sslnet.NewMulticastServer(refereeAddress, c.onNewRefereeData)
	c.trackerServer = sslnet.NewMulticastServer(trackerAddress, c.onNewTrackerData)
	c.simControlPort = simControlPort
	c.ballReplacer.c = c
	c.robotCountMaintainer.c = c
	c.robotSpecsSetter.c = c
	c.robotSpecsSetter.LoadRobotSpecs(robotSpecConfig)
	return
}

func (c *SimulationController) onNewVisionData(data []byte, remoteAddr *net.UDPAddr) {
	wrapper := vision.SSL_WrapperPacket{}
	if err := proto.Unmarshal(data, &wrapper); err != nil {
		log.Println("Could not unmarshal vision wrapper packet", err)
		return
	}

	if wrapper.Geometry != nil && wrapper.Geometry.Field != nil {
		c.fieldSize = wrapper.Geometry.Field
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
	c.lastRefereeMsg = refereeMsg
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
		c.lastTrackedFrame = frame
		c.handle()
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

	c.ballReplacer.handleReplaceBall()
	c.robotCountMaintainer.handleRobotCount()
	c.robotSpecsSetter.handleRobotSpecs()
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
