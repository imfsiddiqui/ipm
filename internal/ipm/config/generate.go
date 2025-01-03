// Package config provides utilities for managing configuration files
package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ipm/internal/ipm/utils"
)

// GenerateManagerConfig generates a new package manager config.
//
// Parameters:
//   - managerName: The name of the package manager for which to generate the config.
//   - configDir: The directory where the configuration files are stored.
//
// Example usage:
//
//	config.GenerateManagerConfig("apt", "/path/to/config/dir")
//
// This function performs the following steps:
//  1. Constructs the path to the configuration file for the specified package manager.
//  2. Initializes a new CommandConfig struct and a map to hold the commands.
//  3. Prompts the user to enter commands for various package manager operations.
//  4. Reads the user input and trims any whitespace.
//  5. Stores the commands in the CommandConfig struct.
//  6. Prompts the user to enable or disable the package manager.
//  7. Marshals the CommandConfig struct into JSON data.
//  8. Writes the JSON data to the configuration file.
//  9. Prints a message indicating that the package manager config has been generated.
func GenerateManagerConfig(managerName string, configDir string) {
	// Construct the path to the configuration file
	configFile := filepath.Join(configDir, managerName+".json")

	// Initialize a new CommandConfig struct and a map to hold the commands
	var config utils.CommandConfig
	config.Commands = make(map[string]string)

	// Create a new reader to read user input from stdin
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter commands for various package manager operations
	fmt.Printf("Enter command for 'info': ")
	infoCmd, _ := reader.ReadString('\n')
	config.Commands["info"] = strings.TrimSpace(infoCmd)

	fmt.Printf("Enter command for 'install': ")
	installCmd, _ := reader.ReadString('\n')
	config.Commands["install"] = strings.TrimSpace(installCmd)

	fmt.Printf("Enter command for 'list': ")
	listInstalledCmd, _ := reader.ReadString('\n')
	config.Commands["list"] = strings.TrimSpace(listInstalledCmd)

	fmt.Printf("Enter command for 'search': ")
	searchCmd, _ := reader.ReadString('\n')
	config.Commands["search"] = strings.TrimSpace(searchCmd)

	fmt.Printf("Enter command for 'uninstall': ")
	uninstallCmd, _ := reader.ReadString('\n')
	config.Commands["uninstall"] = strings.TrimSpace(uninstallCmd)

	fmt.Printf("Enter command for 'update': ")
	updateIndexCmd, _ := reader.ReadString('\n')
	config.Commands["update"] = strings.TrimSpace(updateIndexCmd)

	fmt.Printf("Enter command for 'upgrade': ")
	upgradeCmd, _ := reader.ReadString('\n')
	config.Commands["upgrade"] = strings.TrimSpace(upgradeCmd)

	fmt.Printf("Enter command for 'upgrade-all': ")
	upgradeAllCmd, _ := reader.ReadString('\n')
	config.Commands["upgrade-all"] = strings.TrimSpace(upgradeAllCmd)

	// Prompt the user to enable or disable the package manager
	var enabledInput string
	for {
		fmt.Printf("Enable this manager? (yes/no): ")
		enabledInput, _ = reader.ReadString('\n')
		enabledInput = strings.ToLower(strings.TrimSpace(enabledInput))
		if enabledInput == "yes" {
			config.Enabled = true
			break
		} else if enabledInput == "no" {
			config.Enabled = false
			break
		} else {
			fmt.Println("Invalid input. Please enter 'yes' or 'no'.")
		}
	}

	// Marshal the config data
	newData := marshalConfig(config, configFile)

	// Write the config file
	writeConfigFile(configFile, newData)

	// Print a message indicating that the package manager config has been generated
	fmt.Printf("Manager %s config has been generated", managerName)
}
