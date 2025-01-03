// Package config provides utilities for managing configuration files
package config

import (
	"fmt"
	"path/filepath"
)

// SetManager enables or disables a package manager based on the provided status.
// It reads the configuration file, updates the enabled status, and writes the updated configuration back to the file.
//
// Parameters:
//   - managerName: The name of the package manager to enable or disable.
//   - configDir: The directory where the configuration files are stored.
//   - status: A boolean value indicating whether to enable (true) or disable (false) the package manager.
//
// Example usage:
//
//	config.SetManager("apt", "/path/to/config/dir", true)  // Enable the apt package manager
//	config.SetManager("apt", "/path/to/config/dir", false) // Disable the apt package manager
//
// This function performs the following steps:
//  1. Constructs the path to the configuration file for the specified package manager.
//  2. Checks if the configuration file exists.
//  3. Reads the content of the configuration file.
//  4. Unmarshals the JSON data into a CommandConfig struct.
//  5. Updates the enabled status of the package manager based on the provided status.
//  6. Marshals the updated CommandConfig struct back into JSON data.
//  7. Writes the updated JSON data back to the configuration file.
//  8. Prints a message indicating whether the package manager has been enabled or disabled.
func SetManager(managerName string, configDir string, status bool) {
	configFile := filepath.Join(configDir, managerName+".json")

	// Check if the config file exists
	checkConfigFileExists(configFile)

	// Read the config file
	data := ReadConfigFile(configFile)

	// Unmarshal the config data
	config := UnmarshalConfig(data, configFile)

	// Update the enabled status based on the provided status
	config.Enabled = status

	// Marshal the config data
	newData := marshalConfig(config, configFile)

	// Write the config file
	writeConfigFile(configFile, newData)

	// Print a message indicating whether the package manager has been enabled or disabled
	if status {
		fmt.Printf("Manager %s has been enabled\n", managerName)
	} else {
		fmt.Printf("Manager %s has been disabled\n", managerName)
	}
}
