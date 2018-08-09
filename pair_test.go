package main

import (
	"testing"
)

func TestCheckArgs(t *testing.T) {
	t.Run("rejects invalid arguments", func(t *testing.T) {
		knownInvalid := []struct {
			Description string
			Args        []string
		}{
			{"missing command", []string{"pair"}},
			{"unknown command", []string{"pair", "roneous"}},
			{"too many arguments for 'with' command", []string{"pair", "with", "mutiple", "people"}},
			{"not enough arguments for 'with' command", []string{"pair", "with"}},
			{"too many arguments for 'stop' command", []string{"pair", "stop", "now"}},
		}

		for _, tc := range knownInvalid {
			t.Run(tc.Description, func(t *testing.T) {
				if err := checkArgs(tc.Args); err == nil {
					t.Errorf("Expected '%s' to be invalid arguments, but were found valid", tc.Args)
				}
			})
		}
	})

	t.Run("accepts valid arguments", func(t *testing.T) {
		knownInvalid := []struct {
			Description string
			Args        []string
		}{
			{"'with' command", []string{"pair", "with", "me"}},
			{"'stop' command", []string{"pair", "stop"}},
		}

		for _, tc := range knownInvalid {
			t.Run(tc.Description, func(t *testing.T) {
				if err := checkArgs(tc.Args); err != nil {
					t.Errorf("Expected '%s' to be valid arguments, but were found invalid", tc.Args)
				}
			})
		}
	})
}
