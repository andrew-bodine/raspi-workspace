#!/bin/bash

set -e

output=${WORKSPACE}/build

mkdir -p $output

GIT_SHA=$(git rev-parse --verify HEAD || echo "unknown")

echo "Starting build: " $(date -u +"%Y-%m-%dT%H:%M:%SZ")

pkgs=""
pkgs="monitor ${pkgs}"
pkgs="release ${pkgs}"

backupIFS=${IFS}
IFS=" "

# Build the target packages to the corresponding output build context.
for pkg in ${pkgs}; do
    pushd ${pkg} > /dev/null
        echo "Building raspi-${pkg} target from github.com/andrew-bodine/raspi/${pkg} package."
        
        # Cross-compile the target package.
        #
        # TODO: Should this be GOARM=7 ?
        time env GOOS=linux GOARCH=arm GOARM=5 go build \
            -o ${output}/raspi-${pkg}/raspi-${pkg} \
            github.com/andrew-bodine/raspi/${pkg}

        # If the target package contains a systemd unit file, copy it to
        # the build context for packaging.
        if [ -f ${pkg}/raspi-${pkg}@.service ]; then
            cp ${pkg}/raspi-${pkg}@.service \
                ${output}/raspi-${pkg}/
        fi
    popd > /dev/null
done

IFS=${backupIFS}

echo "Build finished: " $(date -u +"%Y-%m-%dT%H:%M:%SZ")
