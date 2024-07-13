#!/bin/bash

go_required="go1.22.2 darwin/arm64"
go_version=$(go version)
if [[ $go_version != *$go_required* ]]; then
    echo "❌ Go version $go_version is not supported. Please install $go_required."
    return 1
  else
    echo "✅ Go version $go_required is installed."
fi

goose_required="v3.20.0"
goose_version=$(goose --version)
if [[ $goose_version != *$goose_required* ]]; then
    echo "📝 Installing Goose..."
    go install github.com/pressly/goose/v3/cmd/goose@v3.20.0
fi

echo "✅ Goose version $goose_required is installed."

bun_required="1.0.19"
bun_version=$(bun --version)
if [[ $bun_version != *$bun_required* ]]; then
    echo "❌ Bun version $bun_version is not supported. Please install $bun_version."
    return 1
  else
    echo "✅ Bun version $bun_required is installed."
fi
