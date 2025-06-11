#!/bin/bash

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"
FILE="unknown"

case "$OS-$ARCH" in
    linux-x86_64)   FILE="gmfi-linux-amd" ;;
    linux-aarch64)  FILE="gmfi-linux-arm" ;;
    darwin-arm64)   FILE="gmfi-macos-arm" ;;
    *)
        echo "Unsupported OS/arch: $OS-$ARCH"
        exit 1
        ;;
esac

URL="https://github.com/jvqtil/gmfi/releases/download/gmfi/$FILE"
curl -L -O "$URL" || { echo "Download failed"; exit 1; }

chmod +x "$FILE"
echo "Successfully downloaded $FILE!"
echo "Moving to /usr/local/bin, may require password"
echo "If this fails please move the file to some bin directory in your PATH yourself"
sudo mv ./$FILE /usr/local/bin/
echo "Successfully moved to /usr/local/bin, now run with a simple *gmfi*"
