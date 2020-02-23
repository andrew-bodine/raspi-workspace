# This is the Makefile for raspi-workspace.

ginkgo:
	go install github.com/onsi/ginkgo/ginkgo

# test: unit-tests
#
# unit-tests: ginkgo fakes
# 	./scripts/run-unit-tests.bash

# integration-tests:
# 	./scripts/run-integration-tests
#
# integration-tests-local: ginkgo
# 	./scripts/run-integration-tests-local

build:
	./scripts/run-builds.bash

fakes:
	./scripts/create-fakes.bash

# Install all necessary local golang dependencies.
go-deps:
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega
	go get github.com/stianeikeland/go-rpio
	go get github.com/satori/go.uuid
	go get go.uber.org/zap

clean:
	rm -rf build

clean-all: clean
	rm -rf vendor/*
