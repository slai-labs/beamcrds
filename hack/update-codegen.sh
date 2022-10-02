#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG="${HOME}/go/pkg/mod/k8s.io/code-generator@v0.25.2/"
GO_PKG="slai.io/beamcrds/pkg"

# usage: generate-groups.sh <generators> <output-package> <apis-package> <groups-versions>
bash "${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,informer,lister" \
  ${GO_PKG}/beamrunner/v1 \
  ${GO_PKG} \
  beamrunner:v1 \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

# move and cleanup generated code
cp -rf ${GO_PKG}/beamrunner/v1/* pkg/beamrunner/v1
rm -rf slai.io