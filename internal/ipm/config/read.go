// Package config provides utilities for managing configuration files
package config

import (
	"encoding/json"
	"ipm/internal/ipm/utils"
	"log"
	"os"
)

// ReadConfigFile reads the specified configuration file and returns its content as a byte slice.
// If the file cannot be read, it logs a fatal error and terminates the program.
//
// Parameters:
//   - configFile: The path to the configuration file to read.
//
// Returns:
//   - []byte: The content of the configuration file as a byte slice.
//
// Example usage:
//
//	data := ReadConfigFile("/path/to/config.json")
//
// This function is typically used to read the content of a configuration file
// before unmarshalling it into a struct or processing it further.
func ReadConfigFile(configFile string) []byte {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", configFile, err)
	}
	return data
}

// UnmarshalConfig unmarshals the provided JSON data into a CommandConfig struct.
// If the data cannot be unmarshalled, it logs a fatal error and terminates the program.
//
// Parameters:
//   - data: The JSON data to unmarshal.
//   - configFile: The path to the configuration file (used for logging purposes).
//
// Returns:
//   - utils.CommandConfig: The unmarshalled CommandConfig struct.
//
// Example usage:
//
//	config := UnmarshalConfig(data, "/path/to/config.json")
//
// This function is typically used to convert JSON data read from a configuration file
// into a CommandConfig struct for further processing or manipulation.
func UnmarshalConfig(data []byte, configFile string) utils.CommandConfig {
	var config utils.CommandConfig
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to unmarshal %s: %v", configFile, err)
	}
	return config
}
