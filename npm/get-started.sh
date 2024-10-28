#!/bin/sh

# Get system architecture
ARCH=$(uname -m)
PLATFORM=$(uname -s)

# echo "Platform: $PLATFORM"
# echo "Architecture: $ARCH"

# Determine binary name based on platform and architecture
if [ "$PLATFORM" = "Darwin" ]; then
    BINARY_NAME="restack-get-started-darwin-amd64"
elif [ "$PLATFORM" = "Linux" ]; then
    if [ "$ARCH" = "x86_64" ]; then
        BINARY_NAME="restack-get-started-linux-amd64"
    fi
else
    echo "Unsupported platform"
    exit 1
fi

# Get the version from package.json
VERSION=$(node -p "require('$(dirname "$0")/package.json').version")
BINARY_URL="https://github.com/restackio/get-started/releases/download/v${VERSION}/${BINARY_NAME}"
BINARY_PATH="./${BINARY_NAME}"

# Download the binary
echo "Downloading from ${BINARY_URL}..."
if ! curl -L -o "$BINARY_PATH" "$BINARY_URL"; then
    echo "Error downloading binary"
    exit 1
fi

# Make binary executable
chmod +x "$BINARY_PATH"

# Execute the binary with better error handling
if ! "$BINARY_PATH" 2>&1; then
    echo "Error executing binary"
    exit 1
fi
