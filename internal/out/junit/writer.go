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

	if outputPath == "" {
		var err error
		outputPath, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	// if output points to folder, assume default filename for writing
	if filesystem.PathIsDir(outputPath) {
		outputPath = filepath.Join(outputPath, "junit-report.xml")
	}

	// assume output to file, ensure path is writeable
	outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0744)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	log.Info().
		Str("output_path", outputPath).
		Msg(fmt.Sprintf("Writing JUnit report to path %s", outputPath))
	out, _ := xml.MarshalIndent(report, "", "  ")
	outputFile.WriteString(xml.Header + string(out))

	return nil
}
