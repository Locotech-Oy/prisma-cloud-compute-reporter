package junit_test

import (
	"testing"
	"time"

	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/encoding/junit"
	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestEncodeScanReport(t *testing.T) {

	t.Run("Returns JUnitReport instance", func(t *testing.T) {

		rpt := parser.ScanReport{
			Results: []parser.Result{
				{
					Id:                   "result/id",
					Name:                 "result_name",
					Distro:               "some distro 11",
					DistroRelease:        "distro release",
					ComplianceScanPassed: true,
					ComplianceDistribution: parser.Distribution{
						Critical: 0,
						High:     1,
						Medium:   2,
						Low:      3,
						Total:    6,
					},
					Collections:             []string{"coll1", "coll2"},
					ScanTime:                time.Date(2022, 10, 25, 13, 56, 59, 0, &time.Location{}),
					VulnerabilityScanPassed: false,
					VulnerabilityDistribution: parser.Distribution{
						Critical: 3,
						High:     2,
						Medium:   1,
						Low:      0,
						Total:    6,
					},
				},
			},
		}

		rtn := junit.EncodeScanReport(rpt)

		assert.NotNil(t, rtn, "rtn should not be nil")
		assert.IsType(t, junit.JUnitReport{}, rtn)
		assert.Equal(t, 2, len(rtn.TestSuites))
		assert.Equal(t, junit.TestSuite{
			Id:        1,
			Name:      "Prisma Cloud compliance scan",
			Hostname:  "localhost",
			Timestamp: "2022-10-25T13:56:59",
			Properties: []junit.Property{
				{
					Name:  "id",
					Value: "result/id",
				},
				{
					Name:  "name",
					Value: "result_name",
				},
				{
					Name:  "distro",
					Value: "some distro 11",
				},
				{
					Name:  "distroRelease",
					Value: "distro release",
				},
				{
					Name:  "collections",
					Value: "coll1, coll2",
				},
				{
					Name:  "complianceScanPassed",
					Value: "true",
				},
				{
					Name:  "complianceDistribution_critical",
					Value: "0",
				},
				{
					Name:  "complianceDistribution_high",
					Value: "1",
				},
				{
					Name:  "complianceDistribution_medium",
					Value: "2",
				},
				{
					Name:  "complianceDistribution_low",
					Value: "3",
				},
				{
					Name:  "complianceDistribution_total",
					Value: "6",
				},
			},
		}, rtn.TestSuites[0], "Compliance result structure not as expected")

		assert.Equal(t, junit.TestSuite{
			Id:        2,
			Name:      "Prisma Cloud vulnerability scan",
			Hostname:  "localhost",
			Timestamp: "2022-10-25T13:56:59",
			Properties: []junit.Property{
				{
					Name:  "id",
					Value: "result/id",
				},
				{
					Name:  "name",
					Value: "result_name",
				},
				{
					Name:  "distro",
					Value: "some distro 11",
				},
				{
					Name:  "distroRelease",
					Value: "distro release",
				},
				{
					Name:  "collections",
					Value: "coll1, coll2",
				},
				{
					Name:  "vulnerabilityScanPassed",
					Value: "false",
				},
				{
					Name:  "vulnerabilityDistribution_critical",
					Value: "3",
				},
				{
					Name:  "vulnerabilityDistribution_high",
					Value: "2",
				},
				{
					Name:  "vulnerabilityDistribution_medium",
					Value: "1",
				},
				{
					Name:  "vulnerabilityDistribution_low",
					Value: "0",
				},
				{
					Name:  "vulnerabilityDistribution_total",
					Value: "6",
				},
			},
		}, rtn.TestSuites[1], "Vulnerability result structure not as expected")
	})
}
