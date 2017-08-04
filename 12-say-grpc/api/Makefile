# Copyright 2017 Google Inc. All rights reserved.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to writing, software distributed
# under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
# CONDITIONS OF ANY KIND, either express or implied.
#
# See the License for the specific language governing permissions and
# limitations under the License.
.DEFAULT_GOAL := help
.PHONY: help

deps: ## Installs dependencies
	go get google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go

developer-setup-mac: ## Installs protoc for mac
	brew install protobuf

build: ## Builds the api
	protoc -I . say.proto --go_out=plugins=grpc:.

help: ## shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
