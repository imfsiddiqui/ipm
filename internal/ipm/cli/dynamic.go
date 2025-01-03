// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"ipm/internal/ipm/config"
	"ipm/internal/ipm/manager"
	"ipm/internal/ipm/utils"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// SetupDynamicManagerCommands sets up the default manager commands based on the OS.
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
//  2. Dynamically reads manager names from the config directory.
//  3. Creates commands for each detected package manager.
//
// Example usage:
//
//	rootCmd := &cobra.Command{Use: "ipm"}
//	SetupDynamicManagerCommands(rootCmd, "/path/to/configDir", "/path/to/schemaFile", os.Args, os.Args[1], os.Args[2])
//
// This function is useful for dynamically setting up commands for package managers
// based on the configuration files present in the config directory. It ensures that
// the commands for each package manager are available in the CLI.
func SetupDynamicManagerCommands(rootCmd *cobra.Command, configDir string, schemaFile string, argsLength []string, firstArg string, secondArg string) {
	// Check if the validate command is being run
	if len(argsLength) < 3 || firstArg != "manager" || secondArg != "validate" {
		// Dynamically read manager names from the config directory
		managerFiles, err := filepath.Glob(filepath.Join(configDir, "*.json"))
		if err != nil {
			log.Fatalf("Failed to read manager files: %v", err)
		}

		// Iterate over each manager file and create commands
		for _, managerFile := range managerFiles {
			managerName := strings.TrimSuffix(filepath.Base(managerFile), ".json")
			managerCmd := createManagerCommand(managerName, configDir, schemaFile)
			if managerCmd != nil {
				rootCmd.AddCommand(managerCmd)
			}
		}
	}
}

// createManagerCommand creates a command for the specified package manager.
//
// Parameters:
//   - managerName: The name of the package manager.
//   - configDir: The directory containing the configuration files for package managers.
//   - schemaFile: The path to the JSON schema file used for validation.
//
// This function performs the following steps:
//  1. Validates JSON files against the schema.
//  2. Reads the commands from the JSON file.
//  3. Unmarshals the config data.
//  4. Checks if the commands are enabled.
//  5. Creates and returns a cobra.Command for the package manager.
//
// Example usage:
//
//	managerCmd := createManagerCommand("apt", "/path/to/configDir", "/path/to/schemaFile")
//
// This function is useful for creating a command for a specified package manager.
// It reads the configuration from a JSON file, validates it against the schema, and
// creates a cobra.Command if the commands are enabled.
func createManagerCommand(managerName string, configDir string, schemaFile string) *cobra.Command {
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
		return nil
	}

	// Extract and sort the keys
	keys := make([]string, 0, len(config.Commands))
	for key := range config.Commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Create a new cobra.Command for the package manager
	managerCmd := &cobra.Command{
		Use:   managerName,
		Short: "Run " + managerName + " commands",
	}

	// Create default commands for each command in the JSON file
	for _, command := range keys {
		cmd := &cobra.Command{
			Use:   command + " [params]",
			Short: "Execute " + command + " command for " + managerName + " package manager",
			Run: func(cmd *cobra.Command, args []string) {
				manager.ExecuteCommandTemplate(command, config.Commands[command], args)
			},
		}
		managerCmd.AddCommand(cmd)
	}

	return managerCmd
}
