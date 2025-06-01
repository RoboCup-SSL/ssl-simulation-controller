FROM golang:1.24-alpine@sha256:b4f875e650466fa0fe62c6fd3f02517a392123eea85f1d7e69d85f780e4db1c1 AS build
WORKDIR /go/src/github.com/RoboCup-SSL/ssl-simulation-controller
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY cmd cmd
COPY internal internal
RUN go install ./...

# Start fresh from a smaller image
FROM alpine:3@sha256:8a1f59ffb675680d47db6337b49d22281a139e9d709335b492be023728e11715
COPY --from=build /go/bin/ssl-simulation-controller /app/ssl-simulation-controller
USER 1000
ENTRYPOINT ["/app/ssl-simulation-controller"]
CMD []
