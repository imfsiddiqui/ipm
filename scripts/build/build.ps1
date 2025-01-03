# Define the target platforms
# This array contains a list of target platforms for which the binary will be built.
# Each element is a hashtable with two keys:
#   - GOOS: The target operating system (e.g., "windows", "linux", "darwin")
#   - GOARCH: The target architecture (e.g., "386", "amd64")
$targets = @(
    @{ GOOS = "windows"; GOARCH = "386" },
    @{ GOOS = "windows"; GOARCH = "amd64" },
    @{ GOOS = "linux"; GOARCH = "386" },
    @{ GOOS = "linux"; GOARCH = "amd64" },
    @{ GOOS = "darwin"; GOARCH = "amd64" }
)

# Function to build the binary for a specified OS and architecture
function Invoke-BinaryBuild {
    param (
        [string]$GOOS,      # The target operating system
        [string]$GOARCH,    # The target architecture
        [string]$outputDir, # The output directory for the built binary
        [string]$archiveDir # The output directory for the archive
    )

    Write-Host "Building for $GOOS/$GOARCH..."
    
    # Set environment variables for cross-compilation
    $env:GOOS = $GOOS
    $env:GOARCH = $GOARCH
    
    # Define the binary name, adding .exe extension for Windows
    $binaryName = "ipm"
    if ($GOOS -eq "windows") {
        $binaryName = "$binaryName.exe"
    }
    
    # Build the binary using the Go compiler
    go build -o "$outputDir/$binaryName" $mainPackage
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Failed to build for $GOOS/$GOARCH"
        exit 1
    }

    # Copy the config directory to the output directory
    Write-Host "Copying config directory to $outputDir..."
    Copy-Item -Recurse -Path "config" -Destination "$outputDir/config"

    # Create the archive directory if it does not exist
    if (-not (Test-Path -Path $archiveDir)) {
        [void](New-Item -ItemType Directory -Path $archiveDir)
    }

    # Create the archive
    if ($GOOS -eq "windows") {
        $archiveName = "$archiveDir/ipm-$GOOS-$GOARCH.zip"
        Write-Host "Creating zip archive $archiveName..."
        # Use the Compress-Archive cmdlet to create the ZIP archive
        Compress-Archive -Path "$outputDir/*" -DestinationPath $archiveName
    } else {
        $archiveName = "$archiveDir/ipm-$GOOS-$GOARCH.tar.gz"
        Write-Host "Creating tar.gz archive $archiveName..."
        # Use the tar command to create the TAR.GZ archive
        tar -czf $archiveName -C $outputDir .
    }
}

# Define the output binary name
Write-Host "Defining binary name..."
$binaryName = "ipm"

# Define the main package
Write-Host "Defining main package..."
$mainPackage = "cmd/ipm/main.go"

# Clear the dist directory if it exists
Write-Host "Clearing dist directory..."
if (Test-Path -Path "dist") {
    Remove-Item -Recurse -Force "dist"
}

# Create the dist/bin directory
Write-Host "Creating dist/bin directory..."
$outputBinDir = "./dist/bin"
[void](New-Item -ItemType Directory -Path $outputBinDir)

# Define the archive directory
$archiveDir = "./dist/release"

# Build and archive for each target platform
# This loop iterates over each target platform defined in the $targets array.
# For each platform, it sets the GOOS and GOARCH environment variables,
# creates the output directory, and calls the Invoke-BinaryBuild function.
foreach ($target in $targets) {
    # Extract the target operating system and architecture from the current hashtable
    $GOOS = $target.GOOS
    $GOARCH = $target.GOARCH

    # Define the output directory for the built binary
    # The directory name includes the target operating system and architecture
    $outputDir = "$outputBinDir/ipm-$GOOS-$GOARCH"

    # Create the output directory if it does not exist
    [void](New-Item -ItemType Directory -Path $outputDir)

    # Call the Invoke-BinaryBuild function to build the binary and create the archive
    # Parameters:
    #   - GOOS: The target operating system
    #   - GOARCH: The target architecture
    #   - outputDir: The output directory for the built binary
    #   - archiveDir: The output directory for the archive
    Invoke-BinaryBuild -GOOS $GOOS -GOARCH $GOARCH -outputDir $outputDir -archiveDir $archiveDir
}

# Print completion messages
# These messages indicate that the build and archive process has completed,
# and provide the locations of the built binaries and archives.
Write-Host "Build and archive process completed."
Write-Host "Binaries are located in the '$outputBinDir' directory."
Write-Host "Archives are located in the '$archiveDir' directory."
