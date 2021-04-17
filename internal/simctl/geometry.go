package simctl

import (
	_ "embed"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/vision"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

//go:embed geometry-div-a.txt
var geometryConfigDivA string

//go:embed geometry-div-b.txt
var geometryConfigDivB string

type GeometrySetter struct {
	c *SimulationController

	geometryDivA *vision.SSL_GeometryData
	geometryDivB *vision.SSL_GeometryData
	timeLastSent *time.Time
}

func (r *GeometrySetter) LoadGeometry() {
	r.geometryDivA = new(vision.SSL_GeometryData)
	if err := proto.UnmarshalText(geometryConfigDivA, r.geometryDivA); err != nil {
		log.Println("Could not unmarshal geometry file: ", err)
	}
	r.geometryDivB = new(vision.SSL_GeometryData)
	if err := proto.UnmarshalText(geometryConfigDivB, r.geometryDivB); err != nil {
		log.Println("Could not unmarshal geometry file: ", err)
	}
}

func (r *GeometrySetter) handleGeometry() {
	if *r.c.lastRefereeMsg.Stage != referee.Referee_NORMAL_FIRST_HALF_PRE {
		// Only before the game starts
		return
	}

	if r.timeLastSent != nil && time.Now().Sub(*r.timeLastSent) < 5*time.Second {
		// Wait for the config to be applied and the update being sent back
		return
	}

	maxBots := *r.c.lastRefereeMsg.Yellow.MaxAllowedBots
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

func (r *GeometrySetter) sendConfig(geometry *vision.SSL_GeometryData) {
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
