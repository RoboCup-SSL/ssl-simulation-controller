FROM golang:1.25-alpine@sha256:b6ed3fd0452c0e9bcdef5597f29cc1418f61672e9d3a2f55bf02e7222c014abd AS build
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-simulation-controller
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY cmd cmd
COPY internal internal
RUN go install ./...

# Start fresh from a smaller image
FROM alpine:3@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412
COPY --from=build /go/bin/ssl-simulation-controller /app/ssl-simulation-controller
USER 1000
ENTRYPOINT ["/app/ssl-simulation-controller"]
CMD []
