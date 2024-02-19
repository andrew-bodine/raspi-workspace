#!/bin/bash

set -eu

echo "Beginning workspace fakes generation."

# Paths to be searched for `go:generate` comments.
pathsOfInterest="${WORKSPACE}/monitor ${WORKSPACE}/release"

# Contains uniq found dirs where `go generate` needs to be run.
uniqueCodegenDirs=""

for path in ${pathsOfInterest}; do

    # Find file paths that contain 'go:generate' comments in path.
    allComments=$(grep -r "go:generate" ${path} | cut -d ':' -f 1)

    # Compress search results to uniq containing directories.
    uniqueDirs=$(for comment in ${allComments}; do dirname $comment; done | uniq)

    uniqueCodegenDirs="${uniqueDirs} ${uniqueCodegenDirs}"
done

for dir in ${uniqueCodegenDirs}; do
    pushd ${dir} > /dev/null
        echo "Generating fakes at ${dir}"
        go generate
    popd > /dev/null
done

echo "Done generating workspace fakes."
