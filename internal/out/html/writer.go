package junit

import (
	"embed"
	"strings"

	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/filesystem"
	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/parser"
	"github.com/rs/zerolog/log"
)

//go:embed template/scanreport.html
var templateHtmlFiles embed.FS

func Write(outputPath string, report parser.ScanReport) error {
	log.Debug().Msg("report output as html")

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
		outputPath = filepath.Join(outputPath, "pcc-junit-report.html")
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
		Msg(fmt.Sprintf("Writing HTML report to path %s", outputPath))

	funcMap := template.FuncMap{
		"StringsJoin": strings.Join,
	}

	templateHtml, err := template.New("scanreport.html").Funcs(funcMap).ParseFS(templateHtmlFiles, "template/scanreport.html")
	if err != nil {
		return err
	}

	err = templateHtml.Execute(outputFile, report)
	if err != nil {
		return err
	}

	return nil
}
