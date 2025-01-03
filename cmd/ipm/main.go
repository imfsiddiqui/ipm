// Package main is the entry point for the IPM application.
package main

import (
	"path/filepath"

	"ipm/internal/ipm/cli"
)

// cliCmd is the name of the command-line interface (CLI) application.
var cliCmd = "ipm"

// main is the entry point for the IPM application.
func main() {
	// Get the executable path
	// This function retrieves the path of the currently running executable.
	// It is useful for determining the directory structure relative to the executable.
	exePath := cli.GetExecutablePath()

	// Get the directory of the executable
	// This function extracts the directory part of the executable path.
	exeDir := filepath.Dir(exePath)

	// Define the configuration directory and schema file
	// The configuration directory contains the configuration files for package managers.
	// The schema file is used to validate the configuration files.
	configDir := filepath.Join(exeDir, "config", "manager", "config")
	schemaFile := filepath.Join(exeDir, "config", "manager", "schema", "manager.json")

	// Initialize the CLI commands and structure
	// This function sets up the command-line interface (CLI) commands and their structure.
	// It uses the configuration directory and schema file to manage package manager configurations.
	cli.InitializeCLI(configDir, schemaFile, cliCmd)
}
