SHELL := /bin/bash

include Makefile.environment
export $(shell sed 's/=.*//' Makefile.environment)

DIRS_TO_CHECK=$(shell ls -d */ | grep -vE "vendor|tmp")
PKGS_TO_CHECK=$(shell go list ./... | grep -vE "/vendor")
OS=$(shell uname | awk '{print tolower($$0)}')
PASS=$(shell printf "\033[32mPASS\033[0m")
FAIL=$(shell printf "\033[31mFAIL\033[0m")
COLORIZE=sed ''/PASS/s//$(PASS)/'' | sed ''/FAIL/s//$(FAIL)/'' && [ $${PIPESTATUS[0]} -eq 0 ]

ifneq (${PKG},)
	PKGS_TO_CHECK="tediouscoder/paper-robot/${PKG}"
endif

ifneq (${RUN},)
	TEST_TO_RUN=-run ${RUN}
endif

VERSION="latest"
ifneq (${VER},)
	VERSION=${VER}
endif

PROJECT_NAME:=paper-robot
PACKAGE_NAME:=tediouscoder/paper-robot
BINARY_NAME:=paper-robot

.PHONY: help
help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  all            to check, build, test and release"
	@echo "  check          to vet and lint"
	@echo "  build          to build all"
	@echo "  test           to run test"
	@echo "  test-benchmark to run test with benchmark"
	@echo "  test-coverage  to run test with coverage"
	@echo "  start          to start server"
	@echo "  clean          to clean the test and built files"

.PHONY: all
all: check build

.PHONY: check
check: format vet lint

.PHONY: format
format:
	@echo "Formatting packages using gofmt..."
	@find . -path "*/vendor/*" -o -path "*/tmp/*" -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;
	@echo "Done"

.PHONY: vet
vet:
	@echo "Checking packages using go tool vet, skip vendor packages..."
	@go vet
	@echo "Done"

.PHONY: lint
lint:
	@echo "Checking packages using golint, skip vendor packages..."
	@lint=$$(for pkg in ${PKGS_TO_CHECK}; do golint $${pkg}; done); \
	 if [[ -n $${lint} ]]; then echo "$${lint}"; exit 1; fi
	@echo "Done"

.PHONY: build
build: check
	@mkdir -p ./bin
	@echo "Building ${BINARY_NAME}..."
	@GOOS=${OS} GOARCH=amd64 go build -o ./bin/${BINARY_NAME} .
	@echo "Done"

.PHONY: test
test:
	@echo "Running test..."
	@go test -v ${TEST_TO_RUN} ${PKGS_TO_CHECK} | $(COLORIZE)
	@echo "Done"

.PHONY: test-benchmark
test-benchmark:
	@echo "Running test with benchmark..."
	@go test -v -bench=. -benchmem ${TEST_TO_RUN} ${PKGS_TO_CHECK} | $(COLORIZE)
	@echo "Done"

.PHONY: test-coverage
test-coverage:
	@echo "Running test with coverage..."
	@for pkg in ${PKGS_TO_CHECK}; do \
		output="coverage$${pkg#${PACKAGE_NAME}}"; \
		mkdir -p $${output}; \
		go test -v -cover -coverprofile="$${output}/profile.out" $${pkg} | $(COLORIZE); \
		if [[ -e "$${output}/profile.out" ]]; then \
			go tool cover -html="$${output}/profile.out" \
			              -o "$${output}/profile.html"; \
		fi; \
	 done
	@echo "Done"

.PHONY: start
start: build
	@echo "Starting server..."
	@./bin/${BINARY_NAME}
	@echo "Done"

.PHONY: clean
clean:
	@echo "Clean the test and built files"
	rm -rf ./bin
	rm -rf ./release
	rm -rf ./coverage
	@echo "Done"
