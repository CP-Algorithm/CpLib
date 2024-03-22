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
