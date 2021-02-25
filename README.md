[![CircleCI](https://circleci.com/gh/RoboCup-SSL/ssl-simulation-controller/tree/master.svg?style=svg)](https://circleci.com/gh/RoboCup-SSL/ssl-simulation-controller/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/RoboCup-SSL/ssl-simulation-controller?style=flat-square)](https://goreportcard.com/report/github.com/RoboCup-SSL/ssl-simulation-controller)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/RoboCup-SSL/ssl-simulation-controller/pkg/vision)
[![Release](https://img.shields.io/github/release/RoboCup-SSL/ssl-simulation-controller.svg?style=flat-square)](https://github.com/RoboCup-SSL/ssl-simulation-controller/releases/latest)

# ssl-simulation-controller

A controller for SSL simulation tournaments, that:
 * places the ball automatically, if it can not be placed by the teams
 * TODO: Adds and removes robots, if the robot count does not fit
   * During gameplay, if the [conditions](https://robocup-ssl.github.io/ssl-rules/sslrules.html#_robot_substitution) are met
   * During stop by selecting robots nearest to either crossing of half-way line and touch line)
 * TODO: Set robot limits based on some configuration

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
