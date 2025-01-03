// Package config provides utilities for managing configuration files
package config

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

// ListManagers lists package managers based on flags.
//
// Parameters:
//   - configDir: The directory where the configuration files are stored.
//   - all: A boolean flag indicating whether to list all package managers.
//   - enabled: A boolean flag indicating whether to list only enabled package managers.
//   - disabled: A boolean flag indicating whether to list only disabled package managers.
//
// Example usage:
//
//	config.ListManagers("/path/to/config/dir", true, false, false)  // List all package managers
//	config.ListManagers("/path/to/config/dir", false, true, false)  // List only enabled package managers
//	config.ListManagers("/path/to/config/dir", false, false, true)  // List only disabled package managers
//
// This function performs the following steps:
//  1. Uses a wildcard to get all JSON files in the specified directory.
//  2. Checks if each configuration file exists.
//  3. Reads the content of each configuration file.
//  4. Unmarshals the JSON data into a CommandConfig struct.
//  5. Lists the package manager names based on the provided flags.
func ListManagers(configDir string, all bool, enabled bool, disabled bool) {
	// Use wildcard to get all JSON files in the specified directory
	managerFiles, err := filepath.Glob(filepath.Join(configDir, "*.json"))
	if err != nil {
		log.Fatalf("Failed to read manager files: %v", err)
	}

	// Iterate over each manager file
	for _, managerFile := range managerFiles {
		// Check if the config file exists
		checkConfigFileExists(managerFile)

		// Read the config file
		data := ReadConfigFile(managerFile)

		// Unmarshal the config data
		config := UnmarshalConfig(data, managerFile)

		// Get the manager name by trimming the file extension
		managerName := strings.TrimSuffix(filepath.Base(managerFile), ".json")

		// List the manager name based on the provided flags
		if all || (enabled && config.Enabled) || (disabled && !config.Enabled) {
			fmt.Println(managerName)
		}
	}
}
