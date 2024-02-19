#!/bin/bash

set -eu

echo "Beginning workspace unit tests."

export GINKGO_DEFAULT_EVENTUALLY=${GINKGO_DEFAULT_EVENTUALLY:-5s}

SKIP_PACKAGES=""
EXIT_CODE=0

echo "Will skip the following packages: ${SKIP_PACKAGES}"

pushd $WORKSPACE > /dev/null
    ginkgo -r -p -keepGoing -randomizeAllSpecs -progress --race -skipPackage=${SKIP_PACKAGES} "$@"
    ((EXIT_CODE += $?)) || :
popd > /dev/null

exit $EXIT_CODE
