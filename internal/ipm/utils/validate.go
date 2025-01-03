// Package utils provides utility functions for the application
package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

// ValidateJSONFiles validates all JSON files in the specified directory against the provided schema.
//
// Parameters:
//   - schemaFile: The path to the JSON schema file.
//   - configDir: The directory containing the JSON files to validate.
//
// Returns:
//   - error: An error if any of the JSON files fail validation, or nil if all files are valid.
//
// Example usage:
//
//	err := ValidateJSONFiles("/path/to/schema.json", "/path/to/json/dir")
//
// This function performs the following steps:
//  1. Uses a wildcard to get all JSON files in the specified directory.
//  2. Validates each JSON file against the schema using the ValidateJSONFile function.
//  3. Collects validation errors in a strings.Builder.
//  4. Returns an error if any of the JSON files fail validation, or nil if all files are valid.
func ValidateJSONFiles(schemaFile string, configDir string) error {
	// Create a strings.Builder to collect validation errors
	var builder strings.Builder

	// Use wildcard to get all JSON files in the specified directory
	configFiles, err := filepath.Glob(filepath.Join(configDir, "*.json"))
	if err != nil {
		return fmt.Errorf("failed to read JSON files: %v", err)
	}

	// Iterate over each JSON file and validate it against the schema
	for _, configFile := range configFiles {
		if err := ValidateJSONFile(schemaFile, configFile); err != nil {
			// Append validation error to the builder
			builder.WriteString("Validation error: ")
			builder.WriteString(err.Error())
			builder.WriteString("\n")
		}
	}

	// Convert the builder to a string
	validationErrors := builder.String()

	// Return an error if there are any validation errors, otherwise return nil
	if len(validationErrors) > 0 {
		return fmt.Errorf("%v", validationErrors)
	} else {
		return nil
	}
}

// ValidateJSONFile validates a single JSON file against the provided schema.
//
// Parameters:
//   - schemaFile: The path to the JSON schema file.
//   - jsonFile: The path to the JSON file to validate.
//
// Returns:
//   - error: An error if the JSON file fails validation, or nil if the file is valid.
//
// Example usage:
//
//	err := ValidateJSONFile("/path/to/schema.json", "/path/to/json/file.json")
//
// This function performs the following steps:
//  1. Loads the schema and JSON data using the loadSchemaAndData function.
//  2. Validates the JSON data against the schema using gojsonschema.Validate.
//  3. Checks if the validation was successful and returns an error if it was not.
func ValidateJSONFile(schemaFile string, jsonFile string) error {
	// Load the schema and JSON data
	schemaLoader, documentLoader, err := loadSchemaAndData(schemaFile, jsonFile)
	if err != nil {
		return err
	}

	// Validate the JSON data against the schema
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("failed to validate JSON file: %v", err)
	}

	// Check if the validation was successful
	if !result.Valid() {
		var validationErrors []string
		for _, desc := range result.Errors() {
			validationErrors = append(validationErrors, desc.String())
		}
		return fmt.Errorf("%v in %s", strings.Join(validationErrors, ", "), jsonFile)
	}

	return nil
}

// loadSchemaAndData loads the JSON schema and data from the specified files.
//
// Parameters:
//   - schemaFile: The path to the JSON schema file.
//   - jsonFile: The path to the JSON file to load.
//
// Returns:
//   - gojsonschema.JSONLoader: The schema loader.
//   - gojsonschema.JSONLoader: The document loader.
//   - error: An error if the schema or JSON data cannot be loaded, or nil if successful.
//
// Example usage:
//
//	schemaLoader, documentLoader, err := loadSchemaAndData("/path/to/schema.json", "/path/to/json/file.json")
//
// This function performs the following steps:
//  1. Loads the JSON schema from the specified file.
//  2. Loads the JSON data from the specified file.
//  3. Returns the schema and document loaders, or an error if the files cannot be loaded.
func loadSchemaAndData(schemaFile string, jsonFile string) (gojsonschema.JSONLoader, gojsonschema.JSONLoader, error) {
	// Get the absolute path for the schema file
	schemaPath, err := filepath.Abs(schemaFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get absolute path for schema: %v", err)
	}

	// Convert backslashes to forward slashes for URL format
	schemaPath = strings.ReplaceAll(schemaPath, "\\", "/")

	// Load the JSON schema
	schemaLoader := gojsonschema.NewReferenceLoader("file:///" + schemaPath)

	// Read the JSON file
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read JSON file: %v", err)
	}

	// Load the JSON data
	documentLoader := gojsonschema.NewBytesLoader(data)

	return schemaLoader, documentLoader, nil
}
