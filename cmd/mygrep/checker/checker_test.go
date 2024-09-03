package checker

import (
	"testing"
)

func TestCheckPatternMatch(t *testing.T) {
	type args struct {
		line    []byte
		pattern *string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Simple match",
			args: args{
				line:    []byte("hello"),
				pattern: strPtr("h"),
			},
			want: true,
		},
		{
			name: "No match",
			args: args{
				line:    []byte("hello"),
				pattern: strPtr("x"),
			},
			want: false,
		},
		{
			name: "Word character match",
			args: args{
				line:    []byte("hello123"),
				pattern: strPtr("\\w"),
			},
			want: true,
		},
		{
			name: "Digit match",
			args: args{
				line:    []byte("hello123"),
				pattern: strPtr("\\d"),
			},
			want: true,
		},
		{
			name: "Multiple character match",
			args: args{
				line:    []byte("hello"),
				pattern: strPtr("he"),
			},
			want: true,
		},
		{
			name: "Partial match (should fail)",
			args: args{
				line:    []byte("hello"),
				pattern: strPtr("hx"),
			},
			want: false,
		},
		{
			name: "Empty pattern",
			args: args{
				line:    []byte("hello"),
				pattern: strPtr(""),
			},
			want: true,
		},
		{
			name: "Empty line",
			args: args{
				line:    []byte{},
				pattern: strPtr("h"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPatternMatch(tt.args.line, tt.args.pattern); got != tt.want {
				t.Errorf("CheckPatternMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to convert string to *string
func strPtr(s string) *string {
	return &s
}
