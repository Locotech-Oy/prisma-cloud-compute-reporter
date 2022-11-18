/*
Copyright © 2022 Locotech Oy <jens.wegar@locotech.fi>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	junitencoder "github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/encoding/junit"
	junitwriter "github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/out/junit"
	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/parser"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse [flags] report-file.json",
	Short: "Parse a twistcli scan report json file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().
			Str("scan_report_file", args[0]).
			Msg(fmt.Sprintf("Reading scan report input from %s", args[0]))

		f, err := os.Open(filepath.Clean(args[0]))

		if err != nil {
			log.Error().
				Err(err).
				Msg("Could not read input file, exiting...")
			os.Exit(1)
		}

		report, err := parser.ParseJSON(f)

		if err != nil {
			log.Error().
				Err(err).
				Msg("Could not parse input file, exiting...")
			os.Exit(1)
		}

		outputPath, _ := cmd.Flags().GetString("output-file")
		includeCompliance, _ := cmd.Flags().GetBool("compliance")
		includeVulnerability, _ := cmd.Flags().GetBool("vulnerability")

		outputReport := junitencoder.JUnitReport{}
		if includeCompliance {
			log.Debug().
				Msg("Including compliance report")

			outputReport.TestSuites = append(outputReport.TestSuites, junitencoder.EncodeComplianceReport(report).TestSuites...)
		}
		if includeVulnerability {
			log.Debug().
				Msg("Including vulnerability report")

			outputReport.TestSuites = append(outputReport.TestSuites, junitencoder.EncodeVulnerabilityReport(report).TestSuites...)
		}
		if !includeCompliance && !includeVulnerability {
			log.Debug().
				Msg("Including all report sections")
			outputReport = junitencoder.EncodeScanReport(report)
		}

		if len(outputReport.TestSuites) > 0 {
			// Export to file based on output format
			err = junitwriter.Write(outputPath, outputReport)
			if err != nil {
				log.Error().
					Err(err).
					Msg("Could not parse input file, exiting...")
				os.Exit(1)
			}
		} else {
			log.Error().
				Err(err).
				Msg("No reports included in output, exiting...")
			os.Exit(1)
		}

	},
	Args: cobra.ExactArgs(1),
}

func init() {
	imageCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringP("output-file", "o", "", "Destination file path for parse output")
	// TODO: Validate supported formats
	parseCmd.Flags().StringP("output-format", "f", "junit", "Format of output file")
	parseCmd.Flags().Bool("compliance", false, "Include compliance results in output")
	parseCmd.Flags().Bool("vulnerability", false, "Include vulnerability results in output")
}
