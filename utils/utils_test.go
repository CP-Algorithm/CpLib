package utils_test

import (
	"testing"

	"github.com/CP-Algorithm/CpLib/utils"
)

func TestEmailValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "empty email",
			email:    "",
			expected: false,
		},
		{
			name:     "double `.`",
			email:    "test..mail@gmail.com",
			expected: false,
		},
		{
			name:     "without `.` in suffinx",
			email:    "test@test",
			expected: false,
		},
		{
			name:     "without `@` symbol",
			email:    "hello_mail",
			expected: false,
		},
		{
			name:     "correct email",
			email:    "valid_email.is.here@domain.com",
			expected: true,
		},
		{
			name:     "single `@` character",
			email:    "@",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := utils.IsEmail(test.email)
			if res != test.expected {
				t.Errorf("expected %#v, got %#v", test.expected, res)
			}
		})
	}
}

func TestFibFunction(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    int
		expected uint64
	}{
		{
			name:     "zero value",
			input:    0,
			expected: 0,
		},
		{
			name:     "invalid input (less than zero)",
			input:    -10,
			expected: 0,
		},
		{
			name:     "First value",
			input:    1,
			expected: 1,
		},
		{
			name:     "Second value",
			input:    2,
			expected: 1,
		},
		{
			name:     "Third value",
			input:    3,
			expected: 2,
		},
		{
			name:     "Fourth value",
			input:    4,
			expected: 3,
		},
		{
			name:     "large input (85)",
			input:    85,
			expected: 259695496911122585,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := utils.FibOfn(test.input)
			if res != test.expected {
				t.Errorf("expected %#v, got %#v", test.expected, res)
			}
		})
	}
}
