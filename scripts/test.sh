#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

cd $(dirname $0)/..

skipIntegrationTests=false

function helpmenu() {
  echo "Test script:
Use ./test.sh [flags]
Flags:
--skip-integration-tests: skip integration tests (not recommended)"
  exit 1
}

if [ "$#" -ne 0 ] && [ $1 == "help" ]; then
  helpmenu
fi

while [ ! $# -eq 0 ]
do
    case "$1" in
        --help | -h)
            helpmenu
            ;;
        --skip-integration-tests)
            skipIntegrationTests=true
            ;;
    esac
    shift
done

# function checkWorkingDir() {
#   if [ $(git status --short | wc -l) -gt 0 ]; then
#     echo "You have uncommited changes."
#     git status --short
#   fi
# }

function runTests() {
  richgo test $(go list ./... | grep -v /integration$) -race
  if [[ $skipIntegrationTests != true ]] ; then
    # Run integration tests here
    echo "Running integration tests"
  fi
}

function run() {
#   CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
#   LAST_TAG="$(git describe HEAD --tags --abbrev=0 2> /dev/null)" || true

  runTests
  echo "All tests and checks necessary for release passed."
}

run