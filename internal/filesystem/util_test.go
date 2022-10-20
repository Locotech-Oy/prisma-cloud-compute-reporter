package filesystem_test

import (
	"os"
	"testing"

	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/filesystem"
)

func TestGetAbsPath(t *testing.T) {

	t.Run("Returns workingdir only if path empty", func(t *testing.T) {
		rtn := filesystem.GetAbsPath("", "/tmp")
		if rtn != "/tmp" {
			t.Errorf("Expecting return value to be /tmp")
		}
	})

	t.Run("Returns path directly if absolute", func(t *testing.T) {
		rtn := filesystem.GetAbsPath("/tmp/my/path", "/tmp")
		if rtn != "/tmp/my/path" {
			t.Errorf("Expecting return value to be /tmp/my/path")
		}
	})

	t.Run("Concatenates relative path with working dir", func(t *testing.T) {
		rtn := filesystem.GetAbsPath("path", "/tmp")
		if rtn != "/tmp/path" {
			t.Errorf("Expecting return value to be /tmp/path")
		}

		rtn = filesystem.GetAbsPath("./path", "/tmp")
		if rtn != "/tmp/path" {
			t.Errorf("Expecting return value to be /tmp/path")
		}
	})
}

func TestPathIsDir(t *testing.T) {
	t.Run("Returns false for non-existing path", func(t *testing.T) {
		rtn := filesystem.PathIsDir("/does/not/exist")
		if rtn {
			t.Errorf("Expecting return value to be false for /does/not/exist")
		}
	})

	t.Run("Returns false for non-existing path", func(t *testing.T) {
		rtn := filesystem.PathIsDir(os.TempDir())
		if !rtn {
			t.Errorf("Expecting return value to be true for %s", os.TempDir())
		}
	})

}
