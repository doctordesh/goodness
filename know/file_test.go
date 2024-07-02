package know

import "testing"

func BenchmarkDo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Do("lorem")
	}
}
