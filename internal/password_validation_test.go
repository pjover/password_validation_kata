package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name       string
		validation Validation
		password   string
		want       bool
	}{
		{
			name:       "8 characters",
			validation: One,
			password:   "Ab34567_",
			want:       true,
		},
		{
			name:       "7 characters",
			validation: One,
			password:   "1234",
			want:       false,
		},
		{
			name:       "With capital letter",
			validation: One,
			password:   "Test_cap1tal",
			want:       true,
		},
		{
			name:       "Without capital letter",
			validation: One,
			password:   "test_capital",
			want:       false,
		},
		{
			name:       "With lowercase letter",
			validation: One,
			password:   "Test_capi4al",
			want:       true,
		},
		{
			name:       "Without lowercase letter",
			validation: One,
			password:   "AAAAAAAAA",
			want:       false,
		},
		{
			name:       "Without lowercase letter",
			validation: One,
			password:   "AAAAAAAAA",
			want:       false,
		},
		{
			name:       "More than 6 chars",
			validation: Two,
			password:   "Ab3ffwf",
			want:       true,
		},
		{
			name:       "6 or less chars",
			validation: Two,
			password:   "Ab3fff",
			want:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePassword(tt.validation, tt.password)
			assert.Equal(t, tt.want, got)
		})
	}
}
