package junit

import (
	"strconv"
	"strings"

	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/parser"
)

func EncodeScanReport(report parser.ScanReport) JUnitReport {

	var testSuites []TestSuite

	complianceSuite := createComplianceSuiteWithProps(report.Results[0])

	for _, comp := range report.Results[0].Compliances {

		if report.Results[0].ComplianceScanPassed {
			complianceSuite.TestCases = append(complianceSuite.TestCases, createSkippedTestCase(comp))
		} else {
			complianceSuite.TestCases = append(complianceSuite.TestCases, createFailureTestCase(comp))
		}

	}

	vulnerabilitySuite := createVulnerabilitySuiteWithProps(report.Results[0])

	testSuites = append(testSuites, complianceSuite, vulnerabilitySuite)

	rtn := JUnitReport{
		TestSuites: testSuites,
	}

	return rtn
}

func createFailureTestCase(comp parser.Complicance) TestCase {
	className := "[" + comp.Category + "_" + strconv.FormatInt(int64(comp.Id), 10) + "]"

	return TestCase{
		Name:      className + " " + comp.Title,
		ClassName: className,
		Failure: Failure{
			Type:    "failure",
			Message: comp.Title,
			Description: `Severity: ` + comp.Severity + `
Category` + comp.Category + `
Description` + comp.Description,
		},
	}
}

func createSkippedTestCase(comp parser.Complicance) TestCase {
	className := "[" + comp.Category + "_" + strconv.FormatInt(int64(comp.Id), 10) + "]"

	return TestCase{
		Name:      className + " " + comp.Title,
		ClassName: className,
		Skipped: Skipped{
			Type:    "skipped",
			Message: comp.Title,
			Description: `Severity: ` + comp.Severity + `
Category` + comp.Category + `
Description` + comp.Description,
		},
	}
}

func createComplianceSuiteWithProps(result parser.Result) TestSuite {
	suite := createTestSuiteWithProps(result, "Prisma Cloud compliance scan")
	suite.Properties = append(suite.Properties,
		Property{
			Name:  "complianceScanPassed",
			Value: strconv.FormatBool(result.ComplianceScanPassed),
		},
		Property{
			Name:  "complianceDistribution_critical",
			Value: strconv.FormatInt(int64(result.ComplianceDistribution.Critical), 10),
		},
		Property{
			Name:  "complianceDistribution_high",
			Value: strconv.FormatInt(int64(result.ComplianceDistribution.High), 10),
		},
		Property{
			Name:  "complianceDistribution_medium",
			Value: strconv.FormatInt(int64(result.ComplianceDistribution.Medium), 10),
		},
		Property{
			Name:  "complianceDistribution_low",
			Value: strconv.FormatInt(int64(result.ComplianceDistribution.Low), 10),
		},
		Property{
			Name:  "complianceDistribution_total",
			Value: strconv.FormatInt(int64(result.ComplianceDistribution.Total), 10),
		},
	)
	return suite
}

func createVulnerabilitySuiteWithProps(result parser.Result) TestSuite {
	suite := createTestSuiteWithProps(result, "Prisma Cloud vulnerability scan")
	suite.Properties = append(suite.Properties,
		Property{
			Name:  "vulnerabilityScanPassed",
			Value: strconv.FormatBool(result.VulnerabilityScanPassed),
		},
		Property{
			Name:  "vulnerabilityDistribution_critical",
			Value: strconv.FormatInt(int64(result.VulnerabilityDistribution.Critical), 10),
		},
		Property{
			Name:  "vulnerabilityDistribution_high",
			Value: strconv.FormatInt(int64(result.VulnerabilityDistribution.High), 10),
		},
		Property{
			Name:  "vulnerabilityDistribution_medium",
			Value: strconv.FormatInt(int64(result.VulnerabilityDistribution.Medium), 10),
		},
		Property{
			Name:  "vulnerabilityDistribution_low",
			Value: strconv.FormatInt(int64(result.VulnerabilityDistribution.Low), 10),
		},
		Property{
			Name:  "vulnerabilityDistribution_total",
			Value: strconv.FormatInt(int64(result.VulnerabilityDistribution.Total), 10),
		},
	)
	return suite
}

func createTestSuiteWithProps(result parser.Result, name string) TestSuite {

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
		Name:       name,
		Properties: properties,
	}

}
