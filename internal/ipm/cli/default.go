// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"ipm/internal/ipm/config"
	"ipm/internal/ipm/manager"
	"ipm/internal/ipm/utils"
	"log"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
)

// SetupDefaultManagerCommands sets up the default manager commands based on the OS.
//
// Parameters:
//   - rootCmd: The root command to which the default manager commands will be added.
//   - configDir: The directory containing the configuration files for package managers.
//   - schemaFile: The path to the JSON schema file used for validation.
//   - argsLength: The length of the command-line arguments.
//   - firstArg: The first command-line argument.
//   - secondArg: The second command-line argument.
//
// This function performs the following steps:
//  1. Checks if the validate command is being run.
//  2. Detects the default package manager based on the OS.
//  3. Creates default commands for the detected package manager.
//
// Example usage:
//
//	rootCmd := &cobra.Command{Use: "ipm"}
//	SetupDefaultManagerCommands(rootCmd, "/path/to/configDir", "/path/to/schemaFile", os.Args, os.Args[1], os.Args[2])
//
// This function is useful for setting up default commands for the package manager
// detected based on the OS. It ensures that the default package manager commands
// are available in the CLI.
func SetupDefaultManagerCommands(rootCmd *cobra.Command, configDir string, schemaFile string, argsLength []string, firstArg string, secondArg string) {
	// Check if the validate command is being run
	if len(argsLength) < 3 || firstArg != "manager" || secondArg != "validate" {
		// Detect default package manager based on the OS
		defaultManager := utils.DetectDefaultPackageManager()
		if defaultManager != "" {
			createDefaultCommands(rootCmd, defaultManager, configDir, schemaFile)
		}
	}
}

// createDefaultCommands creates default commands for the specified package manager.
//
// Parameters:
//   - rootCmd: The root command to which the default commands will be added.
//   - managerName: The name of the package manager.
//   - configDir: The directory containing the configuration files for package managers.
//   - schemaFile: The path to the JSON schema file used for validation.
//
// This function performs the following steps:
//  1. Validates JSON files against the schema.
//  2. Reads the commands from the JSON file.
//  3. Unmarshals the config data.
//  4. Checks if the commands are enabled.
//  5. Adds the commands to the root command.
//
// Example usage:
//
//	rootCmd := &cobra.Command{Use: "ipm"}
//	createDefaultCommands(rootCmd, "defaultManager", "/path/to/configDir", "/path/to/schemaFile")
//
// This function is useful for creating default commands for a specified package manager.
// It reads the configuration from a JSON file, validates it against the schema, and
// adds the commands to the root command if they are enabled.
func createDefaultCommands(rootCmd *cobra.Command, managerName string, configDir string, schemaFile string) {
	// Validate JSON files against the schema
	if err := utils.ValidateJSONFiles(schemaFile, configDir); err != nil {
		log.Fatalf("%v", err)
	}

	// Read the commands from the JSON file
	configFile := filepath.Join(configDir, managerName+".json")

	// Read the config file
	data := config.ReadConfigFile(configFile)

	// Unmarshal the config data
	config := config.UnmarshalConfig(data, configFile)

	// Check if the commands are enabled
	if !config.Enabled {
		return
	}

	// Extract and sort the keys
	keys := make([]string, 0, len(config.Commands))
	for key := range config.Commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Create default commands for each command in the JSON file
	for _, command := range keys {
		cmd := &cobra.Command{
			Use:   command + " [params]",
			Short: "Execute " + command + " command for " + managerName,
			Run: func(cmd *cobra.Command, args []string) {
				manager.ExecuteCommandTemplate(command, config.Commands[command], args)
			},
		}
		rootCmd.AddCommand(cmd)
	}
}
