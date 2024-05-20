package myip

import (
	_ "embed"
	"runtime/debug"
)

//go:embed LICENSE
var license string

// License returns the license text for this program.
func License() string {
	return license
}

// Version returns a string based on module (vcs) information.
func Version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}

	return "unknown"
}
