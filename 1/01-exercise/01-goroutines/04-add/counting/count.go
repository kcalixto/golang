package counting

import (
	"math/rand"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	var sum int64
	// Utilize all cores on machine
	max := runtime.NumCPU() / 4
	parts := len(numbers) / max

	// Divide the input into parts
	incr := func(arr []int, c *chan int64) {
		var s int64
		for _, n := range arr {
			s += int64(n)
		}
		*c <- s
	}

	c := make(chan int64)
	for i := 0; i < max; i++ {
		pt := numbers[parts * (i) : parts * (i + 1)]
		go incr(pt, &c)
		sum += <-c

	}
	// Run computation for each part in seperate goroutine.

	// Add part sum to cummulative sum
	return sum
}
