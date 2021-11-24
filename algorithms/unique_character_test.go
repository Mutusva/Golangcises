package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueCharacter(t *testing.T){
	test := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Test 01",
			s:    "statistics",
			want: 3,
		},
		{
			name: "Test 02",
			s:    "haeck",
			want: 1,
		},
	}

	for _,tc := range test {
		t.Run(tc.name, func (t *testing.T){
			got := UniqueCharacter(tc.s)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestUniqueCharacterImproved(t *testing.T){
	test := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Test 01",
			s:    "statistics",
			want: 3,
		},
		{
			name: "Test 02",
			s:    "haeck",
			want: 1,
		},
	}

	for _,tc := range test {
		t.Run(tc.name, func (t *testing.T){
			got := UniqueCharacterImproved(tc.s)
			require.Equal(t, tc.want, got)
		})
	}
}