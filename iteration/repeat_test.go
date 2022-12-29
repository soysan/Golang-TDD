package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// When typing Prefix with ExampleXXX, it will show on godoc
func ExampleRepeat() {
	repeated := Repeat("a", 8)
	fmt.Println(repeated)
	// Output "aaaaaaaa"
}

// When typing Prefix with BenchmarkXXX and executing “go test -bench=$path", it can check benchmark.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
