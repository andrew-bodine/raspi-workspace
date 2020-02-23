#!/bin/bash

set -e

output=${WORKSPACE}/build

GIT_SHA=$(git rev-parse --verify HEAD || echo "unknown")

echo "Starting build: " $(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Build the target packages to the corresponding output build context.
pushd ${WORKSPACE}/src/github.com/andrew-bodine > /dev/null
    for pathSegment in $(find * -name cmd); do
        for buildTarget in $(ls ${pathSegment}); do
            echo "Building ${buildTarget} target from github.com/andrew-bodine/${pathSegment} package."

            time \
            env GOOS=linux GOARCH=arm GOARM=5 \
            go build \
                -o ${output}/${buildTarget} \
                github.com/andrew-bodine/${pathSegment}/${buildTarget}
        done
    done
popd > /dev/null

echo "Build finished: " $(date -u +"%Y-%m-%dT%H:%M:%SZ")
