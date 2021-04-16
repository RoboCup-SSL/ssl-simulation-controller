package simctl

import (
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/geom"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/tracker"
	"github.com/golang/protobuf/proto"
	"log"
	"math"
	"sort"
	"time"
)

type RobotCountMaintainer struct {
	c *SimulationController

	lastTimeSendCommand time.Time
	haltTime            *time.Time
}

func (r *RobotCountMaintainer) handleRobotCount() {

	if time.Now().Sub(r.lastTimeSendCommand) < 500*time.Millisecond {
		// Placed ball just recently
		return
	}

	if *r.c.lastRefereeMsg.Command != referee.Referee_HALT &&
		len(r.c.lastTrackedFrame.TrackedFrame.Balls) > 0 &&
		math.Abs(float64(*r.c.lastTrackedFrame.TrackedFrame.Balls[0].Pos.X)) < 2 {
		// Rule: The ball must be at least 2 meters away from the halfway line.
		return
	}

	var blueRobots []*tracker.TrackedRobot
	var yellowRobots []*tracker.TrackedRobot
	for _, robot := range r.c.lastTrackedFrame.TrackedFrame.Robots {
		if *robot.RobotId.Team == referee.Team_BLUE {
			blueRobots = append(blueRobots, robot)
		} else if *robot.RobotId.Team == referee.Team_YELLOW {
			yellowRobots = append(yellowRobots, robot)
		}
	}

	r.updateRobotCount(blueRobots, int(*r.c.lastRefereeMsg.Blue.MaxAllowedBots), referee.Team_BLUE)
	r.updateRobotCount(yellowRobots, int(*r.c.lastRefereeMsg.Yellow.MaxAllowedBots), referee.Team_YELLOW)
}

func (r *RobotCountMaintainer) updateRobotCount(robots []*tracker.TrackedRobot, maxRobots int, team referee.Team) {
	substCenterPos := geom.NewVector2(0, float64(*r.c.fieldSize.FieldWidth)/2000+float64(*r.c.fieldSize.BoundaryWidth)/2000.0-0.1)
	substCenterNeg := geom.NewVector2Float32(0, -*substCenterPos.Y)
	substRectPos := geom.NewRectangleFromCenter(substCenterPos, 2, float64(*r.c.fieldSize.BoundaryWidth)/1000+0.2)
	substRectNeg := geom.NewRectangleFromCenter(substCenterNeg, 2, float64(*r.c.fieldSize.BoundaryWidth)/1000+0.2)
	if len(robots) > 0 && len(robots) > maxRobots {
		r.sortRobotsByDistanceToSubstitutionPos(robots)
		if *r.c.lastRefereeMsg.Command == referee.Referee_HALT ||
			substRectPos.IsPointInside(robots[0].Pos) ||
			substRectNeg.IsPointInside(robots[0].Pos) {
			r.removeRobot(robots[0].RobotId)
		}
	}
	if len(robots) < maxRobots {
		y := float64(*r.c.fieldSize.FieldWidth) / 2000.0
		if team == referee.Team_BLUE {
			y *= -1
		}
		x := 0.1
		for i := 0; i < 100; i++ {
			pos := geom.NewVector2(x, y)
			if r.isFreeOfObstacles(pos) {
				id := r.nextFreeRobotId(team)
				r.addRobot(id, pos)
				break
			}
			x *= -1
			if x > 0 {
				x += 0.1
			}
		}
	}
}

func (r *RobotCountMaintainer) nextFreeRobotId(team referee.Team) *referee.RobotId {
	for i := 0; i < 16; i++ {
		id := uint32(i)
		robotId := &referee.RobotId{
			Id:   &id,
			Team: &team,
		}
		if r.isRobotIdFree(robotId) {
			return robotId
		}
	}
	return nil
}

func (r *RobotCountMaintainer) isRobotIdFree(id *referee.RobotId) bool {

	for _, robot := range r.c.lastTrackedFrame.TrackedFrame.Robots {
		if *robot.RobotId.Id == *id.Id && *robot.RobotId.Team == *id.Team {
			return false
		}
	}
	return true
}

func (r *RobotCountMaintainer) isFreeOfObstacles(pos *geom.Vector2) bool {
	for _, robot := range r.c.lastTrackedFrame.TrackedFrame.Robots {
		if robot.Pos.DistanceTo(pos) < 0.2 {
			return false
		}
	}
	for _, ball := range r.c.lastTrackedFrame.TrackedFrame.Balls {
		pos2d := geom.NewVector2Float32(*ball.Pos.X, *ball.Pos.Y)
		if pos2d.DistanceTo(pos) < 0.1 {
			return false
		}
	}
	return true
}

func (r *RobotCountMaintainer) sortRobotsByDistanceToSubstitutionPos(robots []*tracker.TrackedRobot) {
	negSubstPos := geom.NewVector2(0, -float64(*r.c.fieldSize.FieldWidth)/2000)
	posSubstPos := geom.NewVector2(0, +float64(*r.c.fieldSize.FieldWidth)/2000)
	sort.Slice(robots, func(i, j int) bool {
		distI := math.Min(robots[i].Pos.DistanceTo(negSubstPos), robots[i].Pos.DistanceTo(posSubstPos))
		distJ := math.Min(robots[j].Pos.DistanceTo(negSubstPos), robots[j].Pos.DistanceTo(posSubstPos))
		return distI < distJ
	})
}

func (r *RobotCountMaintainer) removeRobot(id *referee.RobotId) {
	log.Printf("Remove robot %v", id)

	present := false
	command := SimulatorCommand{
		Control: &SimulatorControl{
			TeleportRobot: []*TeleportRobot{
				{
					Id:      id,
					Present: &present,
				},
			},
		},
	}

	r.sendControlCommand(&command)
}

func (r *RobotCountMaintainer) addRobot(id *referee.RobotId, pos *geom.Vector2) {
	log.Printf("Add robot %v @ %v", id, pos)

	present := true
	orientation := float32(0)
	command := SimulatorCommand{
		Control: &SimulatorControl{
			TeleportRobot: []*TeleportRobot{
				{
					Id:          id,
					X:           pos.X,
					Y:           pos.Y,
					Orientation: &orientation,
					Present:     &present,
				},
			},
		},
	}

	r.sendControlCommand(&command)
}

func (r *RobotCountMaintainer) sendControlCommand(command *SimulatorCommand) {

	if data, err := proto.Marshal(command); err != nil {
		log.Println("Could not marshal command: ", err)
	} else {
		r.c.simControlClient.Send(data)
		r.lastTimeSendCommand = time.Now()
	}
}
