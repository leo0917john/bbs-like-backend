package version

import (
	"fmt"
	"runtime"
)

var (
	// Version number for git tag
	Version string = ""
	// BuildDate
	BuildDate string = ""
)

func PrintCLIVersion() string {
	return fmt.Sprintf(
		"version %s, built on %s, %s",
		Version,
		BuildDate,
		runtime.Version(),
	)
}
