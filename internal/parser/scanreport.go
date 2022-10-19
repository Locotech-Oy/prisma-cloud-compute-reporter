package parser

import "time"

// ScanReport represents the root of the twistcli json report
type ScanReport struct {
	Results []Result `json:"results"`
}

type Result struct {
	Id                        string          `json:"id"`
	Name                      string          `json:"name"`
	Distro                    string          `json:"distro"`
	DistroRelease             string          `json:"distroRelease"`
	Collections               []string        `json:"collections"`
	Packages                  []Package       `json:"packages"`
	Applications              []Application   `json:"applications"`
	Compliances               []Complicance   `json:"compliances"`
	ComplianceDistribution    Distribution    `json:"complianceDistribution"`
	ComplianceScanPassed      bool            `json:"complianceScanPassed"`
	Vulnerabilities           []Vulnerability `json:"vulnerabilities"`
	VulnerabilityDistribution Distribution    `json:"vulnerabilityDistribution"`
	VulnerabilityScanPassed   bool            `json:"vulnerabilityScanPassed"`
	ScanTime                  time.Time       `json:"scanTime"`
}

type Package struct {
	Type     string   `json:"type"`
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Licenses []string `json:"licenses"`
}

type Application struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

type Distribution struct {
	Critical int `json:"critical"`
	High     int `json:"high"`
	Medium   int `json:"medium"`
	Low      int `json:"low"`
	Total    int `json:"total"`
}

type Complicance struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	LayerTime   string `json:"layerTime"`
	Category    string `json:"category"`
}

type Vulnerability struct {
	Id               string    `json:"id"`
	Severity         string    `json:"severity"`
	Status           string    `json:"status"`
	Cvss             float32   `json:"cvss"`
	Vector           string    `json:"vector"`
	Description      string    `json:"description"`
	PackageName      string    `json:"packageName"`
	PackageVersion   string    `json:"packageVersion"`
	Link             string    `json:"link"`
	RiskFactors      []string  `json:"riskFactors"`
	ImpactedVersions []string  `json:"impactedVersions"`
	PublishedDate    time.Time `json:"publishedDate"`
	DiscoveredDate   time.Time `json:"discoveredDate"`
	FixDate          time.Time `json:"fixDate"`
	LayerTime        time.Time `json:"layerTime"`
}
