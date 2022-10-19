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

	for _, vuln := range report.Results[0].Vulnerabilities {

		if report.Results[0].VulnerabilityScanPassed {
			vulnerabilitySuite.TestCases = append(vulnerabilitySuite.TestCases, createSkippedVTestCase(vuln))
		} else {
			vulnerabilitySuite.TestCases = append(vulnerabilitySuite.TestCases, createFailureVTestCase(vuln))
		}

	}

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
		Time:      0,
		Failure: &Failure{
			Type:    "compliance",
			Message: comp.Title,
			Description: `Severity: ` + comp.Severity + `
Category: ` + comp.Category + `
Description: ` + comp.Description,
		},
	}
}

func createSkippedTestCase(comp parser.Complicance) TestCase {
	className := "[" + comp.Category + "_" + strconv.FormatInt(int64(comp.Id), 10) + "]"

	return TestCase{
		Name:      className + " " + comp.Title,
		ClassName: className,
		Time:      0,
		Skipped: &Skipped{
			Message: comp.Title,
			Description: `Severity: ` + comp.Severity + `
Category: ` + comp.Category + `
Description: ` + comp.Description,
		},
	}
}

func createFailureVTestCase(vulnerability parser.Vulnerability) TestCase {
	className := "[" + vulnerability.Id + "]"

	return TestCase{
		Name:      className + " " + vulnerability.Description,
		ClassName: vulnerability.PackageName + "_" + vulnerability.PackageVersion,
		Time:      0,
		Failure: &Failure{
			Type:    "vulnerability",
			Message: vulnerability.Description,
			Description: `Severity: ` + vulnerability.Severity + `
Description: ` + vulnerability.Description,
		},
	}
}

func createSkippedVTestCase(vulnerability parser.Vulnerability) TestCase {
	className := "[" + vulnerability.Id + "]"

	return TestCase{
		Name:      className + " " + vulnerability.Description,
		ClassName: vulnerability.PackageName + "_" + vulnerability.PackageVersion,
		Time:      0,
		Skipped: &Skipped{
			Message: vulnerability.Description,
			Description: `Severity: ` + vulnerability.Severity + `
Description: ` + vulnerability.Description,
		},
	}
}

func createComplianceSuiteWithProps(result parser.Result) TestSuite {
	suite := createTestSuiteWithProps(1, result, "Prisma Cloud compliance scan")
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
	suite := createTestSuiteWithProps(2, result, "Prisma Cloud vulnerability scan")
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
