// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"ipm/internal/ipm/config"

	"github.com/spf13/cobra"
)

// AddDisableCommand adds the disable command to the manager command.
//
// Parameters:
//   - managerCmd: The manager command to which the disable command will be added.
//   - configDir: The directory containing the configuration files for package managers.
//
// This function performs the following steps:
//  1. Creates a new "disable" command.
//  2. Sets the command to disable a specified package manager.
//  3. Adds the "disable" command to the manager command.
//
// Example usage:
//
//	managerCmd := &cobra.Command{Use: "manager"}
//	AddDisableCommand(managerCmd, "/path/to/configDir")
//
// This function is useful for disabling a specified package manager. It updates the
// configuration to mark the package manager as disabled and prints the result.
func AddDisableCommand(managerCmd *cobra.Command, configDir string) {
	// Command to disable a package manager
	var disableCmd = &cobra.Command{
		Use:   "disable [manager]",
		Short: "Disable a package manager",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			config.SetManager(args[0], configDir, false)
		},
	}

	// Add the disable command to the manager command
	managerCmd.AddCommand(disableCmd)
}
