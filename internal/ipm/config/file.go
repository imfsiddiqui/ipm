// Package config provides utilities for managing configuration files
package config

import (
	"log"
	"os"
)

// checkConfigFileExists checks if the specified configuration file exists.
// If the file does not exist, it logs a fatal error and terminates the program.
//
// Parameters:
//   - configFile: The path to the configuration file to check.
//
// Example usage:
//
//	checkConfigFileExists("/path/to/config.json")
//
// This function is typically used to ensure that a required configuration file
// is present before attempting to read or write to it. If the file is missing,
// the program will log an error message and exit, preventing further execution.
func checkConfigFileExists(configFile string) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalf("Config file %s does not exist", configFile)
	}
}
