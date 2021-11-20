package algorithms

import "testing"

func TestValid(t *testing.T) {

	tests := []struct {
		name string
		id   string
		want bool
	}{
		{
			name: "Test 01",
			id:   "05 9",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.id); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
