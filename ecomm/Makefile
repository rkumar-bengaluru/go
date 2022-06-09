
check-%:
	@: $(if $(value $*),,$(error $* is undefined))

define getBuildInfo
	cat BUILD_INFO | grep -w $1 | cut -d'=' -f2
endef

PATHS = ./config


PACKAGE="github.com/rkumar-bengaluru/go/ecomm"
VERSION=1
VERSION_CLIENT=1
APP_NAME=ecomm.exe
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_BRANCH=$(shell git branch --show-current)
LDFLAGS= -X '${PACKAGE}/cmd.Version=${VERSION}' -X '${PACKAGE}/cmd.GitCommit=${GIT_COMMIT}' -X '${PACKAGE}/cmd.BuildTime=${BUILD_TIME}'
BUILD=go build -v -o ./bin/${APP_NAME} ./cmd
IMAGE_NAME=ecomm
IMAGE_TAG=${VERSION}
IMAGE_ID="${IMAGE_NAME}:${IMAGE_TAG}"

image-tag: 
	@echo "${VERSION}"

image-name: 	
	@echo "${IMAGE_NAME}:${tag}"


run:
	go run ./cmd
tests:
	@echo "Testing..."
	@go test -v -gcflags=-l -cover $(PATHS) -coverprofile=coverage.out
	@go tool cover -func=coverage.out

build:
	@echo "Building..."
	@$(BUILD)

build-docker-image:
	@echo "Building docker image..."
	@echo $(IMAGE_TAG)
	@echo $(IMAGE_ID)
	@echo $(APP_NAME)
	@docker build --no-cache -f ./deploy/Dockerfile -t $(IMAGE_ID) --build-arg APP_NAME=${APP_NAME} .

build-docker: build \
			build-docker-image

FILES = $(cmd date '+%a %b %d %H:%M:%S %Y')
all:
	@echo "Building..."
	@echo $(IMAGE_TAG)