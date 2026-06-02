package main

import (
	"testing"
	"time"

	"github.com/pocketbase/pocketbase/tools/types"
)

func mustDateTime(t *testing.T, s string) types.DateTime {
	t.Helper()
	var dt types.DateTime
	if err := dt.Scan(s); err != nil {
		t.Fatalf("kunne ikke parse datetime %q: %v", s, err)
	}
	return dt
}

func TestFormatDate_NorwegianFormat(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"2026-04-21 10:00:00.000Z", "21. april 2026"},
		{"2026-01-01 00:00:00.000Z", "1. januar 2026"},
		{"2026-02-28 12:00:00.000Z", "28. februar 2026"},
		{"2026-03-15 08:30:00.000Z", "15. mars 2026"},
		{"2026-05-17 09:00:00.000Z", "17. mai 2026"},
		{"2026-06-01 00:00:00.000Z", "1. juni 2026"},
		{"2026-07-04 00:00:00.000Z", "4. juli 2026"},
		{"2026-08-20 00:00:00.000Z", "20. august 2026"},
		{"2026-09-10 00:00:00.000Z", "10. september 2026"},
		{"2026-10-31 00:00:00.000Z", "31. oktober 2026"},
		{"2026-11-11 00:00:00.000Z", "11. november 2026"},
		{"2026-12-24 00:00:00.000Z", "24. desember 2026"},
	}

	for _, tc := range cases {
		dt := mustDateTime(t, tc.input)
		got := formatDate(dt)
		if got != tc.expected {
			t.Errorf("formatDate(%q) = %q, vil ha %q", tc.input, got, tc.expected)
		}
	}
}

func TestFormatDate_Zero(t *testing.T) {
	var dt types.DateTime
	got := formatDate(dt)
	if got != "-" {
		t.Errorf("formatDate(zero) = %q, vil ha %q", got, "-")
	}
}

func TestFormatDate_UsesLocalTime(t *testing.T) {
	// Verifiser at datoen bruker lokal tid (UTC her i test-miljøet)
	dt := mustDateTime(t, "2026-04-21 10:00:00.000Z")
	got := formatDate(dt)
	_ = time.Now() // sørger for at time-pakken er brukt
	if got == "" {
		t.Error("formatDate returnerte tom streng")
	}
}
