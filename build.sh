#!/usr/bin/env bash

# This script is used to quickly create all binaries, for both macOS and Windows
ARCHS_OSX=("amd64" "arm64")
ARCHS_WIN=("amd64" "arm64")
ARCHS_LINUX=("amd64" "arm64")

mkdir -p target

echo "Building macOS..."
for ARCH in "${ARCHS_OSX[@]}"; do
  echo "Building for darwin/$ARCH..."
  GOOS=darwin GOARCH=$ARCH go build -o "target/latest-macOS-$ARCH"
done

echo "Building Linux..."
for ARCH in "${ARCHS_LINUX[@]}"; do
  echo "Building for linux/$ARCH..."
  GOOS=linux GOARCH=$ARCH go build -o "target/latest-linux-$ARCH.exe"
done

echo "Building Windows..."
for ARCH in "${ARCHS_WIN[@]}"; do
  echo "Building for windows/$ARCH..."
  GOOS=windows GOARCH=$ARCH go build -o "target/latest-windows-$ARCH.exe"
done

echo "Finished!"
