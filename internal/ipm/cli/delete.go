// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"ipm/internal/ipm/config"

	"github.com/spf13/cobra"
)

// AddDeleteCommand adds the delete command to the manager command.
//
// Parameters:
//   - managerCmd: The manager command to which the delete command will be added.
//   - configDir: The directory containing the configuration files for package managers.
//
// This function performs the following steps:
//  1. Creates a new "delete" command.
//  2. Sets the command to delete a specified package manager configuration file.
//  3. Adds the "delete" command to the manager command.
//
// Example usage:
//
//	managerCmd := &cobra.Command{Use: "manager"}
//	AddDeleteCommand(managerCmd, "/path/to/configDir")
//
// This function is useful for deleting a configuration file for a specified
// package manager. It removes the configuration file from the specified directory
// and prints the result.
func AddDeleteCommand(managerCmd *cobra.Command, configDir string) {
	// Command to delete a package manager config
	var deleteCmd = &cobra.Command{
		Use:   "delete [manager]",
		Short: "Delete a package manager config",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			config.DeleteManagerConfig(args[0], configDir)
		},
	}

	// Add the delete command to the manager command
	managerCmd.AddCommand(deleteCmd)
}
