// Package utils provides utility functions for the application
package utils

// CommandConfig represents the structure of the commands in the JSON file.
//
// The CommandConfig struct is used to parse and store the configuration of
// commands from a JSON file. It includes fields for enabling/disabling the
// config and a map of command names to their respective command strings.
//
// Fields:
//   - Enabled: A boolean indicating whether the config are enabled or not.
//   - Commands: A map where the keys are command names and the values are
//     the corresponding command strings.
//
// Example JSON structure:
//
//	{
//	  "enabled": true,
//	  "commands": {
//	    "install": "install-command",
//	    "update": "update-command",
//	    ...
//	  }
//	}
//
// This struct is useful for managing the configuration of commands in a
// structured and easily accessible manner.
type CommandConfig struct {
	Enabled  bool              `json:"enabled"`  // Indicates if the config is enabled
	Commands map[string]string `json:"commands"` // Map of command names to command strings
}
