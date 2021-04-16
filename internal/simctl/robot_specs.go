package simctl

import (
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
	"github.com/golang/protobuf/proto"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type TeamRobotSpecs struct {
	Teams map[string]RobotSpec `yaml:"robot-specs"`
}

type RobotSpec struct {
	Radius             float32 `yaml:"radius"`
	Height             float32 `yaml:"height"`
	Mass               float32 `yaml:"mass"`
	MaxLinearKickSpeed float32 `yaml:"max_linear_kick_speed"`
	MaxChipKickSpeed   float32 `yaml:"max_chip_kick_speed"`
	CenterToDribbler   float32 `yaml:"center_to_dribbler"`
	Limits             Limits  `yaml:"limits"`
}

type Limits struct {
	AccSpeedupAbsoluteMax float32 `yaml:"acc_speedup_absolute_max,omitempty"`
	AccSpeedupAngularMax  float32 `yaml:"acc_speedup_angular_max,omitempty"`
	AccBrakeAbsoluteMax   float32 `yaml:"acc_brake_absolute_max,omitempty"`
	AccBrakeAngularMax    float32 `yaml:"acc_brake_angular_max,omitempty"`
	VelAbsoluteMax        float32 `yaml:"vel_absolute_max,omitempty"`
	VelAngularMax         float32 `yaml:"vel_angular_max,omitempty"`
}

type RobotSpecSetter struct {
	c *SimulationController

	teamRobotSpecs TeamRobotSpecs
	appliedTeams   map[referee.Team]string
}

func (r *RobotSpecSetter) LoadRobotSpecs(configFile string) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println("Could not read robot spec file: ", err)
	}
	if err := yaml.Unmarshal(data, &r.teamRobotSpecs); err != nil {
		log.Println("Could not unmarshal robot spec file: ", err)
	}
	r.appliedTeams = map[referee.Team]string{}
}

func (r *RobotSpecSetter) handleRobotSpecs() {
	if *r.c.lastRefereeMsg.Command != referee.Referee_HALT {
		// Only during HALT
		return
	}

	r.updateTeam(referee.Team_BLUE, *r.c.lastRefereeMsg.Blue.Name)
	r.updateTeam(referee.Team_YELLOW, *r.c.lastRefereeMsg.Yellow.Name)
}

func (r *RobotSpecSetter) updateTeam(team referee.Team, teamName string) {
	if r.appliedTeams[team] != teamName {
		if spec, ok := r.teamRobotSpecs.Teams[teamName]; ok {
			protoSpec := mapRobotSpec(spec)
			protoSpec.Id = new(referee.RobotId)
			protoSpec.Id.Id = new(uint32)
			protoSpec.Id.Team = new(referee.Team)
			*protoSpec.Id.Id = 0
			*protoSpec.Id.Team = team
			r.sendConfig(protoSpec)
			r.appliedTeams[team] = teamName
		}
	}
}

func (r *RobotSpecSetter) sendConfig(robotSpec *RobotSpecs) {
	log.Printf("Sending robot spec %v", robotSpec)

	command := SimulatorCommand{
		Config: &SimulatorConfig{
			RobotSpecs: []*RobotSpecs{
				robotSpec,
			}},
	}

	if data, err := proto.Marshal(&command); err != nil {
		log.Println("Could not marshal command: ", err)
	} else {
		r.c.simControlClient.Send(data)
	}
}

func mapRobotSpec(spec RobotSpec) (protoSpec *RobotSpecs) {
	protoSpec = new(RobotSpecs)
	protoSpec.Radius = &spec.Radius
	protoSpec.Height = &spec.Height
	protoSpec.Mass = &spec.Mass
	protoSpec.MaxLinearKickSpeed = &spec.MaxLinearKickSpeed
	protoSpec.MaxChipKickSpeed = &spec.MaxChipKickSpeed
	protoSpec.CenterToDribbler = &spec.CenterToDribbler
	protoSpec.Limits = mapRobotLimits(spec.Limits)
	return
}

func mapRobotLimits(limits Limits) (protoLimits *RobotLimits) {
	protoLimits = new(RobotLimits)
	protoLimits.AccSpeedupAbsoluteMax = &limits.AccSpeedupAbsoluteMax
	protoLimits.AccSpeedupAngularMax = &limits.AccSpeedupAngularMax
	protoLimits.AccBrakeAbsoluteMax = &limits.AccBrakeAbsoluteMax
	protoLimits.AccBrakeAngularMax = &limits.AccBrakeAngularMax
	protoLimits.VelAbsoluteMax = &limits.VelAbsoluteMax
	protoLimits.VelAngularMax = &limits.VelAngularMax
	return
}
