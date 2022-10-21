package junit

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/encoding/junit"
	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/filesystem"
	"github.com/rs/zerolog/log"
)

func Write(outputPath string, report junit.JUnitReport) error {
	log.Debug().Msg("report output as junit")

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if outputPath == "" {
		outputPath = workingDir
	}

	outputPath = filesystem.GetAbsPath(outputPath, workingDir)

	// if output points to folder, assume default filename for writing
	if filesystem.PathIsDir(outputPath) {
		outputPath = filepath.Join(outputPath, "pcc-junit-report.xml")
	}

	// assume output to file, ensure path is writeable
	outputFile, err := os.OpenFile(filepath.Clean(outputPath), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer func() {
		if err := outputFile.Close(); err != nil {
			log.Error().AnErr("error", err).Msg("Failed to close file")
		}
	}()

	log.Info().
		Str("output_path", outputPath).
		Msg(fmt.Sprintf("Writing JUnit report to path %s", outputPath))
	out, err := xml.MarshalIndent(report, "", "  ")

	if err != nil {
		return err
	}
	_, err = outputFile.WriteString(xml.Header + string(out))
	if err != nil {
		return err
	}

	return nil
}
