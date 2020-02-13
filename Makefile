# This is the Makefile for raspi-workspace.

ginkgo:
	go install github.com/onsi/ginkgo/ginkgo

test: unit-tests

unit-tests: ginkgo
	./scripts/run-unit-tests

# integration-tests:
# 	./scripts/run-integration-tests
#
# integration-tests-local: ginkgo
# 	./scripts/run-integration-tests-local

build:
	./scripts/run-go-builds

# create-workspace-fakes:
# 	./scripts/create-workspace-fakes

# Install all necessary local golang dependencies.
go-deps:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/stianeikeland/go-rpio

clean:
	rm -rf build
