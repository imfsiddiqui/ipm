---
layout: default
---

<!-- markdownlint-disable MD024 MD033 MD041 -->

<a id="top"></a>

<div align=center>

<p>
  🌍 <strong><a href="https://imfsiddiqui.github.io/{{ site.repository_name }}">Web Page</a></strong>
  |
  💻 <strong><a href="https://github.com/imfsiddiqui/{{ site.repository_name }}">Source Code</a></strong>
  |
  🚀 <strong><a href="https://github.com/imfsiddiqui/{{ site.repository_name }}/releases">Releases</a></strong>
</p>

</div>

# 📦 ipm - Integrated Package Manager

A cross-platform tool that unifies CLI for various package managers, simplifies installation, updates, and removal of packages by standardizing commands.

<div align="center">
  <img
    src="./assets/images/ipm-banner-standard.svg"
    style="border-radius: 10px"
    alt="ipm project banner"
  />
</div>

## 📚 Table of Contents

- [📦 ipm - Integrated Package Manager](#-ipm---integrated-package-manager)
  - [📚 Table of Contents](#-table-of-contents)
  - [📌 About](#-about)
  - [✨ Key Features](#-key-features)
  - [📦 Supported Package Managers](#-supported-package-managers)
  - [🛠️ Installation](#️-installation)
    - [📦 Pre-built Binaries](#-pre-built-binaries)
      - [⬇️ Download the Binary](#️-download-the-binary)
      - [🚂 Move the Binary to a Directory](#-move-the-binary-to-a-directory)
      - [🫵 Add the Binary to PATH](#-add-the-binary-to-path)
        - [🐧 Linux](#-linux)
        - [🍏 macOS](#-macos)
        - [🪟 Windows](#-windows)
    - [🏗️ Building from Source](#️-building-from-source)
      - [🌀 Clone the Repository](#-clone-the-repository)
      - [🔨 Make Build](#-make-build)
        - [🐧 Linux (Shell Script)](#-linux-shell-script)
        - [🍏 macOS (Shell Script)](#-macos-shell-script)
        - [🪟 Windows (PowerShell Script)](#-windows-powershell-script)
    - [🐳 Dockerfiles](#-dockerfiles)
      - [🏗️ Building Docker Images](#️-building-docker-images)
      - [🚀 Running Docker Container](#-running-docker-container)
  - [📋 Usage](#-usage)
    - [🔍 Default Package Manager (auto-detected)](#-default-package-manager-auto-detected)
      - [📝 Basic Commands](#-basic-commands)
        - [📃 List Installed Packages](#-list-installed-packages)
        - [🔎 Search for a Package](#-search-for-a-package)
        - [📖 Check Information about a Package](#-check-information-about-a-package)
        - [📥 Install a Package](#-install-a-package)
        - [🔄 Update Index of the Package Manager](#-update-index-of-the-package-manager)
        - [⬆️ Upgrade a Package](#️-upgrade-a-package)
        - [⬆️ Upgrade all Packages](#️-upgrade-all-packages)
        - [🗑️ Remove a package](#️-remove-a-package)
      - [💡 Example](#-example)
    - [🛠️ Custom Package Manager](#️-custom-package-manager)
      - [📝 Basic Commands](#-basic-commands-1)
        - [📃 List Installed Packages](#-list-installed-packages-1)
        - [🔎 Search for a Package](#-search-for-a-package-1)
        - [📖 Check Information about a Package](#-check-information-about-a-package-1)
        - [📥 Install a Package](#-install-a-package-1)
        - [🔄 Update Index of the Package Manager](#-update-index-of-the-package-manager-1)
        - [⬆️ Upgrade a Package](#️-upgrade-a-package-1)
        - [⬆️ Upgrade all Packages](#️-upgrade-all-packages-1)
        - [🗑️ Remove a package](#️-remove-a-package-1)
      - [💡 Example](#-example-1)
  - [⚙️ Configuration](#️-configuration)
    - [🪞 Example Configuration](#-example-configuration)
  - [📄 Important Documents](#-important-documents)
  - [🙏 Acknowledgements](#-acknowledgements)
    - [🌟 Special Thanks](#-special-thanks)
  - [📜 License](#-license)

<p align="right"><a href="#top">☝️</a></p>

## 📌 About

`ipm` - Integrated Package Manager is a powerful, cross-platform tool designed to unify the command-line interface (CLI) commands for various package managers. It acts as a universal frontend, working on top of other package managers to provide a seamless and consistent experience for managing software packages across different operating systems.

In today's diverse development environments, developers, software engineers, and system engineers often face the challenge of dealing with multiple package managers, each with its own set of commands and interfaces. This can lead to inefficiencies, errors, and a steep learning curve, especially when switching between different platforms.

`ipm` aims to solve these issues by offering a single, unified CLI that abstracts away the complexities of individual package managers. Whether you're working on Windows, Linux, or macOS, `ipm` provides a consistent and intuitive interface for installing, updating, and removing software packages. By standardizing package management commands, `ipm` simplifies the workflow, reduces the potential for errors, and enhances productivity.

```console
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

<p align="right"><a href="#top">☝️</a></p>

## ✨ Key Features

- **Cross-Platform Compatibility**: Supports multiple operating systems including Windows, Linux, and macOS.
- **Unified Interface**: Provides a single, consistent CLI interface for managing packages across different platforms.
- **Efficient Package Management**: Simplifies the installation, update, and removal of software packages.
- **Customizable Configurations**: Allows users to define custom commands and configurations for different package managers.
- **Extensible**: Easily extendable to support additional package managers and custom commands.

<p align="right"><a href="#top">☝️</a></p>

## 📦 Supported Package Managers

| **Package Manager** | **Enabled (default)** | **Disabled (default)** |
| :-----------------: | :-------------------: | :--------------------: |
|        `apk`        |          ✅           |           ❌           |
|        `apt`        |          ✅           |           ❌           |
|       `brew`        |          ✅           |           ❌           |
|       `cards`       |          ❌           |           ✅           |
|       `choco`       |          ✅           |           ❌           |
|        `dnf`        |          ✅           |           ❌           |
|      `emerge`       |          ✅           |           ❌           |
|       `eopkg`       |          ✅           |           ❌           |
|      `flatpak`      |          ❌           |           ✅           |
|       `guix`        |          ❌           |           ✅           |
|       `nala`        |          ❌           |           ✅           |
|      `nix-env`      |          ❌           |           ✅           |
|        `npm`        |          ✅           |           ❌           |
|       `opkg`        |          ❌           |           ✅           |
|      `pacman`       |          ❌           |           ✅           |
|        `pip`        |          ✅           |           ❌           |
|       `pip3`        |          ✅           |           ❌           |
|       `scoop`       |          ❌           |           ✅           |
|     `slackpkg`      |          ❌           |           ✅           |
|       `snap`        |          ❌           |           ✅           |
|      `winget`       |          ✅           |           ❌           |
|       `xbps`        |          ❌           |           ✅           |
|        `yum`        |          ✅           |           ❌           |
|      `zypper`       |          ✅           |           ❌           |

<p align="right"><a href="#top">☝️</a></p>

## 🛠️ Installation

### 📦 Pre-built Binaries

Pre-built binaries for various platforms are available in the [releases](https://github.com/imfsiddiqui/ipm/releases) section. Download the appropriate binary for your platform and add it to your system's PATH.

#### ⬇️ Download the Binary

Navigate to the [releases](https://github.com/imfsiddiqui/ipm/releases) page and download the binary for your operating system.

#### 🚂 Move the Binary to a Directory

Move the downloaded binary to a directory of your choice.

#### 🫵 Add the Binary to PATH

Add the directory containing the binary to your system's PATH. This allows you to run `ipm` from any terminal session.

##### 🐧 Linux

```console
# Add the directory to your PATH
echo 'export PATH=$PATH:/path/to/ipm' >> ~/.bashrc
source ~/.bashrc
```

##### 🍏 macOS

```console
# Add the directory to your PATH
echo 'export PATH=$PATH:/path/to/ipm' >> ~/.bashrc
source ~/.bashrc
```

##### 🪟 Windows

```console
# Add the directory to your PATH
[System.Environment]::SetEnvironmentVariable(
  "Path",
  $env:Path + ";path\to\ipm",
  [System.EnvironmentVariableTarget]::Machine
)
```

### 🏗️ Building from Source

To build `ipm` from source, you need to have [Go](https://golang.org/dl/) installed on your system.

#### 🌀 Clone the Repository

```console
git clone https://github.com/your-repo/ipm.git
cd ipm
```

#### 🔨 Make Build

To build `ipm`, you can use the provided build scripts.

##### 🐧 Linux (Shell Script)

- Open a terminal.
- Run the build script: `./scripts/build/build.sh`

##### 🍏 macOS (Shell Script)

- Open a terminal.
- Run the build script: `./scripts/build/build.sh`

##### 🪟 Windows (PowerShell Script)

- Open a PowerShell terminal.
- Run the build script: `.\scripts\build\build.ps1`.

These scripts will build the binaries for the multiple platforms and create archives (ZIP for Windows and TAR.GZ for other platforms) in the `dist/release` directory.

### 🐳 Dockerfiles

Dockerfiles are provided to build Docker images for `ipm`. These Dockerfiles are located in the `dockerfiles` directory.

#### 🏗️ Building Docker Images

To build a Docker image for `ipm`, run the following command from project root:

```console
docker build -t ipm-<package-manager-name>:latest -f dockerfiles/<package-manager-name>/Dockerfile .
```

#### 🚀 Running Docker Container

To run `ipm` Docker container:

```console
docker run --rm -it ipm-<package-manager-name>:latest /bin/sh
```

<p align="right"><a href="#top">☝️</a></p>

## 📋 Usage

### 🔍 Default Package Manager (auto-detected)

#### 📝 Basic Commands

##### 📃 List Installed Packages

```console
ipm list
```

##### 🔎 Search for a Package

```console
ipm search <package-name>
```

##### 📖 Check Information about a Package

```console
ipm info <package-name>
```

##### 📥 Install a Package

```console
ipm install <package-name>
```

##### 🔄 Update Index of the Package Manager

```console
ipm update
```

##### ⬆️ Upgrade a Package

```console
ipm upgrade <package-name>
```

##### ⬆️ Upgrade all Packages

```console
ipm upgrade-all
```

##### 🗑️ Remove a package

```console
ipm uninstall <package-name>
```

#### 💡 Example

To install the `jq` package:

```console
ipm install jq
```

### 🛠️ Custom Package Manager

#### 📝 Basic Commands

##### 📃 List Installed Packages

```console
ipm <package-manager> list
```

##### 🔎 Search for a Package

```console
ipm <package-manager> search <package-name>
```

##### 📖 Check Information about a Package

```console
ipm <package-manager> info <package-name>
```

##### 📥 Install a Package

```console
ipm <package-manager> install <package-name>
```

##### 🔄 Update Index of the Package Manager

```console
ipm <package-manager> update
```

##### ⬆️ Upgrade a Package

```console
ipm <package-manager> upgrade <package-name>
```

##### ⬆️ Upgrade all Packages

```console
ipm <package-manager> upgrade-all
```

##### 🗑️ Remove a package

```console
ipm <package-manager> uninstall <package-name>
```

#### 💡 Example

To install the `fast-json-stringify` package using `npm`:

```console
ipm npm install fast-json-stringify
```

<p align="right"><a href="#top">☝️</a></p>

## ⚙️ Configuration

`ipm` uses a JSON configuration file to define custom commands and settings for different package managers. The configuration file is located in the config directory.

### 🪞 Example Configuration

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

<p align="right"><a href="#top">☝️</a></p>

## 📄 Important Documents

- [Changelog](https://github.com/imfsiddiqui/ipm/blob/main/docs/CHANGELOG.md): Changelog of all notable changes.
- [Code of Conduct](https://github.com/imfsiddiqui/ipm/blob/main/docs/CODE-OF-CONDUCT.md): Code of Conduct for contributors.
- [Commit Message Instructions](https://github.com/imfsiddiqui/ipm/blob/main/.github/copilot/commit-message-instructions.md): Commit message guidelines for contributors and Copilot.
- [Contribution Guidelines](https://github.com/imfsiddiqui/ipm/blob/main/docs/CONTRIBUTING.md): How to contribute to this project.
- [License](https://github.com/imfsiddiqui/ipm/blob/main/LICENSE.md): License text.
- [Pull Request Description Instructions](https://github.com/imfsiddiqui/ipm/blob/main/.github/copilot/pull-request-description-instructions.md): Pull request guidelines for contributors and Copilot.
- [Roadmap](https://github.com/imfsiddiqui/ipm/blob/main/docs/ROADMAP.md): High-level strategic plan, long-term goals, milestones, and overall project vision.
- [Security Policy](https://github.com/imfsiddiqui/ipm/blob/main/docs/SECURITY.md): Security policy and reporting instructions.
- [Todo](https://github.com/imfsiddiqui/ipm/blob/main/docs/TODO.md): Day-to-day task tracking and immediate execution.

<p align="right"><a href="#top">☝️</a></p>

## 🙏 Acknowledgements

I would like to extend my heartfelt thanks to all the developers and contributors whose work has made this project possible. Your dedication and contributions to the open-source community are invaluable, and I am grateful for the tools, libraries, and frameworks that you have created and maintained.

### 🌟 Special Thanks

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

<p align="right"><a href="#top">☝️</a></p>

## 📜 License

This project is licensed under the [MIT License](https://github.com/imfsiddiqui/ipm/blob/main/LICENSE.md), allowing anyone to use, modify, and distribute it freely for personal or commercial purposes.

<p align="right"><a href="#top">☝️</a></p>
