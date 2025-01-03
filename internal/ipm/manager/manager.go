// Package manager provides utilities for executing command templates and running commands
package manager

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"text/template"
)

// ExecuteCommandTemplate executes a command template with the given parameters.
//
// Parameters:
//   - command: The base command to execute.
//   - templateStr: The command template string to parse and execute.
//   - params: A slice of strings containing the parameters to pass to the template.
//
// Example usage:
//
//	ExecuteCommandTemplate("apt-get install -y", "{{.Package}}", []string{"jq"})
//
// This function performs the following steps:
//  1. Parses the command template with the given parameters using the parseCommandTemplate function.
//  2. Checks if the final command string is empty and prints a message if the command is not available.
//  3. Prints the final command string to be executed.
//  4. Runs the command using the final command string by calling the runCommand function.
func ExecuteCommandTemplate(command string, templateStr string, params []string) {
	// Parse and execute the command template
	finalCmdStr := parseCommandTemplate(templateStr, params)

	// Checks if the final command string is empty
	if finalCmdStr == "" {
		// Prints a message if the command is not available
		fmt.Printf("Executing %s: %s\n", command, "Command not available")
	} else {
		// Prints the final command string to be executed
		fmt.Printf("Executing %s: %s\n", command, finalCmdStr)
		// Run the command
		runCommand(command, finalCmdStr)
	}
}

// parseCommandTemplate parses and executes the command template with the given parameters.
//
// Parameters:
//   - templateStr: The command template string to parse and execute.
//   - params: A slice of strings containing the parameters to pass to the template.
//
// Returns:
//   - string: The final command string after parsing and executing the template.
//
// Example usage:
//
//	finalCmdStr := parseCommandTemplate("{{.Package}}!", []string{"jq"})
//
// This function performs the following steps:
//  1. Parses the command template using the provided template string.
//  2. Creates a map for template data and populates it with the provided parameters.
//  3. Executes the template with the provided data and stores the result in a buffer.
//  4. Returns the final command string from the buffer.
func parseCommandTemplate(templateStr string, params []string) string {
	// Parse the command template
	tmpl, err := template.New("command").Parse(templateStr)
	if err != nil {
		log.Fatalf("Failed to parse command template: %v", err)
	}

	// Create a map for template data
	templateData := make(map[string]string)
	if len(params) > 0 {
		// If there are parameters, use the first one as the "Package" value
		templateData["Package"] = params[0]
	}

	// Buffer to hold the executed template result
	var cmdBuffer bytes.Buffer

	// Execute the template with the provided data
	if err := tmpl.Execute(&cmdBuffer, templateData); err != nil {
		log.Fatalf("Failed to execute command template: %v", err)
	}

	// Get the final command string from the buffer
	return cmdBuffer.String()
}

// runCommand creates and runs the command, capturing and printing the output.
//
// Parameters:
//   - command: The base command to execute.
//   - finalCmdStr: The final command string to execute.
//
// Example usage:
//
//	runCommand("apt-get install -y", "jq")
//
// This function performs the following steps:
//  1. Creates the command to be executed based on the operating system.
//  2. Set the output streams to directly stream the output.
//  3. Run the command.
func runCommand(command string, finalCmdStr string) {
	// Create the command to be executed
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// On Windows, use "cmd /C" to execute the command
		cmd = exec.Command("cmd", "/C", finalCmdStr)
	} else {
		// On Unix-like systems, use "sh -c" to execute the command
		cmd = exec.Command("sh", "-c", finalCmdStr)
	}

	// Set the output streams to directly stream the output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to execute %s: %v", command, err)
	}
}
