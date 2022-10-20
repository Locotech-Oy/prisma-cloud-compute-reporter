package parser_test

import (
	"strings"
	"testing"

	"github.com/Locotech-Oy/prisma-cloud-compute-reporter/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseJSON(t *testing.T) {

	t.Run("Returns error on failed parsing", func(t *testing.T) {
		_, err := parser.ParseJSON(strings.NewReader("This is not json"))

		assert.NotNil(t, err, "err should not be nil")
	})

	t.Run("Returns ScanReport", func(t *testing.T) {
		r, err := parser.ParseJSON(strings.NewReader("{\"results\":[{\"id\": \"12345\"}]}"))

		assert.Nil(t, err, "err should be nil")
		assert.Len(t, r.Results, 1, "results length should be 1")
	})

}
