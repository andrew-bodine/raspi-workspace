#!/bin/bash

set -e

GIT_SHA=$(git rev-parse --verify HEAD || echo "unknown")

echo "Starting build: " $(date -u +"%Y-%m-%dT%H:%M:%SZ")

output=${WORKSPACE}/build/monitoring

# Build the target package to the corresponding output build context.
time \
env GOOS=linux GOARCH=arm GOARM=5 \
go build \
    -o ${output}/monitor \
    github.com/andrew-bodine/monitoring/cmd/monitoring

echo "Build finished: " $(date -u +"%Y-%m-%dT%H:%M:%SZ")
