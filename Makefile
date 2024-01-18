# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# If you update this file, please follow
# https://www.thapaliya.com/en/writings/well-documented-makefiles

# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

########################################################################################################################
#                                                       ENV-VARS                                                       #
########################################################################################################################

# **************************************************** DIRECTORIES *****************************************************

# Golang version to use for building the project
GO_VERSION         ?= 1.21.5
GO_CONTAINER_IMAGE ?= docker.io/library/golang:$(GO_VERSION)

# Use GOPROXY environment variable if set
GOPROXY := $(shell go env GOPROXY)
ifeq ($(GOPROXY),)
GOPROXY := https://proxy.golang.org
endif
export GOPROXY

# Full directory of where the Makefile resides
BIN_DIR := bin

# Set build time variables including version details
LDFLAGS := $(shell hack/version.sh)

# VNG-CLOUD based on OpenStack infrastructure and Kamaji control plane
ALL_MANAGERS = core kubeadm-bootstrap


########################################################################################################################
#                                                       BINARIES                                                       #
########################################################################################################################

# ******************************************************* BUILD ********************************************************

all: test managers clusterctl

.PHONY: managers
managers: $(addprefix manager-,$(ALL_MANAGERS)) ## Run all manager-* targets

.PHONY: manager-core
manager-core: ## Build the core manager binary into the ./bin folder
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/manager sigs.k8s.io/cluster-api

.PHONY: manager-kubeadm-bootstrap
manager-kubeadm-bootstrap: ## Build the kubeadm bootstrap manager binary into the ./bin folder
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/kubeadm-bootstrap-manager sigs.k8s.io/cluster-api/bootstrap/kubeadm

.PHONY: clusterctl
clusterctl: ## Build the clusterctl binary
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/clusterctl sigs.k8s.io/cluster-api/cmd/clusterctl


########################################################################################################################
#                                                        TESTING                                                       #
########################################################################################################################

.PHONY: test
test: $(SETUP_ENVTEST) ## Run unit and integration tests
	KUBEBUILDER_ASSETS="$(KUBEBUILDER_ASSETS)" go test ./... $(TEST_ARGS)


########################################################################################################################
#                                                        HELPERS                                                       #
########################################################################################################################

go-version: ## Print the go version we use to compile our binaries and images
	@echo $(GO_VERSION)