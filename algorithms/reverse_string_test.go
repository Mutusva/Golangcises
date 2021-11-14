package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Test 01",
			s:    "cat",
			want: "tac",
		},
		{
			name: "Test 02",
			s:    "inno",
			want: "onni",
		},
		{
			name: "Test 03",
			s:    "tiktoc",
			want: "cotkit",
		},
		{
			name: "Test 04",
			s:    "s",
			want: "s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.s)
			require.Equal(t, tt.want, got)
		})
	}
}
