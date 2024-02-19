# This is the Makefile for raspi-workspace.

all: codegen test build

setup:
	./scripts/setup.bash

test: unit-tests

unit-tests:
	./scripts/run-unit-tests.bash

build:
	./scripts/run-builds.bash

codegen:
	./scripts/run-codegen.bash

clean:
	rm -rf build
