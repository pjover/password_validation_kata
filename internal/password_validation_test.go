package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{
			name:     "8 characters",
			password: "Ab345678",
			want:     true,
		},
		{
			name:     "7 characters",
			password: "1234",
			want:     false,
		},
		{
			name:     "With capital letter",
			password: "Test_capital",
			want:     true,
		},
		{
			name:     "Without capital letter",
			password: "test_capital",
			want:     false,
		},
		{
			name:     "With lowercase letter",
			password: "Test_capital",
			want:     true,
		},
		{
			name:     "Without lowercase letter",
			password: "AAAAAAAAA",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePassword(tt.password)
			assert.Equal(t, tt.want, got)
		})
	}
}
