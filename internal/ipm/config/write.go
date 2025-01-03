// Package config provides utilities for managing configuration files
package config

import (
	"encoding/json"
	"ipm/internal/ipm/utils"
	"log"
	"os"
)

// marshalConfig marshals the CommandConfig struct into JSON data.
// If the struct cannot be marshalled, it logs a fatal error and terminates the program.
//
// Parameters:
//   - config: The CommandConfig struct to marshal.
//   - configFile: The path to the configuration file (used for logging purposes).
//
// Returns:
//   - []byte: The marshalled JSON data.
//
// Example usage:
//
//	data := marshalConfig(config, "/path/to/config.json")
//
// This function is typically used to convert a CommandConfig struct into JSON data
// before writing it to a configuration file.
func marshalConfig(config utils.CommandConfig, configFile string) []byte {
	newData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal %s: %v", configFile, err)
	}
	return newData
}

// writeConfigFile writes the JSON data to the specified configuration file.
// If the file cannot be written, it logs a fatal error and terminates the program.
//
// Parameters:
//   - configFile: The path to the configuration file to write.
//   - data: The JSON data to write to the file.
//
// Example usage:
//
//	writeConfigFile("/path/to/config.json", data)
//
// This function is typically used to write JSON data to a configuration file
// after marshalling a struct or processing the data in some way.
func writeConfigFile(configFile string, data []byte) {
	if err := os.WriteFile(configFile, data, 0644); err != nil {
		log.Fatalf("Failed to write %s: %v", configFile, err)
	}
}
