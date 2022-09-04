#!/bin/bash
for pid in $(ps -ef | grep -e "go run main.go" -e "go-build" | awk '{print $2}'); do kill -9 $pid; done
