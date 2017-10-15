package main

import (
	"net/url"
	"testing"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestFormatValues(t *testing.T) {
	want := []string{
		"v := url.Values{",
		`	"T1": {"V1"},`,
		`	"T2": {"V2","V3"},`,
		"}",
	}

	v := url.Values{
		"T1": {"V1"},
		"T2": {"V2", "V3"},
	}

	result := formatValues(v)
	count := 0
	for line := range result {
		if !contains(want, line) {
			t.Errorf("expected line %v is missing", line)
		}
		count++
	}

	if count != len(want) {
		t.Errorf("Generated string list is not the same as expected list.")
	}
}
