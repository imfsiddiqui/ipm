// Package config provides utilities for managing configuration files
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"ipm/internal/ipm/utils"
)

// ValidateConfigs validates the configuration files.
//
// Parameters:
//   - args: A slice of strings containing the names of the package managers to validate.
//   - schemaFile: The path to the JSON schema file used for validation.
//   - configDir: The directory where the configuration files are stored.
//
// This function performs the following steps:
//  1. If specific package manager names are provided in args, it validates only those configuration files.
//  2. If no package manager names are provided, it validates all configuration files in the configDir.
//  3. For each configuration file, it checks if the file exists.
//  4. If the file exists, it validates the file against the provided JSON schema.
//  5. It prints the validation result for each configuration file.
func ValidateConfigs(args []string, schemaFile string, configDir string) {
	// Check if specific package manager names are provided in args
	if len(args) > 0 {
		// Iterate over each package manager name provided in args
		for _, managerName := range args {
			// Construct the path to the configuration file for the specified package manager
			configFile := filepath.Join(configDir, managerName+".json")

			// Check if the configuration file exists
			if _, err := os.Stat(configFile); os.IsNotExist(err) {
				// Print a message if the configuration file does not exist
				fmt.Printf("Validation failed for %s: config file does not exist\n", configFile)
				continue
			}

			// Validate the configuration file against the provided JSON schema
			if err := utils.ValidateJSONFile(schemaFile, configFile); err == nil {
				// Print a message if the validation is successful
				fmt.Printf("Validation successful for %s\n", configFile)
			} else {
				// Print a message if there is a validation error
				fmt.Printf("Validation error: %v\n", err)
			}
		}
	} else {
		// Validate all configuration files in the configDir if no specific package manager names are provided
		if err := utils.ValidateJSONFiles(schemaFile, configDir); err == nil {
			// Print a message if the validation is successful for all configuration files
			fmt.Println("Validation successful for all configs")
		} else {
			// Print the validation errors if any
			fmt.Printf("%v", err)
		}
	}
}
