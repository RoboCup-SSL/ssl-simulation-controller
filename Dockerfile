FROM golang:1.24-alpine@sha256:daae04ebad0c21149979cd8e9db38f565ecefd8547cf4a591240dc1972cf1399 AS build
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-simulation-controller
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY cmd cmd
COPY internal internal
RUN go install ./...

# Start fresh from a smaller image
FROM alpine:3@sha256:4bcff63911fcb4448bd4fdacec207030997caf25e9bea4045fa6c8c44de311d1
COPY --from=build /go/bin/ssl-simulation-controller /app/ssl-simulation-controller
USER 1000
ENTRYPOINT ["/app/ssl-simulation-controller"]
CMD []
