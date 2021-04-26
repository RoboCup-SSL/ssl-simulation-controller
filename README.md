[![CircleCI](https://circleci.com/gh/RoboCup-SSL/ssl-simulation-controller/tree/master.svg?style=svg)](https://circleci.com/gh/RoboCup-SSL/ssl-simulation-controller/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/RoboCup-SSL/ssl-simulation-controller?style=flat-square)](https://goreportcard.com/report/github.com/RoboCup-SSL/ssl-simulation-controller)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/RoboCup-SSL/ssl-simulation-controller/pkg/vision)
[![Release](https://img.shields.io/github/release/RoboCup-SSL/ssl-simulation-controller.svg?style=flat-square)](https://github.com/RoboCup-SSL/ssl-simulation-controller/releases/latest)

# ssl-simulation-controller

A controller for SSL simulation tournaments.

## Behaviors

### Manual ball placement
The ball is manually moved to the designated target position,
if the game is in HALT for at least 0.5s.

### Maintain robot count
Adds and removes robots, if the robot count does not match the max robot count in the referee message:
* The [robot substitution rules](https://robocup-ssl.github.io/ssl-rules/sslrules.html#_robot_substitution) are applied as far as possible.
* Robots are only added or removed, if the game is either in HALT, or the ball is not within 2m to the halfway line (see robot substitution rules).
* Robots are put near the halfway line at a free position.
* Robots nearest to either crossing of halfway line and touch lines are removed first.

### Apply Geometry
The geometry for division A or B is applied, based on the max number of robots in the referee message.
The geometry is currently hard-coded into the binary.
It will only be applied in `NORMAL_FIRST_HALF_PRE` stage.

### Apply robot specs
Send the robot specs for the respective teams to the simulator:
* The robot specs are taken from a config file specified via command line. There is also an example config in [config/robot-specs.yaml](config/robot-specs.yaml).
* The team names are taken from the referee message and must match exactly with the keys in the config file.
* Robot specs are only applied during pre-stages.
* The simulator will receive equal robot specs for all 16 robots in one message.
* Robot specs are only send once on startup and when a team changes. The simulator-controller will not resend the specs after a simulator crash and must be restarted, or the team name must be switched to force resending the specs.

## Usage
If you just want to use this app, simply download the latest [release binary](https://github.com/RoboCup-SSL/ssl-simulation-controller/releases/latest).
The binary is self-contained. No dependencies are required.

You can also use pre-build docker images:
```shell script
docker pull robocupssl/ssl-simulation-controller
docker run --network host robocupssl/ssl-simulation-controller
```

## Development

### Requirements
You need to install following dependencies first: 
 * Go >= 1.14

### Prepare
Download and install to [GOPATH](https://github.com/golang/go/wiki/GOPATH):
```bash
go get -u github.com/RoboCup-SSL/ssl-simulation-controller/...
```
Switch to project root directory
```bash
cd $GOPATH/src/github.com/RoboCup-SSL/ssl-simulation-controller/
```

### Run
Run the backend:
```bash
go run cmd/ssl-simulation-controller/main.go
```
