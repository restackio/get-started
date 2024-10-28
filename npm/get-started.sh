#!/bin/sh

# Get system architecture
ARCH=$(uname -m)
PLATFORM=$(uname -s)

echo "Platform: $PLATFORM"
echo "Architecture: $ARCH"

# Determine binary path based on platform and architecture
if [ "$PLATFORM" = "Darwin" ]; then
    BINARY_PATH="./restack-get-started-darwin-amd64"
elif [ "$PLATFORM" = "Linux" ]; then
    if [ "$ARCH" = "x86_64" ]; then
        BINARY_PATH="./restack-get-started-linux-amd64"
    fi
else
    echo "Unsupported platform"
    exit 1
fi

echo "Binary path: $BINARY_PATH"

# Check if binary exists
if [ ! -f "$BINARY_PATH" ]; then
    echo "Error: Binary not found at $BINARY_PATH"
    exit 1
fi

# Make binary executable
chmod +x "$BINARY_PATH"

# Execute the binary with better error handling
if ! "$BINARY_PATH" 2>&1; then
    echo "Error executing binary"
    exit 1
fi
