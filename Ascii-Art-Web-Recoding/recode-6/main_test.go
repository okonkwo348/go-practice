package main

import (
	"testing"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{"Valid normal string", "hello", false},
		{"strings with space", "hello world", false},
		{"valid string with number", "hello123", false},
		{"string with invalid character", "\x01", true},
		{"empty string", "", false},
		{"string with \\n seperator", "hello\nworld", false},
		{"string with special character", "hello!!", false},
	}

	for _, tt := range tests {
		_, err := ValidateInput(tt.input)
		if tt.expectErr && err == nil {
			t.Errorf("%s: expected error but got none", tt.name)
		}

		if !tt.expectErr && err != nil {
			t.Errorf("%v: didn't expect an err but go %s", tt.name, err)
		}
	}
}
