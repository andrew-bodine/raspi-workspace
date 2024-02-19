# This is the Makefile for raspi-workspace.

all: codegen test build

setup:
	./scripts/setup.bash

test:
	./scripts/test.bash

build:
	./scripts/build.bash

codegen:
	./scripts/codegen.bash

clean:
	rm -rf build
