package simctl

import (
	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type TeamRobotSpecs struct {
	Teams map[string]RobotSpec `yaml:"robot-specs"`
}

type RobotSpec struct {
	Radius             float32                `yaml:"radius"`
	Height             float32                `yaml:"height"`
	Mass               float32                `yaml:"mass"`
	MaxLinearKickSpeed float32                `yaml:"max_linear_kick_speed"`
	MaxChipKickSpeed   float32                `yaml:"max_chip_kick_speed"`
	CenterToDribbler   float32                `yaml:"center_to_dribbler"`
	Limits             Limits                 `yaml:"limits"`
	CustomErforce      CustomRobotSpecErForce `yaml:"custom_erforce"`
}

type Limits struct {
	AccSpeedupAbsoluteMax float32 `yaml:"acc_speedup_absolute_max,omitempty"`
	AccSpeedupAngularMax  float32 `yaml:"acc_speedup_angular_max,omitempty"`
	AccBrakeAbsoluteMax   float32 `yaml:"acc_brake_absolute_max,omitempty"`
	AccBrakeAngularMax    float32 `yaml:"acc_brake_angular_max,omitempty"`
	VelAbsoluteMax        float32 `yaml:"vel_absolute_max,omitempty"`
	VelAngularMax         float32 `yaml:"vel_angular_max,omitempty"`
}

type CustomRobotSpecErForce struct {
	ShootRadius   float32 `yaml:"shoot_radius"`
	DribblerWidth float32 `yaml:"dribbler_width"`
}

type RobotSpecHandler struct {
	c *SimulationController

	teamRobotSpecs TeamRobotSpecs
	appliedTeams   map[referee.Team]string
}

func NewRobotSpecHandler(c *SimulationController, configFile string) (r *RobotSpecHandler) {
	r = new(RobotSpecHandler)
	r.c = c
	r.loadRobotSpecs(configFile)
	return r
}

func (r *RobotSpecHandler) Reset() {
	r.appliedTeams = map[referee.Team]string{}
}

func (r *RobotSpecHandler) loadRobotSpecs(configFile string) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println("Could not read robot spec file: ", err)
	} else if err := yaml.Unmarshal(data, &r.teamRobotSpecs); err != nil {
		log.Println("Could not unmarshal robot spec file: ", err)
	}
}

func (r *RobotSpecHandler) handleRobotSpecs() {
	switch *r.c.lastRefereeMsg.Stage {
	case referee.Referee_NORMAL_FIRST_HALF_PRE,
		referee.Referee_NORMAL_SECOND_HALF_PRE,
		referee.Referee_EXTRA_FIRST_HALF_PRE,
		referee.Referee_EXTRA_SECOND_HALF_PRE:
		// accept
	default:
		// Only in pre-stages
		return
	}

	r.updateTeam(referee.Team_BLUE, *r.c.lastRefereeMsg.Blue.Name)
	r.updateTeam(referee.Team_YELLOW, *r.c.lastRefereeMsg.Yellow.Name)
}

func (r *RobotSpecHandler) updateTeam(team referee.Team, teamName string) {
	if r.appliedTeams[team] != teamName {
		if spec, ok := r.teamRobotSpecs.Teams[teamName]; ok {
			r.applySpecs(team, teamName, spec)
		} else if spec, ok := r.teamRobotSpecs.Teams["Unknown"]; ok {
			log.Printf("Team %v not found, using fallback", teamName)
			r.applySpecs(team, teamName, spec)
		} else {
			log.Printf("Team %v not found and also no fallback found", teamName)
		}
	}
}

func (r *RobotSpecHandler) applySpecs(team referee.Team, teamName string, spec RobotSpec) {
	var protoSpecs []*RobotSpecs
	for id := 0; id < 16; id++ {
		protoSpec := mapRobotSpec(spec)
		protoSpec.Id = new(referee.RobotId)
		protoSpec.Id.Id = new(uint32)
		protoSpec.Id.Team = new(referee.Team)
		*protoSpec.Id.Team = team
		*protoSpec.Id.Id = uint32(id)
		protoSpecs = append(protoSpecs, protoSpec)
	}
	r.sendConfig(protoSpecs)
	r.appliedTeams[team] = teamName
}

func (r *RobotSpecHandler) sendConfig(robotSpec []*RobotSpecs) {
	log.Printf("Sending robot spec %v", robotSpec)

	command := SimulatorCommand{
		Config: &SimulatorConfig{
			RobotSpecs: robotSpec,
		},
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

	customErForce := RobotSpecErForce{
		ShootRadius:   &spec.CustomErforce.ShootRadius,
		DribblerWidth: &spec.CustomErforce.DribblerWidth,
	}
	customErForceSerialized, err := proto.Marshal(&customErForce)
	if err != nil {
		log.Println("Could not serialize custom ER-Force robot specs: ", err)
	}
	customErForce.ProtoReflect().Descriptor().FullName()
	protoSpec.Custom = append(protoSpec.Custom, &anypb.Any{
		TypeUrl: "type.googleapis.com/sslsim.RobotSpecErForce",
		Value:   customErForceSerialized,
	})
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
