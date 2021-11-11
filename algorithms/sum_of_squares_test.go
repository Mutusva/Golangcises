package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDifference(t *testing.T) {
	
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Test 01",
			n: 1,
			want: 0,
		},
		{
			name: "Test 02",
			n: 2,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Difference(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}

 func BenchmarkDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Difference(3)
	}
}