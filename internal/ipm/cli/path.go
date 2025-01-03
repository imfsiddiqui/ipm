// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"log"
	"os"
)

// GetExecutablePath gets the path of the executable and handles errors.
//
// This function retrieves the absolute path of the currently running executable.
// It is useful for determining the directory structure relative to the executable.
//
// Returns:
//   - string: The absolute path of the currently running executable.
//
// Example usage:
//
//	exePath := cli.GetExecutablePath()
//
// This function performs the following steps:
//  1. Calls os.Executable() to get the absolute path of the executable.
//  2. Checks for errors and logs a fatal error if os.Executable() fails.
//  3. Returns the absolute path of the executable.
func GetExecutablePath() string {
	// Call os.Executable() to get the absolute path of the executable
	exePath, err := os.Executable()
	if err != nil {
		// Log a fatal error if os.Executable() fails
		log.Fatalf("Failed to get executable path: %v", err)
	}
	// Return the absolute path of the executable
	return exePath
}
