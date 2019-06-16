package acmdl

import (
	"net/url"
	"testing"
)

func TestQueryEncoding(t *testing.T) {
	tests := []struct {
		words    []string
		year     int
		start    int
		expected url.Values
	}{
		{
			[]string{"vmm"},
			-1,
			0,
			url.Values{"dte": []string{""}, "query": []string{"(vmm)"}, "start": []string{"0"}},
		},
		{
			[]string{"vmm", "+security"},
			-1,
			0,
			url.Values{"dte": []string{""}, "query": []string{"(vmm %2Bsecurity)"}, "start": []string{"0"}},
		},
		{
			[]string{"vmm", "-security"},
			2014,
			20,
			url.Values{"dte": []string{"2014"}, "query": []string{"(vmm -security)"}, "start": []string{"20"}},
		},
	}

	for _, tt := range tests {
		q := NewQuery()

		for _, word := range tt.words {
			q.AddWord(word)
		}

		q.Year = tt.year
		q.Start = tt.start

		if q.Encode() != tt.expected.Encode() {
			t.Errorf("encoded query should be %s. got %s", tt.expected.Encode(), q.Encode())
		}
	}
}
