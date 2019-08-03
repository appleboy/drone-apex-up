# drone-apex-up

[![GoDoc](https://godoc.org/github.com/appleboy/drone-apex-up?status.svg)](https://godoc.org/github.com/appleboy/drone-apex-up)
[![Build Status](https://drone.ggz.tw/api/badges/appleboy/drone-apex-up/status.svg)](https://drone.ggz.tw/appleboy/drone-apex-up)
[![codecov](https://codecov.io/gh/appleboy/drone-apex-up/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-apex-up)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-apex-up)](https://goreportcard.com/report/github.com/appleboy/drone-apex-up)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-apex-up.svg)](https://hub.docker.com/r/appleboy/drone-apex-up/)
[![badger](https://images.microbadger.com/badges/image/appleboy/drone-apex-up.svg)](https://microbadger.com/images/appleboy/drone-apex-up "Get your own image badge on microbadger.com")
[![Build status](https://ci.appveyor.com/api/projects/status/pmkfbnwtlf1fm45l/branch/master?svg=true)](https://ci.appveyor.com/project/appleboy/drone-apex-up/branch/master)

Drone plugin for deploying infinitely scalable serverless apps using apex/up command.

## Build or Download a binary

The pre-compiled binaries can be downloaded from [release page](https://github.com/appleboy/drone-apex-up/releases). Support the following OS type.

* Windows amd64/386
* Linux arm/amd64/386
* Darwin amd64/386

With `Go` installed

```sh
go get -u -v github.com/appleboy/drone-apex-up
```

or build the binary with the following command:

```sh
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go test -cover ./...

go build -v -a -tags netgo -o release/linux/amd64/drone-apex-up .
```

## Docker

Build the docker image with the following commands:

```sh
make docker
```
