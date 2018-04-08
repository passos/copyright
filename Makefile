hide := @

PROJECT := copyright
BUILD_DIR := ${GOPATH}/src/${PROJECT}

BUILD_TIME := $(shell date +"%Y-%m-%dT%H:%M:%SZ")
COMMIT := $(shell git rev-parse HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

FLAGS :=
#FLAGS := -race
LD_FLAGS := "-X main.Version=1.0.0 -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)"
GO := go
GO_BUILD := $(GO) build $(FLAGS) -ldflags $(LD_FLAGS)

SOURCES := $(wildcard *.go)

all: build
	$(info ; all done)

.PHONY: build
build: copyright-server

.PHONY: copyright-server
copyright-server: $(SOURCES)
	$(info ; building $(PROJECT))
	$(hide) $(GO_BUILD)
	$(hide) mv $(PROJECT) bin/

.PHONY: install-deps
install-deps:
	$(info ; installing dependencies)
	$(hide) bin/install-deps

.PHONY: restart-super
restart-super:
	$(info ; restarting supervisor job)
	$(hide) supervisorctl restart copyright-server; supervisorctl status

repo ?= copyright-server
remote ?= copyright-server

.PHONY: publish
publish:
	$(info ; deploy to $(remote))
	$(hide) git push -f $(repo)
	$(hide) ssh $(remote) 'cd $$GOPATH/src/$(PROJECT); git reset --hard master; git log -3 --pretty=oneline'


.PHONY: deploy
deploy: publish
	$(hide) ssh $(remote) 'cd $$GOPATH/src/$(PROJECT); make install-deps; make; make restart-super'

.PHONY: clean
clean:
	$(hide) -rm -f bin/copyright-server
