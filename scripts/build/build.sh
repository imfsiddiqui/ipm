#!/bin/bash

# Define the target platforms
# This array contains a list of target platforms for which the binary will be built.
# Each element is a tuple with two values:
#   - GOOS: The target operating system (e.g., "windows", "linux", "darwin")
#   - GOARCH: The target architecture (e.g., "386", "amd64", "arm64")
TARGETS=(
    "windows 386"
    "windows amd64"
    "linux 386"
    "linux amd64"
    "darwin amd64"
    "darwin arm64"
)

# Function to build the binary for a specified OS and architecture
invoke_binary_build() {
    local GOOS=$1        # The target operating system
    local GOARCH=$2      # The target architecture
    local OUTPUT_DIR=$3  # The output directory for the built binary
    local ARCHIVE_DIR=$4 # The output directory for the archive

    echo "Building for $GOOS/$GOARCH..."

    # Set environment variables for cross-compilation
    export GOOS=$GOOS
    export GOARCH=$GOARCH

    # Define the binary name, adding .exe extension for Windows
    local BINARY_NAME=$BINARY_NAME
    if [ "$GOOS" == "windows" ]; then
        BINARY_NAME="$BINARY_NAME.exe"
    fi

    # Build the binary using the Go compiler
    go build -o "$OUTPUT_DIR/$BINARY_NAME" $MAIN_PACKAGE
    if [ $? -ne 0 ]; then
        echo "Failed to build for $GOOS/$GOARCH"
        exit 1
    fi

    # Copy the config directory to the output directory
    echo "Copying config directory to $OUTPUT_DIR..."
    cp -r config "$OUTPUT_DIR/config"

    # Create the archive directory if it does not exist
    mkdir -p $ARCHIVE_DIR

    # Create the archive
    if [ "$GOOS" == "windows" ]; then
        local ARCHIVE_NAME="$ARCHIVE_DIR/ipm-$GOOS-$GOARCH.zip"
        echo "Creating zip archive $ARCHIVE_NAME..."
        # Use the zip command to create the ZIP archive
        (cd $OUTPUT_DIR && zip -q -T -r ../../../$ARCHIVE_NAME .)
    else
        local ARCHIVE_NAME="$ARCHIVE_DIR/ipm-$GOOS-$GOARCH.tar.gz"
        echo "Creating tar.gz archive $ARCHIVE_NAME..."
        # Use the tar command to create the TAR.GZ archive
        tar -czf $ARCHIVE_NAME -C $OUTPUT_DIR .
    fi
}

# Define the output binary name
echo "Defining binary name..."
BINARY_NAME="ipm"

# Define the main package
echo "Defining main package..."
MAIN_PACKAGE="cmd/ipm/main.go"

# Clear the dist directory if it exists
echo "Clearing dist directory..."
if [ -d "dist" ]; then
    rm -rf "dist"
fi

# Create the dist/bin directory
echo "Creating dist/bin directory..."
OUTPUT_BIN_DIR="./dist/bin"
mkdir -p $OUTPUT_BIN_DIR

# Define the archive directory
ARCHIVE_DIR="./dist/release"

# Build and archive for each target platform
for TARGET in "${TARGETS[@]}"; do
    # Extract the target operating system and architecture from the current tuple
    GOOS=$(echo $TARGET | cut -d ' ' -f 1)
    GOARCH=$(echo $TARGET | cut -d ' ' -f 2)

    # Define the output directory for the built binary
    # The directory name includes the target operating system and architecture
    OUTPUT_DIR="$OUTPUT_BIN_DIR/ipm-$GOOS-$GOARCH"

    # Create the output directory if it does not exist
    mkdir -p $OUTPUT_DIR

    # Call the invoke_binary_build function to build the binary and create the archive
    invoke_binary_build $GOOS $GOARCH $OUTPUT_DIR $ARCHIVE_DIR
done

# Print completion messages
# These messages indicate that the build and archive process has completed,
# and provide the locations of the built binaries and archives.
echo "Build and archive process completed."
echo "Binaries are located in the '$OUTPUT_BIN_DIR' directory."
echo "Archives are located in the '$ARCHIVE_DIR' directory."
