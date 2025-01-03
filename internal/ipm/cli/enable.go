// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"ipm/internal/ipm/config"

	"github.com/spf13/cobra"
)

// AddEnableCommand adds the enable command to the manager command.
//
// Parameters:
//   - managerCmd: The manager command to which the enable command will be added.
//   - configDir: The directory containing the configuration files for package managers.
//
// This function performs the following steps:
//  1. Creates a new "enable" command.
//  2. Sets the command to enable a specified package manager.
//  3. Adds the "enable" command to the manager command.
//
// Example usage:
//
//	managerCmd := &cobra.Command{Use: "manager"}
//	AddEnableCommand(managerCmd, "/path/to/configDir")
//
// This function is useful for enabling a specified package manager. It updates the
// configuration to mark the package manager as enabled and prints the result.
func AddEnableCommand(managerCmd *cobra.Command, configDir string) {
	// Command to enable a package manager
	var enableCmd = &cobra.Command{
		Use:   "enable [manager]",
		Short: "Enable a package manager",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			config.SetManager(args[0], configDir, true)
		},
	}

	// Add the enable command to the manager command
	managerCmd.AddCommand(enableCmd)
}
