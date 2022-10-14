package parser

import (
	"encoding/json"
	"io"
)

func ParseJSON(r io.Reader) (ScanReport, error) {

	d := json.NewDecoder(r)
	var report ScanReport

	if err := d.Decode(&report); err != nil {
		return ScanReport{}, err
	}

	return report, nil

}
