package bloom

import (
	"math/rand"
	"testing"
)

func TestConstantTimeBasic(t *testing.T) {
	f := New(1000, 4)
	n1 := []byte("Bess")
	n2 := []byte("Jane")
	n3 := []byte("Emma")
	f.Add(n1)
	n3a := f.ConstantTimeTest(n3)
	f.Add(n3)
	n1b := f.ConstantTimeTest(n1)
	n2b := f.ConstantTimeTest(n2)
	n3b := f.ConstantTimeTest(n3)
	if !n1b {
		t.Errorf("%v should be in.", n1)
	}
	if n2b {
		t.Errorf("%v should not be in.", n2)
	}
	if n3a {
		t.Errorf("%v should not be in the first time we look.", n3)
	}
	if !n3b {
		t.Errorf("%v should be in the second time we look.", n3)
	}
}

func BenchmarkTest(b *testing.B) {
	f := NewWithEstimates(1000, 0.0001)
	key := make([]byte, 100)
	rand.Read(key)
	f.Add(key)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Test(key)
	}
}

func BenchmarkConstantTimeTest(b *testing.B) {
	f := NewWithEstimates(1000, 0.0001)
	key := make([]byte, 100)
	rand.Read(key)
	f.Add(key)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.ConstantTimeTest(key)
	}
}
