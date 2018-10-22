# drone-apex-up

[![GoDoc](https://godoc.org/github.com/appleboy/drone-apex-up?status.svg)](https://godoc.org/github.com/appleboy/drone-apex-up)
[![Build Status](http://drone.wu-boy.com/api/badges/appleboy/drone-apex-up/status.svg)](http://drone.wu-boy.com/appleboy/drone-apex-up)
[![codecov](https://codecov.io/gh/appleboy/drone-apex-up/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-apex-up)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-apex-up)](https://goreportcard.com/report/github.com/appleboy/drone-apex-up)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-apex-up.svg)](https://hub.docker.com/r/appleboy/drone-apex-up/)
[![](https://images.microbadger.com/badges/image/appleboy/drone-apex-up.svg)](https://microbadger.com/images/appleboy/drone-apex-up "Get your own image badge on microbadger.com")
[![Build status](https://ci.appveyor.com/api/projects/status/pmkfbnwtlf1fm45l/branch/master?svg=true)](https://ci.appveyor.com/project/appleboy/drone-apex-up/branch/master)

Drone plugin for deploying infinitely scalable serverless apps using apex/up command.

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-apex-up
docker build --rm -t appleboy/drone-apex-up .
```
