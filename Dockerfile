FROM golang:1.16-alpine AS build
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-simulation-controller
COPY go.mod go.mod
RUN go mod download
COPY cmd cmd
COPY internal internal
RUN go install ./...

# Start fresh from a smaller image
FROM alpine:3.9
COPY --from=build /go/bin/ssl-simulation-controller /app/ssl-simulation-controller
ENTRYPOINT ["/app/ssl-simulation-controller"]
CMD []
