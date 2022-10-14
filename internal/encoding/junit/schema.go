package junit

import "encoding/xml"

type JUnitReport struct {
	XMLName    xml.Name    `xml:"testsuites"`
	Disabled   int         `xml:"disabled,attr"`
	Errors     int         `xml:"errors,attr"`
	Failures   int         `xml:"failures,attr"`
	Tests      int         `xml:"tests,attr"`
	Time       string      `xml:"time,attr"`
	TestSuites []TestSuite `xml:"testsuite"`
}

type TestSuite struct {
	XMLName    xml.Name   `xml:"testsuite"`
	Name       string     `xml:"name,attr"`
	Disabled   int        `xml:"disabled,attr"`
	Errors     int        `xml:"errors,attr"`
	Failures   int        `xml:"failures,attr"`
	Skipped    int        `xml:"skipped,attr"`
	Tests      int        `xml:"tests,attr"`
	Time       string     `xml:"time,attr"`
	Properties []Property `xml:"property"`
	TestCases  []TestCase `xml:"testcase"`
}

type Property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type TestCase struct {
	XMLName   xml.Name `xml:"testcase"`
	Name      string   `xml:"name,attr"`
	ClassName string   `xml:"classname,attr"`
	Skipped   Skipped  `xml:"skipped,omitempty"`
	Failure   Failure  `xml:"failure,omitempty"`
}

type Failure struct {
	Type        string `xml:"type,attr"`
	Message     string `xml:"message,attr"`
	Description string `xml:",chardata"`
}

type Skipped struct {
	Type        string `xml:"type,attr"`
	Message     string `xml:"message,attr"`
	Description string `xml:",chardata"`
}
