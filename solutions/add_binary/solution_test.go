package add_binary

import "testing"

func BenchmarkAddBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddBinary("10100001", "1111111111")
	}
}
