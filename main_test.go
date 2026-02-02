package main

import "testing"

/* ===== PART 1 TESTS ===== */

func TestFactorial(t *testing.T) {
	tests := []struct {
		input   int
		want    int
		wantErr bool
	}{
		{0, 1, false},
		{5, 120, false},
		{-1, 0, true},
	}

	for _, tt := range tests {
		got, err := Factorial(tt.input)
		if (err != nil) != tt.wantErr {
			t.Fail()
		}
		if got != tt.want {
			t.Fail()
		}
	}
}

func TestIsPrime(t *testing.T) {
	if ok, _ := IsPrime(17); !ok {
		t.Fail()
	}
	if ok, _ := IsPrime(20); ok {
		t.Fail()
	}
}

func TestPower(t *testing.T) {
	if v, _ := Power(2, 8); v != 256 {
		t.Fail()
	}
}

/* ===== PART 2 TESTS ===== */

func TestMakeCounter(t *testing.T) {
	c := MakeCounter(0)
	if c() != 1 || c() != 2 {
		t.Fail()
	}
}

func TestMakeMultiplier(t *testing.T) {
	double := MakeMultiplier(2)
	if double(5) != 10 {
		t.Fail()
	}
}

func TestMakeAccumulator(t *testing.T) {
	add, sub, get := MakeAccumulator(100)
	add(50)
	sub(30)
	if get() != 120 {
		t.Fail()
	}
}
