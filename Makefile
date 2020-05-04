

GO           ?= go
GOFMT        ?= $(GO)fmt
FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
STATICCHECK  := $(FIRST_GOPATH)/bin/staticcheck
GOVENDOR     := $(FIRST_GOPATH)/bin/govendor
GODEP				 := $(FIRST_GOPATH)/bin/dep
RPM          := ./scripts/build_rpm.sh
pkgs          = ./...

BIN_DIR                 ?= $(shell pwd)/build
VERSION ?= $(shell cat VERSION)
REVERSION ?=$(shell git log -1 --pretty="%H")
BRANCH ?=$(shell git rev-parse --abbrev-ref HEAD)
TIME ?=$(shell date)
HOST ?=$(shell hostname)  
DOCKER := $(shell { command -v podman || command -v docker; } 2>/dev/null)

all:   fmt style  build  docker-build rpm docker-rpm

 
style:
	@echo ">> checking code style"
	! $(GOFMT) -d $$(find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

check_license:
	@echo ">> checking license header"
	@licRes=$$(for file in $$(find . -type f -iname '*.go' ! -path './vendor/*') ; do \
               awk 'NR<=3' $$file | grep -Eq "(Copyright|generated|GENERATED)" || echo $$file; \
       done); \
       if [ -n "$${licRes}" ]; then \
               echo "license header checking failed:"; echo "$${licRes}"; \
               exit 1; \
       fi

build: 
	@echo ">> building binaries"
	$(GO) build  -o $(BIN_DIR)/netapp_exporter  -ldflags  '-X "main.Vsersion=$(VERSION)" -X  "main.BuildRevision=$(REVERSION)" -X  "main.BuildBranch=$(BRANCH)" -X "main.BuildTime=$(TIME)" -X "main.BuildHost=$(HOST)"'

docker-build:
	$(DOCKER) run -v `pwd`:/go/src/github.com/jenningsloy318/netapp_exporter  -w /go/src/github.com/jenningsloy318/netapp_exporter docker.io/jenningsloy318/prom-builder  make build

rpm: | build
	@echo ">> building rpm package"
	$(RPM) 

docker-rpm:
	$(DOCKER) run -v `pwd`:/go/src/github.com/jenningsloy318/netapp_exporter  -w /go/src/github.com/jenningsloy318/netapp_exporter docker.io/jenningsloy318/prom-builder  make rpm

clean:
	rm -rf $(BIN_DIR)


fmt:
	@echo ">> format code style"
	$(GOFMT) -w $$(find . -path ./vendor -prune -o -name '*.go' -print) 



.PHONY: all style check_license  build docker-build  rpm docker-rpm fmt 