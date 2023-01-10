package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name       string
		validation Validation
		password   string
		want       []error
	}{
		{
			name:       "8 characters",
			validation: One,
			password:   "Ab3456s7_",
			want:       []error{},
		},
		{
			name:       "7 characters",
			validation: One,
			password:   "Ul_234",
			want:       []error{},
		},
		{
			name:       "With capital letter",
			validation: One,
			password:   "Test_cap1tal",
			want:       []error{},
		},
		{
			name:       "Without capital letter",
			validation: One,
			password:   "test_capital0",
			want:       []error{},
		},
		{
			name:       "With lowercase letter",
			validation: One,
			password:   "Test_capi4al",
			want:       []error{},
		},
		{
			name:       "Without lowercase letter",
			validation: One,
			password:   "AA_0AAAAAA",
			want:       []error{},
		},
		{
			name:       "Without lowercase letter",
			validation: One,
			password:   "A_0AAAAAAAA",
			want:       []error{},
		},
		{
			name:       "Validation 3",
			validation: One,
			password:   "abfffwfasfdolksff0j",
			want: []error{
				errors.New("an upper case char is required"),
				errors.New("an underscore char is required"),
			},
		},
		{
			name:       "More than 6 chars",
			validation: Two,
			password:   "Ab3ffwf",
			want:       []error{},
		},
		{
			name:       "6 or less chars",
			validation: Two,
			password:   "Ab3fff",
			want:       []error{},
		},
		{
			name:       "Lacks capital letter",
			validation: Two,
			password:   "ab3ffwf",
			want:       []error{},
		},
		{
			name:       "Lacks lowercase",
			validation: Two,
			password:   "AB3FFFH",
			want:       []error{},
		},
		{
			name:       "Lacks number",
			validation: Two,
			password:   "ABxFFFH",
			want:       []error{},
		},
		{
			name:       "Validation 3",
			validation: Two,
			password:   "abfffwfasfdolksffj",
			want: []error{
				errors.New("an upper case char is required"),
				errors.New("a digit is required"),
			},
		},
		{
			name:       "More than 16 chars",
			validation: Three,
			password:   "Abff_wfas_dñlk_fj",
			want:       []error{},
		},
		{
			name:       "Lacks uppercase",
			validation: Three,
			password:   "aab0ff_wfas_dhlk_fsssj",
			want:       []error{},
		},
		{
			name:       "Lacks lowercase",
			validation: Three,
			password:   "ABBB_BBBB_BBBB_BB",
			want:       []error{},
		},
		{
			name:       "Lacks underscore",
			validation: Three,
			password:   "Abfffwfasfdolksff0j",
			want:       []error{},
		},
		{
			name:       "Validation 3",
			validation: Three,
			password:   "abfffwfasfdolksff0j",
			want: []error{
				errors.New("an upper case char is required"),
				errors.New("an underscore char is required"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePassword(tt.validation, tt.password)
			assert.Equal(t, tt.want, got)
		})
	}
}
