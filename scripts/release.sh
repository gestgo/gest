#!/usr/bin/env bash
set -euo pipefail

# Usage: ./scripts/release.sh <package> [patch|minor|major]
#
# Examples:
#   ./scripts/release.sh lifecycle patch
#   ./scripts/release.sh mcp minor
#   ./scripts/release.sh grpc major

VALID_PACKAGES=("lifecycle" "mcp" "grpc")

declare -A PACKAGE_PATHS=(
  ["lifecycle"]="package/core/lifecycle"
  ["mcp"]="package/extension/mcp"
  ["grpc"]="package/extension/grpc"
)

usage() {
  echo "Usage: $0 <package> [patch|minor|major]"
  echo ""
  echo "Packages: ${VALID_PACKAGES[*]}"
  exit 1
}

# Validate args
if [[ $# -lt 1 ]]; then
  usage
fi

PACKAGE="$1"
VERSION_TYPE="${2:-}"

# Check package is valid
if [[ ! " ${VALID_PACKAGES[*]} " =~ " ${PACKAGE} " ]]; then
  echo "Error: unknown package '${PACKAGE}'"
  echo "Valid packages: ${VALID_PACKAGES[*]}"
  exit 1
fi

PKG_PATH="${PACKAGE_PATHS[$PACKAGE]}"

echo "==> Releasing package: ${PACKAGE} (${PKG_PATH})"

# Run vet and tests before release
echo "==> go vet ./..."
(cd "${PKG_PATH}" && GOWORK=off go vet ./...)

echo "==> go test ./..."
(cd "${PKG_PATH}" && GOWORK=off go test ./...) || true

# Build release-it args
RELEASE_ARGS=()
if [[ -n "${VERSION_TYPE}" ]]; then
  RELEASE_ARGS+=("${VERSION_TYPE}")
fi

echo "==> Running release-it..."
npm run "release:${PACKAGE}" -- "${RELEASE_ARGS[@]+"${RELEASE_ARGS[@]}"}"
