package math

import "testing"

func TestSubstract(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 2, 3},
		{3, 3, 0},
		{9, 1, 8},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Substract(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		} else {
			t.Logf("Substract(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}