WHAT := server client

PWD ?= $(shell pwd)

VERSION   ?= $(shell git describe --tags)
REVISION  ?= $(shell git rev-parse HEAD)
BRANCH    ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILDUSER ?= $(shell id -un)
BUILDTIME ?= $(shell date '+%Y%m%d-%H:%M:%S')

.PHONY: build build-darwin-amd64 build-linux-amd64 build-windows-amd64 clean release test

build:
	for target in $(WHAT); do \
		go build -ldflags "-X github.com/Selahattinn/go-system-agent/pkg/version.Version=${VERSION} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.Revision=${REVISION} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.Branch=${BRANCH} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/$$target ./cmd/$$target; \
	done

build-darwin-amd64:
	for target in $(WHAT); do \
		CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -a -installsuffix cgo -ldflags "-X github.com/Selahattinn/go-system-agent/pkg/version.Version=${VERSION} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.Revision=${REVISION} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.Branch=${BRANCH} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/go-system-agent-${VERSION}-linux-amd64/$$target ./cmd/$$target; \
	done

build-linux-amd64:
	for target in $(WHAT); do \
		CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -ldflags "-X github.com/Selahattinn/go-system-agent/pkg/version.Version=${VERSION} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.Revision=${REVISION} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.Branch=${BRANCH} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/go-system-agent-${VERSION}-linux-amd64/$$target ./cmd/$$target; \
	done

build-windows-amd64:
	for target in $(WHAT); do \
		CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -a -installsuffix cgo -ldflags "-X github.com/Selahattinn/go-system-agent/pkg/version.Version=${VERSION} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.Revision=${REVISION} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.Branch=${BRANCH} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/Selahattinn/go-system-agent/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/go-system-agent-${VERSION}-windows-amd64/$$target.exe ./cmd/$$target; \
	done

clean:
	rm -rf ./bin;

prepare-release: clean 
	$(info "Run: 'git commit' and 'git tag', then 'make release'")

release: build-darwin-amd64 build-linux-amd64 build-windows-amd64
	cp ${PWD}/LICENSE ${PWD}/bin/go-system-agent-${VERSION}-darwin-amd64
	cp ${PWD}/LICENSE ${PWD}/bin/go-system-agent-${VERSION}-linux-amd64
	cp ${PWD}/LICENSE ${PWD}/bin/go-system-agent-${VERSION}-windows-amd64
	cd ${PWD}/bin; tar cfvz go-system-agent-${VERSION}-darwin-amd64.tar.gz ./go-system-agent-${VERSION}-darwin-amd64
	cd ${PWD}/bin; tar cfvz go-system-agent-${VERSION}-linux-amd64.tar.gz ./go-system-agent-${VERSION}-linux-amd64
	cd ${PWD}/bin; tar cfvz go-system-agent-${VERSION}-windows-amd64.tar.gz ./go-system-agent-${VERSION}-windows-amd64

test:
	go test ./...
