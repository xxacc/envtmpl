# The name of the project
NAME := $(shell basename $(CURDIR))

# The architecture to build for
GOARCH ?= amd64

ALL_GOARCH := amd64

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe)

DIRS := $(addprefix bin/, $(GOARCH))

# Go related variables
GO ?= $(shell which go)
GOFMT ?= $(shell which gofmt)
GOLINT ?= $(shell which golint)
GOBUILDFLAGS ?= -installsuffix "static"
GOTESTFLAGS ?= -cover -timeout 30s
LINTTOOL := $(BIN_DIR)/revive

# List of packages to include in the binaries
PKGS := $(shell go list ./...)
# List of go files, excluding vendor
FILES ?= $(shell find . -name '*.go' -not -wholename './vendor/*' -print)

BIN_DIR := $(GOPATH)/bin

all: lint test build

build-%:
	@$(MAKE) --no-print-directory ARCH=$* build

all-build: $(addprefix build-, $(ALL_GOARCH))

build: bin/$(GOARCH)/$(NAME)

bin/$(GOARCH)/$(NAME): $(FILES)
	@echo "building: $@"
	CGO_ENABLED=0 \
	GO111MODULE=on \
	$(GO) build -v -o $@ $(GOBUILDFLAGS) .

.PHONY: install
install: build
	cp bin/$(GOARCH)/$(NAME) $(GOPATH)/bin

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: dirs
dirs: $(DIRS)
$(DIRS):
	mkdir -p $@
	ls $@

.PHONY: test
test:
	@$(GO) test $(GOBUILDFLAGS) $(GOTESTFLAGS) $(PKGS)

.PHONY: lint
lint: CGO_ENABLED := 0
lint:
	@echo "Format:"
	@$(GOFMT) -l $(FILES)
	@echo "\nLint:"
	@$(GOLINT) $(PKGS)
	@echo "\nVet:"
	@$(GO) vet $(PKGS)

.PHONY: clean
clean:
	rm -rf .go bin
