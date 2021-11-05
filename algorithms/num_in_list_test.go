package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumInList(t *testing.T) {

	tests := []struct {
		name string
		list []int
		num  int
		want bool
	}{
		{
			name: "Test 01",
			list: []int{1, 3, 5, 6, 7},
			num:  3,
			want: true,
		},
		{
			name: "Test 02",
			list: []int{1, 4, 9, 6, 7, 12},
			num:  3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumInList(tt.list, tt.num)
			require.Equal(t, tt.want, got)
		})
	}
}
