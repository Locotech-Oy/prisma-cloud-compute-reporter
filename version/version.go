package version

import "fmt"

var version string = "dev"

// VersionStr returns a string containing build information about this application
func VersionStr() string {
	return fmt.Sprintf("Prisma Cloud Compute reporter\nversion: %s\n", version)
}
