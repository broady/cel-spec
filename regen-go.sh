#!/bin/bash
#
# Copyright 2018 Google Inc.
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

# This script rebuilds the generated code for the protocol buffers.
# To run this you will need protoc and goprotobuf installed;
# see https://github.com/golang/protobuf for instructions.
# You also need Go and Git installed.

set -euo pipefail

#CEL_REPO=https://github.com/google/cel-spec
CEL_REPO=https://github.com/broady/cel-spec
PROTO_REPO=https://github.com/google/protobuf
GOOGLEAPIS_REPO=https://github.com/googleapis/googleapis

function die() {
  echo 1>&2 $*
  exit 1
}

# Sanity check that the right tools are accessible.
for tool in go git protoc protoc-gen-go; do
  q=$(which $tool) || die "didn't find $tool"
  echo 1>&2 "$tool: $q"
done

mkdir -p tmp
if [[ -z "${NOCLEANTMP:-}" ]]; then
  trap 'rm -rf tmp' EXIT
fi

test -d tmp/spec       || git clone --depth=1 $CEL_REPO tmp/spec
test -d tmp/protobuf   || git clone --depth=1 $PROTO_REPO tmp/protobuf
test -d tmp/googleapis || git clone --depth=1 $GOOGLEAPIS_REPO tmp/googleapis

#ROOT=tmp/spec
ROOT=.
IMPORTS="-I tmp/googleapis/ -I tmp/protobuf/src/ -I $ROOT"
OUTFLAG="--go_out=paths=source_relative:."
protoc $OUTFLAG $ROOT/proto/v1/*.proto $IMPORTS
#protoc $OUTFLAG $ROOT/proto/checked/v1/*.proto $IMPORTS
