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

FROM golang:1.10 as builder

RUN apt-get update && apt-get install -y flite-dev

ENV IMPORTPATH github.com/campoy/justforfunc/13-flite-cgo
RUN mkdir -p "/go/src/${IMPORTPATH}"
COPY . "/go/src/${IMPORTPATH}"
RUN mv "/go/src/${IMPORTPATH}/vendor" "/go/src/vendor"
RUN go install -tags flite "${IMPORTPATH}/backend"

FROM debian:jessie-slim
RUN apt-get update && apt-get install -y flite-dev
COPY --from=builder /go/bin/backend /bin/backend
ENTRYPOINT ["/bin/backend"]