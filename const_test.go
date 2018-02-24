package bloom

import "testing"

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
