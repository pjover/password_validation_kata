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
		{
			name:       "Lacks capital letter",
			validation: Two,
			password:   "ab3ffwf",
			want:       false,
		},
		{
			name:       "Lacks lowercase",
			validation: Two,
			password:   "AB3FFFH",
			want:       false,
		},
		{
			name:       "Lacks number",
			validation: Two,
			password:   "ABxFFFH",
			want:       false,
		},
		{
			name:       "More than 16 chars",
			validation: Three,
			password:   "Abff_wfas_dñlk_fj",
			want:       true,
		},
		{
			name:       "Lacks uppercase",
			validation: Three,
			password:   "abff_wfas_dhlk_fj",
			want:       false,
		},
		{
			name:       "Lacks lowercase",
			validation: Three,
			password:   "ABBB_BBBB_BBBB_BB",
			want:       false,
		},
		{
			name:       "Lacks underscore",
			validation: Three,
			password:   "Abfffwfasfdñlkffj",
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
