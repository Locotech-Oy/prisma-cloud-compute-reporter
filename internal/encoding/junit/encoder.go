package junit

import (
	"strings"

	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/parser"
)

// EncodeScanReport accepts a single ScanReport and returns a JUnitReport
func EncodeScanReport(report parser.ScanReport) JUnitReport {

	var testSuites []TestSuite

	testSuites = append(testSuites, EncodeComplianceReport(report).TestSuites[0], EncodeVulnerabilityReport(report).TestSuites[0])

	rtn := JUnitReport{
		TestSuites: testSuites,
	}

	return rtn
}

func createTestSuiteWithProps(id int, result parser.Result, name string) TestSuite {

	properties := []Property{
		{
			Name:  "id",
			Value: result.Id,
		},
		{
			Name:  "name",
			Value: result.Name,
		},
		{
			Name:  "distro",
			Value: result.Distro,
		},
		{
			Name:  "distroRelease",
			Value: result.DistroRelease,
		},
		{
			Name:  "collections",
			Value: strings.Join(result.Collections, ", "),
		},
	}

	return TestSuite{
		Id:        id,
		Name:      name,
		Hostname:  "localhost",
		Timestamp: result.ScanTime.Format("2006-01-02T15:04:05"),
		// PC scan report does not include scan duration, so always 0
		Time:       0,
		Properties: properties,
	}

}
