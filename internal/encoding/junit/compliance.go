package junit

import (
	"strconv"

	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/parser"
)

// EncodeComplianceReport accepts a single ScanReport and returns a JUnitReport for compliances
func EncodeComplianceReport(report parser.ScanReport) JUnitReport {
	var testSuites []TestSuite

	complianceSuite := createComplianceSuiteWithProps(report.Results[0])
	for _, comp := range report.Results[0].Compliances {

		if report.Results[0].ComplianceScanPassed {
			complianceSuite.TestCases = append(complianceSuite.TestCases, createSkippedTestCase(comp))
		} else {
			complianceSuite.TestCases = append(complianceSuite.TestCases, createFailureTestCase(comp))
		}

	}

	testSuites = append(testSuites, complianceSuite)
	return JUnitReport{
		TestSuites: testSuites,
	}

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
