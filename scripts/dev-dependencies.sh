#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

function main() {
    passGoDevDependencies
}

function passGoDevDependencies() {
  echo "Installing developer tools."

  echo "goveralls https://github.com/mattn/goveralls"
  go get -u github.com/mattn/goveralls

  echo "errcheck https://github.com/kisielk/errcheck"
  go get -u github.com/kisielk/errcheck

  echo "golint golang.org/x/lint/golint"
  go get -u golang.org/x/lint/golint

  echo "staticcheck honnef.co/go/tools/cmd/staticcheck"
  go get -u honnef.co/go/tools/cmd/staticcheck

  echo "swag github.com/swaggo/swag/cmd/swag"
  go get -u github.com/swaggo/swag/cmd/swag

  echo "testify github.com/stretchr/testify"
  go get -u github.com/stretchr/testify

  echo "richgo github.com/kyoh86/richgo"
  go get -u github.com/kyoh86/richgo

  echo "reflex github.com/cespare/reflex"
  go get -u github.com/cespare/reflex
}

main