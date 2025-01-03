<!-- markdownlint-disable MD024 -->

# `ipm` - Integrated Package Manager 🚀

## Table of Contents 📑

- [`ipm` - Integrated Package Manager 🚀](#ipm---integrated-package-manager-)
  - [Table of Contents 📑](#table-of-contents-)
  - [About ℹ️](#about-ℹ️)
  - [Key Features ✨](#key-features-)
  - [Installation 🛠️](#installation-️)
    - [Pre-built Binaries 📦](#pre-built-binaries-)
    - [Building from Source 🏗️](#building-from-source-️)
      - [Clone the Repository 🌀](#clone-the-repository-)
      - [Make Build 🔨](#make-build-)
        - [PowerShell Script (Windows) 💻](#powershell-script-windows-)
        - [Shell Script (Linux/macOS) 🐧🍏](#shell-script-linuxmacos-)
  - [Usage 📋](#usage-)
    - [Default Package Manager (auto-detected) 🔍](#default-package-manager-auto-detected-)
      - [Basic Commands 📝](#basic-commands-)
        - [List Installed Packages 📜](#list-installed-packages-)
        - [Search for a Package 🔎](#search-for-a-package-)
        - [Check Information about a Package ℹ️](#check-information-about-a-package-ℹ️)
        - [Install a Package 📥](#install-a-package-)
        - [Update Index of the Package Manager 🔄](#update-index-of-the-package-manager-)
        - [Upgrade a Package ⬆️](#upgrade-a-package-️)
        - [Upgrade all Packages ⬆️⬆️](#upgrade-all-packages-️️)
        - [Remove a package 🗑️](#remove-a-package-️)
      - [Example 💡](#example-)
    - [Custom Package Manager 🛠️](#custom-package-manager-️)
      - [Basic Commands 📝](#basic-commands--1)
        - [List Installed Packages 📜](#list-installed-packages--1)
        - [Search for a Package 🔎](#search-for-a-package--1)
        - [Check Information about a Package ℹ️](#check-information-about-a-package-ℹ️-1)
        - [Install a Package 📥](#install-a-package--1)
        - [Update Index of the Package Manager 🔄](#update-index-of-the-package-manager--1)
        - [Upgrade a Package ⬆️](#upgrade-a-package-️-1)
        - [Upgrade all Packages ⬆️⬆️](#upgrade-all-packages-️️-1)
        - [Remove a package 🗑️](#remove-a-package-️-1)
      - [Example 💡](#example--1)
  - [Configuration ⚙️](#configuration-️)
    - [Example Configuration 📄](#example-configuration-)
  - [Contributing 🤝](#contributing-)
    - [How to Contribute 🛠️](#how-to-contribute-️)
    - [Code of Conduct 📜](#code-of-conduct-)
  - [Acknowledgements 🙏](#acknowledgements-)
    - [Special Thanks 🌟](#special-thanks-)
  - [License 📄](#license-)

## About ℹ️

Integrated Package Manager (`ipm`) is a powerful, cross-platform tool designed to unify the command-line interface (CLI) commands for various package managers. It acts as a universal frontend, working on top of other package managers to provide a seamless and consistent experience for managing software packages across different operating systems.

In today's diverse development environments, developers, software engineers, and system engineers often face the challenge of dealing with multiple package managers, each with its own set of commands and interfaces. This can lead to inefficiencies, errors, and a steep learning curve, especially when switching between different platforms.

`ipm` aims to solve these issues by offering a single, unified CLI that abstracts away the complexities of individual package managers. Whether you're working on Windows, Linux, or macOS, `ipm` provides a consistent and intuitive interface for installing, updating, and removing software packages. By standardizing package management commands, `ipm` simplifies the workflow, reduces the potential for errors, and enhances productivity.

```txt
$ ipm 
Usage:
  ipm [command]

Available Commands:
  manager     Manage package manager configurations
  info        Execute info command for apt
  install     Execute install command for apt
  list        Execute list command for apt
  search      Execute search command for apt
  uninstall   Execute uninstall command for apt
  update      Execute update command for apt
  upgrade     Execute upgrade command for apt
  upgrade-all Execute upgrade-all command for apt
  apk         Run apk commands
  apt         Run apt commands
  brew        Run brew commands
  choco       Run choco commands
  dnf         Run dnf commands
  emerge      Run emerge commands
  eopkg       Run eopkg commands
  npm         Run npm commands
  pip         Run pip commands
  pip3        Run pip3 commands
  winget      Run winget commands
  yum         Run yum commands
  zypper      Run zypper commands

Flags:
  -h, --help   help for ipm

Use "ipm [command] --help" for more information about a command.
```

## Key Features ✨

- **Cross-Platform Compatibility**: Supports multiple operating systems including Windows, Linux, and macOS.
- **Unified Interface**: Provides a single, consistent CLI interface for managing packages across different platforms.
- **Efficient Package Management**: Simplifies the installation, update, and removal of software packages.
- **Customizable Configurations**: Allows users to define custom commands and configurations for different package managers.
- **Extensible**: Easily extendable to support additional package managers and custom commands.

## Installation 🛠️

### Pre-built Binaries 📦

Pre-built binaries for various platforms are available in the [releases](https://github.com/imfsiddiqui/ipm/releases) section. Download the appropriate binary for your platform and add it to your system's PATH.

1. **Download the binary**: Navigate to the [releases](https://github.com/imfsiddiqui/ipm/releases) page and download the binary for your operating system.
2. **Move the binary to a directory**: Move the downloaded binary to a directory of your choice.
3. **Add the binary to your PATH**: Add the directory containing the binary to your system's PATH. This allows you to run `ipm` from any terminal session.
4. For example:

  On Unix-like systems (Linux/macOS):

  ```sh
  # Add the directory to your PATH
  echo 'export PATH=$PATH:/path/to/ipm' >> ~/.bashrc
  source ~/.bashrc
  ```

  On Windows:

  ```powershell
  # Add the directory to your PATH
  [System.Environment]::SetEnvironmentVariable("Path", $env:Path + ";path\to\ipm", [System.EnvironmentVariableTarget]::Machine)
  ```

### Building from Source 🏗️

To build `ipm` from source, you need to have [Go](https://golang.org/dl/) installed on your system.

#### Clone the Repository 🌀

```sh
git clone https://github.com/your-repo/ipm.git
cd ipm
```

#### Make Build 🔨

To build `ipm`, you can use the provided build scripts.

##### PowerShell Script (Windows) 💻

1. Open a PowerShell terminal.
2. Run the build script:

  ```powershell
  .\scripts\build\build.ps1
  ```

##### Shell Script (Linux/macOS) 🐧🍏

1. Open a terminal.
2. Run the build script:

  ```sh
  ./scripts/build/build.sh
  ```

These scripts will build the binaries for the multiple platforms and create archives (ZIP for Windows and TAR.GZ for other platforms) in the `dist/release` directory.

## Usage 📋

### Default Package Manager (auto-detected) 🔍

#### Basic Commands 📝

##### List Installed Packages 📜

```sh
ipm list
```

##### Search for a Package 🔎

```sh
ipm search <package-name>
```

##### Check Information about a Package ℹ️

```sh
ipm info <package-name>
```

##### Install a Package 📥

```sh
ipm install <package-name>
```

##### Update Index of the Package Manager 🔄

```sh
ipm update
```

##### Upgrade a Package ⬆️

```sh
ipm upgrade <package-name>
```

##### Upgrade all Packages ⬆️⬆️

```sh
ipm upgrade-all
```

##### Remove a package 🗑️

```sh
ipm uninstall <package-name>
```

#### Example 💡

To install the `jq` package:

```sh
ipm install jq
```

### Custom Package Manager 🛠️

#### Basic Commands 📝

##### List Installed Packages 📜

```sh
ipm <package-manager> list
```

##### Search for a Package 🔎

```sh
ipm <package-manager> search <package-name>
```

##### Check Information about a Package ℹ️

```sh
ipm <package-manager> info <package-name>
```

##### Install a Package 📥

```sh
ipm <package-manager> install <package-name>
```

##### Update Index of the Package Manager 🔄

```sh
ipm <package-manager> update
```

##### Upgrade a Package ⬆️

```sh
ipm <package-manager> upgrade <package-name>
```

##### Upgrade all Packages ⬆️⬆️

```sh
ipm <package-manager> upgrade-all
```

##### Remove a package 🗑️

```sh
ipm <package-manager> uninstall <package-name>
```

#### Example 💡

To install the `fast-json-stringify` package using `npm`:

```sh
ipm npm install fast-json-stringify
```

## Configuration ⚙️

`ipm` uses a JSON configuration file to define custom commands and settings for different package managers. The configuration file is located in the config directory.

### Example Configuration 📄

```json
{
  "enabled": true,
  "commands": {
    "info": "apt-cache show {{.Package}}",
    "install": "apt-get install -y {{.Package}}",
    "list": "dpkg --list",
    "search": "apt-cache search {{.Package}}",
    "uninstall": "apt-get remove -y {{.Package}}",
    "update": "apt-get update",
    "upgrade": "apt-get install --only-upgrade {{.Package}}",
    "upgrade-all": "apt-get upgrade"
  }
}
```

## Contributing 🤝

Contributions are welcome! If you have suggestions for improvements, please open an issue or submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

### How to Contribute 🛠️

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

### Code of Conduct 📜

This project adheres to the Contributor Covenant code of conduct. By participating, you are expected to uphold this code.

## Acknowledgements 🙏

I would like to extend my heartfelt thanks to all the developers and contributors whose work has made this project possible. Your dedication and contributions to the open-source community are invaluable, and I am grateful for the tools, libraries, and frameworks that you have created and maintained.

### Special Thanks 🌟

I would like to specifically acknowledge the developers of the following tools, frameworks and dependencies which are used in this project:

- **Go Language Team**: For providing a powerful and efficient programming language that forms the backbone of this project.
- **Package Manager Developers**: For creating and maintaining the various package managers that `ipm` integrates with, including `apt`, `yum`, `brew`, `choco`, `npm`, `pip`, and many others.
- **Library Authors**: For the numerous libraries and tools that have been utilized in this project, enhancing its functionality and performance, specially:
  - **[github.com/spf13/cobra](https://github.com/spf13/cobra)**: For creating a library for building powerful modern CLI applications.
  - **[github.com/xeipuuv/gojsonschema](https://github.com/xeipuuv/gojsonschema)**: For providing a library to validate JSON schemas.
  - **[github.com/inconshreveable/mousetrap](https://github.com/inconshreveable/mousetrap)**: For helping detect when a Go program is run from a Windows shortcut.
  - **[github.com/spf13/pflag](https://github.com/spf13/pflag)**: For providing a POSIX/GNU-style flag parsing library.
  - **[github.com/xeipuuv/gojsonpointer](https://github.com/xeipuuv/gojsonpointer)**: For implementing JSON Pointer (RFC 6901).
  - **[github.com/xeipuuv/gojsonreference](https://github.com/xeipuuv/gojsonreference)**: For implementing JSON Reference (RFC 6901).

Your hard work and commitment to excellence have made it possible for me to build `ipm` and provide a unified package management experience for developers across different platforms. Thank you for your contributions to the open-source ecosystem.

## License 📄

This project is licensed under the GPL-3.0 license. See the [LICENSE](https://github.com/imfsiddiqui/ipm/blob/main/LICENSE) file for details.
