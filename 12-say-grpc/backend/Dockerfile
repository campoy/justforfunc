# Copyright 2107 Google Inc. All rights reserved.
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

FROM golang:1.8 as builder

RUN apt-get update && apt-get install -y flite-dev
RUN mkdir -p /go/src/app/vendor

ADD main.go /go/src/app
ADD vendor /go/src/app/vendor
RUN go install app

FROM debian:jessie-slim
RUN apt-get update && apt-get install -y flite-dev
RUN apt-get remove bash
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]