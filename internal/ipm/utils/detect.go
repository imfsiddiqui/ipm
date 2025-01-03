// Package utils provides utility functions for the application
package utils

import (
	"os/exec"
	"runtime"
)

// DetectDefaultPackageManager detects the default package manager based on the operating system.
//
// Returns:
//   - string: The name of the detected package manager, or an empty string if no package manager is detected.
//
// Example usage:
//
//	packageManager := DetectDefaultPackageManager()
//
// This function performs the following steps:
//  1. Checks the operating system using runtime.GOOS.
//  2. Calls the appropriate detection function based on the operating system.
//  3. Returns the name of the detected package manager, or an empty string if no package manager is detected.
func DetectDefaultPackageManager() string {
	switch runtime.GOOS {
	case "windows":
		return detectWindowsPackageManager()
	case "darwin":
		return detectMacOSPackageManager()
	case "linux":
		return detectLinuxPackageManager()
	}
	return ""
}

// detectWindowsPackageManager checks for the presence of winget, choco, and scoop on Windows.
//
// Returns:
//   - string: The name of the detected package manager, or an empty string if no package manager is detected.
//
// Example usage:
//
//	packageManager := detectWindowsPackageManager()
//
// This function performs the following steps:
//  1. Checks if the "winget" command is available.
//  2. If "winget" is not available, checks if the "choco" command is available.
//  3. If "choco" is not available, checks if the "scoop" command is available.
//  4. Returns the name of the detected package manager, or an empty string if no package manager is detected.
func detectWindowsPackageManager() string {
	if isCommandAvailable("winget") {
		return "winget"
	}
	if isCommandAvailable("choco") {
		return "choco"
	}
	if isCommandAvailable("scoop") {
		return "scoop"
	}
	return ""
}

// detectMacOSPackageManager checks for the presence of brew and macports on macOS.
//
// Returns:
//   - string: The name of the detected package manager, or an empty string if no package manager is detected.
//
// Example usage:
//
//	packageManager := detectMacOSPackageManager()
//
// This function performs the following steps:
//  1. Checks if the "brew" command is available.
//  2. If "brew" is not available, checks if the "port" command is available.
//  3. Returns the name of the detected package manager, or an empty string if no package manager is detected.
func detectMacOSPackageManager() string {
	if isCommandAvailable("brew") {
		return "brew"
	}
	if isCommandAvailable("port") {
		return "port"
	}
	return ""
}

// detectLinuxPackageManager checks for the presence of various package managers on Linux.
//
// Returns:
//   - string: The name of the detected package manager, or an empty string if no package manager is detected.
//
// Example usage:
//
//	packageManager := detectLinuxPackageManager()
//
// This function performs the following steps:
//  1. Checks if the "apk" command is available (for Alpine Linux).
//  2. If "apk" is not available, checks if the "apt" command is available (for Debian-based distributions) and so on.
//  3. Returns the name of the detected package manager, or an empty string if no package manager is detected.
func detectLinuxPackageManager() string {
	if isCommandAvailable("apk") { // apk: Alpine Linux
		return "apk"
	}
	if isCommandAvailable("apt") { // apt: Debian-based distributions
		return "apt"
	}
	if isCommandAvailable("cards") { // cards: NuTyX
		return "cards"
	}
	if isCommandAvailable("dnf") { // dnf: Fedora
		return "dnf"
	}
	if isCommandAvailable("emerge") { // emerge: Gentoo
		return "emerge"
	}
	if isCommandAvailable("eopkg") { // eopkg: Solus
		return "eopkg"
	}
	if isCommandAvailable("guix") { // guix: GNU Guix
		return "guix"
	}
	if isCommandAvailable("nix-env") { // nix-env: Nix
		return "nix-env"
	}
	if isCommandAvailable("opkg") { // opkg: OpenWrt
		return "opkg"
	}
	if isCommandAvailable("pacman") { // pacman: Arch Linux
		return "pacman"
	}
	if isCommandAvailable("slackpkg") { // slackpkg: Slackware
		return "slackpkg"
	}
	if isCommandAvailable("xbps-install") { // xbps-install: Void Linux
		return "xbps"
	}
	if isCommandAvailable("yum") { // yum: CentOS/RHEL
		return "yum"
	}
	if isCommandAvailable("zypper") { // zypper: openSUSE
		return "zypper"
	}

	// Additional checks if none of the above tools are available
	if isCommandAvailable("brew") { // brew: Homebrew on Linux
		return "brew"
	}
	if isCommandAvailable("flatpak") { // flatpak: Flatpak
		return "flatpak"
	}
	if isCommandAvailable("nala") { // nala: Nala (alternative frontend for APT)
		return "nala"
	}
	if isCommandAvailable("snap") { // snap: Snap
		return "snap"
	}

	return ""
}

// isCommandAvailable checks if a command is available on the system.
//
// Parameters:
//   - cmd: The name of the command to check.
//
// Returns:
//   - bool: True if the command is available, false otherwise.
//
// Example usage:
//
//	available := isCommandAvailable("brew")
//
// This function performs the following steps:
//  1. Uses exec.LookPath to check if the command is available in the system's PATH.
//  2. Returns true if the command is found, false otherwise.
func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
