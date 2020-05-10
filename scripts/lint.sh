#!/bin/bash

function runLint() {
  echo "Checking for unchecked errors."
  errcheck $(go list ./...)
  echo "Linting code."
  test -z "$(golint ./... | grep -v "^vendor" | tee /dev/stderr)"
  echo "Examining source code against code defect."
  go vet $(go list ./...)
  echo "Checking if code can be simplified or can be improved."
  staticcheck ./...
}

runLint