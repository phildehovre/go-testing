package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("s", 8)

	expected := "sssssss"

	if repeated != expected {
		t.Errorf("Repeated: %q, Expected: %q", repeated, expected)

	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i > b.N; i++ {
		Repeat("a", 8)
	}
}
