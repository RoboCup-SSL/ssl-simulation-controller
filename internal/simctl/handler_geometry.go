package simctl

import (
	_ "embed"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/vision"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

//go:embed geometry-div-a.txt
var geometryConfigDivA string

//go:embed geometry-div-b.txt
var geometryConfigDivB string

type GeometryHandler struct {
	c *SimulationController

	geometryDivA *vision.SSL_GeometryData
	geometryDivB *vision.SSL_GeometryData
	timeLastSent *time.Time
}

func NewGeometryHandler(c *SimulationController) (r *GeometryHandler) {
	r = new(GeometryHandler)
	r.c = c
	r.loadGeometry()
	return r
}

func (r *GeometryHandler) Reset() {
	r.timeLastSent = nil
}

func (r *GeometryHandler) loadGeometry() {
	r.geometryDivA = new(vision.SSL_GeometryData)
	if err := prototext.Unmarshal([]byte(geometryConfigDivA), r.geometryDivA); err != nil {
		log.Println("Could not unmarshal geometry file: ", err)
	}
	r.geometryDivB = new(vision.SSL_GeometryData)
	if err := prototext.Unmarshal([]byte(geometryConfigDivB), r.geometryDivB); err != nil {
		log.Println("Could not unmarshal geometry file: ", err)
	}
}

func (r *GeometryHandler) handleGeometry() {
	if *r.c.lastRefereeMsg.Command != referee.Referee_HALT {
		// Only during HALT
		return
	}

	if r.timeLastSent != nil && time.Now().Sub(*r.timeLastSent) < 5*time.Second {
		// Wait for the config to be applied and the update being sent back
		return
	}

	maxBots := int(*r.c.lastRefereeMsg.Yellow.MaxAllowedBots) +
		len(r.c.lastRefereeMsg.Yellow.YellowCardTimes) +
		int(*r.c.lastRefereeMsg.Yellow.RedCards)
	var geometry *vision.SSL_GeometryData
	if maxBots == 6 {
		geometry = r.geometryDivB
	} else {
		geometry = r.geometryDivA
	}

	if *r.c.fieldSize.FieldLength != *geometry.Field.FieldLength {
		r.timeLastSent = new(time.Time)
		*r.timeLastSent = time.Now()
		r.sendConfig(geometry)
	}
}

func (r *GeometryHandler) sendConfig(geometry *vision.SSL_GeometryData) {
	log.Printf("Sending geometry %v", geometry)

	command := SimulatorCommand{
		Config: &SimulatorConfig{
			Geometry: geometry,
		},
	}

	if data, err := proto.Marshal(&command); err != nil {
		log.Println("Could not marshal command: ", err)
	} else {
		r.c.simControlClient.Send(data)
	}
}
