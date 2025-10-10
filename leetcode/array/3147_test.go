package larray

import "testing"

func TestMaximumEnergy(t *testing.T) {
	tests := []struct {
		name     string
		energy   []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			energy:   []int{5, 2, -10, -5, 1},
			k:        3,
			expected: 3,
		},
		{
			name:     "Example 2",
			energy:   []int{-2, -3, -1},
			k:        2,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumEnergy(tt.energy, tt.k)
			if result != tt.expected {
				t.Errorf("maximumEnergy(%v, %d) = %d, expected %d",
					tt.energy, tt.k, result, tt.expected)
			}
		})
	}
}
