package main

import (
	"errors"
	"fmt"
	"os"
)

/* =========================
   PART 1: Math Operations
   ========================= */

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

/* =========================
   PART 2: Closures
   ========================= */

func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (func(int), func(int), func() int) {
	value := initial

	add := func(x int) {
		value += x
	}
	sub := func(x int) {
		value -= x
	}
	get := func() int {
		return value
	}

	return add, sub, get
}

/* =========================
   PART 3: Higher-Order Functions
   ========================= */

func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = operation(v)
	}
	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce(nums []int, initial int, operation func(int, int) int) int {
	acc := initial
	for _, v := range nums {
		acc = operation(acc, v)
	}
	return acc
}

func Compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

/* =========================
   PART 4: Process Explorer
   ========================= */

func ExploreProcess() {
	fmt.Println("=== Process Information ===")
	fmt.Println("Process ID:", os.Getpid())
	fmt.Println("Parent Process ID:", os.Getppid())

	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Address of slice header: %p\n", &data)
	fmt.Printf("Address of first element: %p\n", &data[0])

	fmt.Println("Other processes cannot access these memory addresses due to process isolation.")
}

/* =========================
   PART 5: Pointers & Escape Analysis
   ========================= */

func DoubleValue(x int) {
	x *= 2
	// Does NOT modify original value (pass-by-value)
}

func DoublePointer(x *int) {
	*x *= 2
	// Modifies original value (pointer)
}

func CreateOnStack() int {
	x := 10
	return x // stays on stack
}

func CreateOnHeap() *int {
	x := 10
	return &x // escapes to heap
}

func SwapValues(a, b int) (int, int) {
	return b, a
}

func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}
