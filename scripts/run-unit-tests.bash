#!/bin/bash

set -eu

echo "Beginning workspace unit tests."

export GINKGO_DEFAULT_EVENTUALLY=${GINKGO_DEFAULT_EVENTUALLY:-5s}

SKIP_PACKAGES=""

# Filter out integration test packages.
pushd ${WORKSPACE}/src/github.com/andrew-bodine > /dev/null
    for pkg in $(find *monitor* -name cmd); do
        if [ "${SKIP_PACKAGES}" == "" ]; then
            SKIP_PACKAGES="${pkg}"
        else
            SKIP_PACKAGES="${pkg},${SKIP_PACKAGES}"
        fi
    done
popd > /dev/null

EXIT_CODE=0

pushd $WORKSPACE/src/github.com/andrew-bodine > /dev/null
    ginkgo -r -p -keepGoing -randomizeAllSpecs -progress --race -skipPackage=${SKIP_PACKAGES} "$@"
    ((EXIT_CODE += $?)) || :
popd > /dev/null

exit $EXIT_CODE
