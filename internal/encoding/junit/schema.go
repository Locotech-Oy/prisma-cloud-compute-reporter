package junit

import (
	"encoding/xml"
)

type JUnitReport struct {
	XMLName xml.Name `xml:"testsuites"`
	//Disabled   int         `xml:"disabled,attr"`
	//Errors     int         `xml:"errors,attr"`
	//Failures   int         `xml:"failures,attr"`
	//Time       string      `xml:"time,attr"`
	//Tests      int         `xml:"tests,attr"`
	TestSuites []TestSuite `xml:"testsuite"`
}

type TestSuite struct {
	XMLName xml.Name `xml:"testsuite"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	//Disabled   int        `xml:"disabled,attr"`
	Errors     int        `xml:"errors,attr"`
	Failures   int        `xml:"failures,attr"`
	Skipped    int        `xml:"skipped,attr"`
	Tests      int        `xml:"tests,attr"`
	Hostname   string     `xml:"hostname,attr"`
	Time       float64    `xml:"time,attr"`
	Timestamp  string     `xml:"timestamp,attr"`
	Package    string     `xml:"package,attr"`
	Properties []Property `xml:"properties>property"`
	TestCases  []TestCase `xml:"testcase"`
	SystemOut  string     `xml:"system-out"`
	SystemErr  string     `xml:"system-err"`
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
	Time      float64  `xml:"time,attr"`
	Skipped   *Skipped `xml:"skipped,omitempty"`
	Failure   *Failure `xml:"failure,omitempty"`
}

type Failure struct {
	Message     string `xml:"message,attr,omitempty"`
	Description string `xml:",chardata"`
	Type        string `xml:"type,attr"`
}

type Skipped struct {
	Message     string `xml:"message,attr,omitempty"`
	Description string `xml:",chardata"`
}
