// The module directive defines the module's path, which is the import path
// prefix for all packages within the module.
module ipm

// The go directive specifies the version of the Go language used to build
// the module. This helps ensure compatibility with the specified Go version.
go 1.23.4

// The require block lists the module dependencies required by the project.
// Each dependency is specified with its module path and version.
require (
	// github.com/spf13/cobra is a library for creating powerful modern CLI
	// applications. It provides a simple interface to define commands and
	// their flags.
	github.com/spf13/cobra v1.8.1

	// github.com/xeipuuv/gojsonschema is a library for validating JSON
	// schemas. It provides functions to load and validate JSON data against
	// JSON schemas.
	github.com/xeipuuv/gojsonschema v1.2.0
)

// The second require block lists indirect dependencies. These are dependencies
// that are not directly imported by the project but are required by other
// dependencies. Indirect dependencies are automatically managed by Go modules.
require (
	// github.com/inconshreveable/mousetrap is a library that helps detect
	// when a Go program is run from a Windows shortcut. It is used by
	// github.com/spf13/cobra to improve the user experience on Windows.
	github.com/inconshreveable/mousetrap v1.1.0 // indirect

	// github.com/spf13/pflag is a library that provides a POSIX/GNU-style
	// flag parsing. It is used by github.com/spf13/cobra to handle command-line
	// flags.
	github.com/spf13/pflag v1.0.5 // indirect

	// github.com/xeipuuv/gojsonpointer is a library that implements JSON
	// Pointer (RFC 6901). It is used by github.com/xeipuuv/gojsonschema to
	// navigate JSON documents.
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect

	// github.com/xeipuuv/gojsonreference is a library that implements JSON
	// Reference (RFC 6901). It is used by github.com/xeipuuv/gojsonschema to
	// resolve JSON references.
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
)
