// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"ipm/internal/ipm/config"

	"github.com/spf13/cobra"
)

// AddValidateCommand adds the validate command to the manager command.
//
// Parameters:
//   - managerCmd: The manager command to which the validate command will be added.
//   - configDir: The directory containing the configuration files for package managers.
//   - schemaFile: The path to the JSON schema file used for validation.
//
// This function performs the following steps:
//  1. Creates a new "validate" command.
//  2. Sets the command to validate JSON configuration files for package managers.
//  3. Adds the "validate" command to the manager command.
//
// Example usage:
//
//	managerCmd := &cobra.Command{Use: "manager"}
//	AddValidateCommand(managerCmd, "/path/to/configDir", "/path/to/schemaFile")
//
// This function is useful for validating the configuration files of package managers
// against a specified JSON schema. It ensures that the configuration files adhere to
// the defined schema and prints the validation results.
func AddValidateCommand(managerCmd *cobra.Command, configDir string, schemaFile string) {
	// Command to validate JSON files
	var validateCmd = &cobra.Command{
		Use:   "validate [managers...]",
		Short: "Validate package managers config files",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			config.ValidateConfigs(args, schemaFile, configDir)
		},
	}

	// Add the validate command to the manager command
	managerCmd.AddCommand(validateCmd)
}
