GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test
BINARY_NAME_AKO_GATEWAY=ako-gateway
PACKAGE_PATH_AKO=github.com/vmware/load-balancer-and-ingress-services-for-kubernetes
REL_PATH_AKO_GW=$(PACKAGE_PATH_AKO)/cmd/ako-gateway


define GetSupportabilityMatrix
$(shell node -p "require('./buildsettings.json').$(1)")
endef
# TODO(Akshay): This has to be addressed
AVI_MIN_VERSION:=$(call GetSupportabilityMatrix,avi.minVersion)
AVI_MAX_VERSION:=$(call GetSupportabilityMatrix,avi.maxVersion)
K8S_MIN_VERSION:=$(call GetSupportabilityMatrix,kubernetes.minVersion)
K8S_MAX_VERSION:=$(call GetSupportabilityMatrix,kubernetes.maxVersion)
AKO_VERSION:=v$(call GetSupportabilityMatrix,version)
AKO_LDFLAGS:="-X 'main.version=$(AKO_VERSION)' \
		-X '$(PACKAGE_PATH_AKO)/internal/lib.aviMinVersion=$(AVI_MIN_VERSION)' \
		-X '$(PACKAGE_PATH_AKO)/internal/lib.aviMaxVersion=$(AVI_MAX_VERSION)' \
		-X '$(PACKAGE_PATH_AKO)/internal/lib.k8sMinVersion=$(K8S_MIN_VERSION)' \
		-X '$(PACKAGE_PATH_AKO)/internal/lib.k8sMaxVersion=$(K8S_MAX_VERSION)'"

ifdef GOLANG_SRC_REPO
	BUILD_GO_IMG=$(GOLANG_SRC_REPO)
else
	BUILD_GO_IMG=golang:latest
endif

.PHONY: glob-vars
glob-vars:
	$(eval BUILD_ARG_AKO_LDFLAGS=--build-arg AKO_LDFLAGS=$(AKO_LDFLAGS))

ifndef BUILD_TAG
	$(eval BUILD_TAG=$(shell ./hack/jenkins/get_build_version.sh "dummy" 0))
endif

ifndef BUILD_TIME
	$(eval BUILD_TIME=$(shell date +%Y-%m-%d_%H:%M:%S_%Z))
endif

ifdef GOLANG_SRC_REPO
	$(eval BUILD_ARG_GOLANG=--build-arg golang_src_repo=$(GOLANG_SRC_REPO))
else
	$(eval BUILD_ARG_GOLANG=)
endif

ifdef PHOTON_SRC_REPO
	$(eval BUILD_ARG_PHOTON=--build-arg photon_src_repo=$(PHOTON_SRC_REPO))
else
	$(eval BUILD_ARG_PHOTON=)
endif

ifdef UBI_SRC_REPO
	$(eval BUILD_ARG_UBI=--build-arg ubi_src_repo=$(UBI_SRC_REPO))
else
	$(eval BUILD_ARG_UBI=)
endif

.PHONY: all
all: build docker

.PHONY: sync-crd-files
sync-crd-files:
		cp ./helm/ako/crds/* ./ako-operator/helm/ako-operator/crds/

.PHONY: pre-build
pre-build: sync-crd-files

# builds
.PHONY: build
build: pre-build glob-vars
		sudo docker run \
		-w=/go/src/$(PACKAGE_PATH_AKO) \
		-v $(PWD):/go/src/$(PACKAGE_PATH_AKO) $(BUILD_GO_IMG) \
		$(GOBUILD) \
		-o /go/src/$(PACKAGE_PATH_AKO)/bin/$(BINARY_NAME_AKO_GATEWAY) \
		-ldflags $(AKO_LDFLAGS) \
		-mod=vendor \
		/go/src/$(REL_PATH_AKO)


.PHONY: build-local
build-local: pre-build
		$(GOBUILD) \
		-o bin/$(BINARY_NAME_AKO_GATEWAY) \
		-ldflags $(AKO_LDFLAGS) \
		-mod=vendor \
		./cmd/ako-gateway

.PHONY: clean
clean:
		$(GOCLEAN) -mod=vendor $(REL_PATH_AKO)
		rm -f bin/$(BINARY_NAME_AKO_GATEWAY)

.PHONY: deps
deps:
	dep ensure -v

# docker images
.PHONY: docker
docker: glob-vars
	sudo docker build \
	-t $(BINARY_NAME_AKO_GATEWAY):latest \
	--label "BUILD_TAG=$(BUILD_TAG)" \
	--label "BUILD_TIME=$(BUILD_TIME)" \
	$(BUILD_ARG_GOLANG) $(BUILD_ARG_PHOTON) $(BUILD_ARG_AKO_LDFLAGS) \
	-f Dockerfile.ako-gateway .


# tests

# linting and formatting
GO_FILES := $(shell find . -type d -path ./vendor -prune -o -type f -name '*.go' -print)
.PHONY: fmt
fmt:
	@echo
	@echo "Formatting Go files"
	@gofmt -s -l -w $(GO_FILES)

.golangci-bin:
	@echo "Installing Golangci-lint"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $@ v1.50.1

.PHONY: golangci
golangci: .golangci-bin
	@echo "Running golangci"
	@GOOS=linux .golangci-bin/golangci-lint run -c .golangci.yml

.PHONY: golangci-fix
golangci-fix: .golangci-bin
	@echo "Running golangci-fix"
	@GOOS=linux .golangci-bin/golangci-lint run -c .golangci.yml --fix
