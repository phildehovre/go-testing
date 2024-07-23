package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	assertAddition(t, sum, expected)
}

func assertAddition(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("Sum %d, expected %d", sum, expected)
	}
}

// These types of Example function will compare the stdoutput to the "Output" comment.
func ExampleAddition() {
	fmt.Println(Add(3, 4))
	// Output: 7
}
