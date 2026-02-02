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



