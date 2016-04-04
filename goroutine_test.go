package goroutine

import "testing"

func TestGet(t *testing.T) {
	t.Logf("Goroutine ID is: %d\n", GoroutineId())
}

func BenchmarkGet(b *testing.B) {
	_ = GoroutineId()
}
