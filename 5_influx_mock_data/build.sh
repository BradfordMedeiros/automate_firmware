#!/usr/bin/env bash

mkdir -p ./build
go build -o ./build/influx_mock_data -ldflags "-s -w" ./src/*
