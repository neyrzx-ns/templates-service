#!/usr/bin/env sh

set -e

MIGRATION_DIR="${PWD}/migrations/"

docker run --rm --user "$(id -u)" --volume "${MIGRATION_DIR}":"${MIGRATION_DIR}" --workdir "${MIGRATION_DIR}" ryme/goose:latest \
  create "$1" sql
