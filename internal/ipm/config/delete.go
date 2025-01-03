// Package config provides utilities for managing configuration files
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// DeleteManagerConfig deletes a package manager config.
//
// Parameters:
//   - managerName: The name of the package manager whose config is to be deleted.
//   - configDir: The directory where the configuration files are stored.
//
// Example usage:
//
//	config.DeleteManagerConfig("apt", "/path/to/config/dir")
//
// This function performs the following steps:
//  1. Constructs the path to the configuration file for the specified package manager.
//  2. Checks if the configuration file exists.
//  3. Deletes the configuration file.
//  4. Prints a message indicating that the package manager config has been deleted.
func DeleteManagerConfig(managerName string, configDir string) {
	// Construct the path to the configuration file
	configFile := filepath.Join(configDir, managerName+".json")

	// Check if the config file exists
	checkConfigFileExists(configFile)

	// Delete the configuration file
	if err := os.Remove(configFile); err != nil {
		fmt.Printf("Failed to delete %s: %v", configFile, err)
	}

	// Print a message indicating that the package manager config has been deleted
	fmt.Printf("Manager %s config has been deleted", managerName)
}
