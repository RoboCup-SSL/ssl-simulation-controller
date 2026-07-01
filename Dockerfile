FROM golang:1.26-alpine@sha256:f85330846cde1e57ca9ec309382da3b8e6ae3ab943d2739500e08c86393a21b1 AS build
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-simulation-controller
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY cmd cmd
COPY internal internal
RUN go install ./...

# Start fresh from a smaller image
FROM alpine:3@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b
COPY --from=build /go/bin/ssl-simulation-controller /app/ssl-simulation-controller
USER 1000
ENTRYPOINT ["/app/ssl-simulation-controller"]
CMD []
