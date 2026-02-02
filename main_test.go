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

/* ===== PART 3 TESTS ===== */

func TestApply(t *testing.T) {
	nums := []int{1, 2, 3}
	result := Apply(nums, func(x int) int { return x * 2 })
	if result[0] != 2 || result[2] != 6 {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	evens := Filter(nums, func(x int) bool { return x%2 == 0 })
	if len(evens) != 2 {
		t.Fail()
	}
}

func TestReduce(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	sum := Reduce(nums, 0, func(a, b int) int { return a + b })
	if sum != 10 {
		t.Fail()
	}
}

func TestCompose(t *testing.T) {
	addTwo := func(x int) int { return x + 2 }
	double := func(x int) int { return x * 2 }
	f := Compose(addTwo, double)
	if f(5) != 12 {
		t.Fail()
	}
}

/* ===== PART 5 TESTS ===== */

func TestSwapValues(t *testing.T) {
	a, b := SwapValues(5, 10)
	if a != 10 || b != 5 {
		t.Fail()
	}
}

func TestSwapPointers(t *testing.T) {
	a, b := 5, 10
	SwapPointers(&a, &b)
	if a != 10 || b != 5 {
		t.Fail()
	}
}
