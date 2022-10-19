package p1_test

import (
	"testing"

	"github.com/rommms07/projecteuler.net/p1"
)

func TestProblem1Solution1(t *testing.T) {
	expect_out := map[int64]int64{
		10:   23,
		1000: 233168,
	}

	for in, out := range expect_out {
		sout := p1.Solution1(in)
		if sout != out {
			t.Errorf("expected: f(%d) = %d [output: %d]", in, out, sout)
		}
	}
}

func BenchmarkProblem1Solution1(b *testing.B) {
	for k := 0; k < b.N; k++ {
		p1.Solution1(int64(1000))
	}
}

func BenchmarkProblem1Solution2(b *testing.B) {
	for k := 0; k < b.N; k++ {
		p1.Solution2(int64(1000))
	}
}
