package allocsample

import "testing"

func Benchmark_uniqueDeclaration(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uniqueDeclaration()
	}
}
func Benchmark_declarationInsideLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		declarationInsideLoop()
	}
}
