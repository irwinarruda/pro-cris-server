#!/bin/bash

source ~/.zshrc

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

node_required="v22.0.0"
node_version=$(node --version)
if [[ $node_version != *$node_required* ]]; then
    echo "📝 Installing Node version $node_required..."
    nvm install $node_required
    nvm use $node_required
fi
echo "✅ Node version $node_required is installed."

dotenv_required="[--help] [--debug]"
dotenv_version=$(dotenv)
if [[ $dotenv_version != *$dotenv_required* ]]; then
    echo "📝 Installing Dotenv..."
    npm install -g dotenv-cli@latest
fi
echo "✅ Dotenv installed."

echo "📝 Updating external dependencies..."
cd external/watch && npm install
