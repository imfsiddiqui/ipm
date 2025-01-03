// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"ipm/internal/ipm/config"

	"github.com/spf13/cobra"
)

// AddGenerateCommand adds the generate command to the manager command.
//
// Parameters:
//   - managerCmd: The manager command to which the generate command will be added.
//   - configDir: The directory containing the configuration files for package managers.
//
// This function performs the following steps:
//  1. Creates a new "generate" command.
//  2. Sets the command to generate a new package manager configuration file.
//  3. Adds the "generate" command to the manager command.
//
// Example usage:
//
//	managerCmd := &cobra.Command{Use: "manager"}
//	AddGenerateCommand(managerCmd, "/path/to/configDir")
//
// This function is useful for generating a new configuration file for a specified
// package manager. It creates a new configuration file in the specified directory
// and prints the result.
func AddGenerateCommand(managerCmd *cobra.Command, configDir string) {
	// Command to generate a new package manager config
	var generateCmd = &cobra.Command{
		Use:   "generate [manager]",
		Short: "Generate a new package manager config",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			config.GenerateManagerConfig(args[0], configDir)
		},
	}

	// Add the generate command to the manager command
	managerCmd.AddCommand(generateCmd)
}
