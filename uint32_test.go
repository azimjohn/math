package math

import "testing"

var maxUint32Tests = []struct {
	a, b uint32
	want uint32
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
}

func TestMaxUint32(t *testing.T) {
	for i, tt := range maxUint32Tests {
		got := MaxUint32(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxUint(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}