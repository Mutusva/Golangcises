package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShareWith(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Test 01",
			s: "Inno",
			want: "One for Inno, one for me.",
		},
		{
			name: "Test 02",
			want: "One for you, one for me.",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShareWith(tt.s)
			require.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkShareWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShareWith("Inno")
	}
}

func BenchmarkShareWithTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShareWithTwo("Inno")
	}
}

func BenchmarkShareWith03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShareWith03("Inno")
	}
}


