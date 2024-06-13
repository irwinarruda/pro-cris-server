#!/bin/bash

go_required="go1.22.2 darwin/arm64"
go_version=$(go version)
if [[ $go_version != *$go_required* ]]; then
    echo "❌ Go version $go_version is not supported. Please install $go_required"
  else
    echo "✅ Go version $go_required is installed"
fi

air_required="v1.51.0"
air_version=$(air -v)
if [[ $air_version != *$air_required* ]]; then
    echo "📝 Installing Air"
    go install github.com/cosmtrek/air@v1.51.0
fi

echo "✅ Air version $air_required is installed"

echo "📝 Installing Gow"
go install github.com/mitranim/gow@latest

echo "✅ Gow version is installed"

goose_required="v3.20.0"
goose_version=$(goose --version)
if [[ $goose_version != *$goose_required* ]]; then
    echo "📝 Installing Goose"
    go install github.com/pressly/goose/v3/cmd/goose@v3.20.0
fi

echo "✅ Goose version $goose_required is installed"

