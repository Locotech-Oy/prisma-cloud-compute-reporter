package version_test

import (
	"testing"

	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/version"
	"github.com/stretchr/testify/assert"
)

func TestVersionStr(t *testing.T) {

	t.Run("Returns string containing version number", func(t *testing.T) {
		str := version.VersionStr()

		assert.NotNil(t, str, "str should not be nil")
		assert.Contains(t, str, "0.0.0", "str should contain version nr")
	})
}
