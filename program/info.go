package program

import "runtime/debug"

const Name = "laravel-ls"

var VersionOverride = ""

func Version() string {
	if VersionOverride != "" {
		return VersionOverride
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "" {
			return info.Main.Version
		}
	}
	return "(unknown)"
}
