package simctl

import (
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/geom"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

type BallReplaceHandler struct {
	c *SimulationController

	lastTimePlacedBall time.Time
	haltTime           *time.Time
}

func NewBallReplaceHandler(c *SimulationController) (r *BallReplaceHandler) {
	r = new(BallReplaceHandler)
	r.c = c
	return r
}

func (r *BallReplaceHandler) Reset() {
	r.lastTimePlacedBall = time.Time{}
	r.haltTime = nil
}

func (r *BallReplaceHandler) handleReplaceBall() {
	if *r.c.lastRefereeMsg.Command != referee.Referee_HALT {
		// Only during HALT
		r.haltTime = nil
		return
	}
	if r.haltTime == nil {
		r.haltTime = new(time.Time)
		*r.haltTime = time.Now()
	}

	if time.Now().Sub(*r.haltTime) < 500*time.Millisecond {
		// Halt must be set for at least 1 second
		return
	}

	if r.c.lastRefereeMsg.DesignatedPosition == nil {
		// Only when ball placement position set
		return
	}

	if time.Now().Sub(r.lastTimePlacedBall) < 500*time.Millisecond {
		// Placed ball just recently
		return
	}

	targetPos := geom.NewVector2Float32(
		*r.c.lastRefereeMsg.DesignatedPosition.X/1000,
		*r.c.lastRefereeMsg.DesignatedPosition.Y/1000,
	)

	balls := r.c.lastTrackedFrame.TrackedFrame.Balls
	if len(balls) == 0 {
		log.Printf("Ball vanished. Placing ball to %v", targetPos)
		r.placeBall(targetPos)
		return
	}

	currentPos := geom.NewVector2Float32(*balls[0].Pos.X, *balls[0].Pos.Y)

	if targetPos.DistanceTo(currentPos) > 0.1 {
		log.Printf("Placing ball from %v to %v", currentPos, targetPos)
		r.placeBall(targetPos)
	}
}

func (r *BallReplaceHandler) placeBall(ballPos *geom.Vector2) {

	zero := float32(0)
	command := SimulatorCommand{
		Control: &SimulatorControl{
			TeleportBall: &TeleportBall{
				X:  ballPos.X,
				Y:  ballPos.Y,
				Z:  &zero,
				Vx: &zero,
				Vy: &zero,
				Vz: &zero,
			},
		},
	}

	if data, err := proto.Marshal(&command); err != nil {
		log.Println("Could not marshal command: ", err)
	} else {
		r.c.simControlClient.Send(data)
		r.lastTimePlacedBall = time.Now()
	}
}
