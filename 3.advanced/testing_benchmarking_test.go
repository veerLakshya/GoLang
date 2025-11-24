package advanced

import (
	"math/rand"
	"testing"
)

func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := GenerateRandomSlice(size)
	if len(slice) != size {
		t.Errorf("GenerateRandomSlice(%d) = %d, expected %d", size, len(slice), size)
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for range b.N {
		GenerateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GenerateRandomSlice(1000)
	b.ResetTimer()
	for range b.N {
		SumSlice(slice)
	}
}

// func Add(x, y int) int {
// 	return x + y
// }

// func BenchmarkAdd(b *testing.B) {
// 	for range b.N {
// 		Add(2, 3)
// 	}
// }
// func BenchmarkMediumInput(b *testing.B) {
// 	for range b.N {
// 		Add(200, 300)
// 	}
// }
// func BenchmarkLargeInput(b *testing.B) {
// 	for range b.N {
// 		Add(20000, 30000)
// 	}
// }

// Subtest Pattern Tests
/*
func TestAddSubtests(t *testing.T) {
	tests := []struct{ a, b, expected int }{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d, %d),", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("Add(%d,%d) = %d, expected %d", test.a, test.b, result, test.expected)
				return
			}
		})

	}
}
*/

// Table Pattern Tests
/*
func TestAddTableDriven(t *testing.T) {
	tests := []struct{ a, b, expected int }{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}
	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d,%d) = %d, expected %d", test.a, test.b, result, test.expected)
			return
		}

	}
}
*/

// func TestAdd(t *testing.T) {
// 	result := Add(2, 3)
// 	expected := 5
// 	if result != expected {
// 		t.Errorf("Add(2,3) = %d, expected %d", result, expected)
// 		return
// 	}
// }

/*
- test file should end with _test.go
- test function should start with Test
- benchmark function should start with Benchmark
- go test -bench=. to run benchmarks
- -benchmem flag to see memory allocations
- go test -bench=. -memprofile mem.pprof testing_benchmarking_test.go  for memory profiling
- go tool pprof mem.pprof to analyze memory profile


# Why is Testing Important?
	- Reliability
	- Maintainability
	- Documentation

# Profiling
	- Profiling provides detailed insights into the performace of your application, including CPU usage, memory allocation, and goroutine activity.
	- use `pprof` to collect and analyze CPU profile data

# best practices
	- write comprehensive tests
	- maintain test coverage
	- use benchmarks efficiently
	- profile regularly

- testing for quality assurance
- benchmarking for performance optimization
- profiling for performance analysis
*/
