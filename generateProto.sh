#!/bin/bash

# Fail on errors
set -e
# Print commands
set -x

# common GC
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_gc_common.proto
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_gc_geometry.proto

# vision
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_vision_detection.proto
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_vision_geometry.proto
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_vision_wrapper.proto

# tracked vision
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_vision_detection_tracked.proto
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_vision_wrapper_tracked.proto

# game events
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_gc_game_event.proto

# referee message
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_gc_referee_message.proto

# simulation control
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_simulation_error.proto
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_simulation_config.proto
protoc -I"./proto" -I"$GOPATH/src" --go_out="$GOPATH/src" proto/ssl_simulation_control.proto
