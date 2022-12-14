package filesystem

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

// Builds an absolute path
//
// If path already is absolute, the path is returned unaltered.
// If path is relative, path is concatenated with workingDir and returned as absolute.
// If path is empty, workingDir is returned.
func GetAbsPath(path string, workingDir string) string {

	if path != "" && filepath.IsAbs(path) {
		return path
	} else if path != "" {
		return filepath.Join(workingDir, path)
	}

	return workingDir
}

// Checks if the given path points to a directory
func PathIsDir(path string) bool {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		log.Error().AnErr("error", err).Msg("Error opening path")
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Error().AnErr("error", err).Msg("Failed to close file")
			os.Exit(1)
		}
	}()

	fi, err := file.Stat()
	if err != nil {
		log.Error().AnErr("error", err).Msg("Error getting file stats")
		os.Exit(1)
	}

	return fi.IsDir()
}
