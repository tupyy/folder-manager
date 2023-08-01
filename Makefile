MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_PATH := $(patsubst %/,%,$(dir $(MKFILE_PATH)))
# The details of the application:
binary:=folder-manager

LOCAL_BIN_PATH := ${PROJECT_PATH}/bin
GOFMT := gofmt

OPENAPI_GENERATOR ?= ${LOCAL_BIN_PATH}/openapi-generator
NPM ?= "$(shell which npm)"
openapi-generator:
ifeq (, $(shell which ${NPM} 2> /dev/null))
	@echo "npm is not available please install it to be able to install openapi-generator"
	exit 1
endif
ifeq (, $(shell which ${LOCAL_BIN_PATH}/openapi-generator 2> /dev/null))
	@{ \
	set -e ;\
	mkdir -p ${LOCAL_BIN_PATH} ;\
	mkdir -p ${LOCAL_BIN_PATH}/openapi-generator-installation ;\
	cd ${LOCAL_BIN_PATH} ;\
	${NPM} install --prefix ${LOCAL_BIN_PATH}/openapi-generator-installation @openapitools/openapi-generator-cli@cli-4.3.1 ;\
	ln -s openapi-generator-installation/node_modules/.bin/openapi-generator openapi-generator ;\
	}
endif

#####################
# Help targets      #
#####################

#help help.highlevel: show help for high level targets. Use 'make help.all' to display all help messages
help.highlevel:
	@grep -hE '^[a-z_-]+:' $(MAKEFILE_LIST) | LANG=C sort -d | \
	awk 'BEGIN {FS = ":"}; {printf("$(COLOR_YELLOW)%-25s$(RESET_COLOR) %s\n", $$1, $$2)}'
.PHONY: help.highlevel

#help help.all: display all targets' help messages
help.all:
	@grep -hE '^#help|^[a-z_-]+:' $(MAKEFILE_LIST) | sed "s/#help //g" | LANG=C sort -d | \
	awk 'BEGIN {FS = ":"}; {if ($$1 ~ /\./) printf("    $(COLOR_BLUE)%-21s$(RESET_COLOR) %s\n", $$1, $$2); else printf("$(COLOR_YELLOW)%-25s$(RESET_COLOR) %s\n", $$1, $$2)}'
.PHONY: help.all

#help run go mod vendor
build.vendor:
	@rm -fr $(CURDIR)/vendor
	go mod tidy
	go mod vendor
.PHONY: build.vendor

#help build locally a binary, in target/ folder
build.local: ## Build local
	go build -mod=vendor $(BUILD_ARGS) -ldflags "-X main.CommitID=$(GIT_COMMIT) -s -w" \
	-o $(CURDIR)/target/run $(CURDIR)/cli/main.go
.PHONY: build.local

#help build.docker: build docker image
build.docker:
	DOCKER_BUILDKIT=1 docker build --no-cache=true --build-arg build_args="$(BUILD_ARGS)" -t $(IMAGE_NAME):$(IMAGE_TAG) -f Dockerfile .
.PHONY: build.docker

#help openapi.generate: Generate openapi schema
openapi.generate: openapi-generator ## Generate the openapi schema and generated package
	rm -rf pkg/api/v1
	$(OPENAPI_GENERATOR) validate -i openapi/photoz.yaml
	$(OPENAPI_GENERATOR) generate -i openapi/photoz.yaml -g go -o pkg/api/v1 --package-name v1 -t openapi/templates --ignore-file-override ./.openapi-generator-ignore
	$(GOFMT) -w pkg/api/v1
.PHONY: openapi/generate

BASE_CONNSTR="postgresql://$(PG_USER):$(PG_PWD)@$(DB_HOST):$(DB_PORT)"
GEN_CMD=$(TOOLS_DIR)/gen --sqltype=postgres \
	--module=github.com/tupyy/folder-manager/internal/repository/postgres/models --exclude=schema_migrations \
	--gorm --no-json --no-xml --overwrite --out $(CURDIR)/internal/repository/postgres/models

#help generate.models: generate models for the database
generate.models:
	sh -c '$(GEN_CMD) --connstr "$(BASE_CONNSTR)/photoz?sslmode=disable"  --model=pg --database photoz' 						# Generate models for the DB tables
.PHONY: generate.models

#help: test: run tests with ginko
test: ginkgo
	$(GINKGO) -focus=$(FOCUS) -v --cover --coverprofile=cover.out ./...
.PHONY: test

GINKGO = $(shell pwd)/bin/ginkgo

#help ginko: Download ginko locally
ginkgo:
ifeq (, $(shell which ginkgo 2> /dev/null))
	$(call go-install-tool,$(GINKGO),github.com/onsi/ginkgo/v2/ginkgo@v2.1.3)
endif
.PHONY: ginko

# include infra targets.
-include ./infra.mk
