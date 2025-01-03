// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"ipm/internal/ipm/config"

	"github.com/spf13/cobra"
)

// AddListCommand adds the list command to the manager command.
//
// Parameters:
//   - managerCmd: The manager command to which the list command will be added.
//   - configDir: The directory containing the configuration files for package managers.
//
// This function performs the following steps:
//  1. Creates a new "list" command.
//  2. Sets the command to list package managers based on the specified flags.
//  3. Adds the "list" command to the manager command.
//
// Example usage:
//
//	managerCmd := &cobra.Command{Use: "manager"}
//	AddListCommand(managerCmd, "/path/to/configDir")
//
// This function is useful for listing package managers. It provides options to list
// all package managers, only enabled package managers, or only disabled package managers.
func AddListCommand(managerCmd *cobra.Command, configDir string) {
	// Command to list package managers
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List package managers",
		Run: func(cmd *cobra.Command, args []string) {
			all, _ := cmd.Flags().GetBool("all")
			enabled, _ := cmd.Flags().GetBool("enabled")
			disabled, _ := cmd.Flags().GetBool("disabled")

			// If no flag is passed, return help output
			if !all && !enabled && !disabled {
				cmd.Help()
				return
			}

			config.ListManagers(configDir, all, enabled, disabled)
		},
	}

	// Add flags to the list command
	listCmd.Flags().BoolP("all", "a", false, "List all package managers")
	listCmd.Flags().BoolP("enabled", "e", false, "List enabled package managers")
	listCmd.Flags().BoolP("disabled", "d", false, "List disabled package managers")

	// Add the list command to the manager command
	managerCmd.AddCommand(listCmd)
}
