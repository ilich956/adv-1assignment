package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	numbers := make([]float64, 10000)
	for i := 0; i < 1000; i++ {
		numbers[i] = float64(i)
	}

	start := time.Now()
	operation(numbers)
	elapsed := time.Since(start)
	fmt.Printf("Without goroutines: %s\n", elapsed)

	for _, numGoroutines := range []int{1, 10, 100, 1000} {
		start = time.Now()
		gooperation(numbers, numGoroutines)
		elapsed = time.Since(start)
		fmt.Printf("With %d goroutines: %s\n", numGoroutines, elapsed)
	}
}

func operation(numbers []float64) {
	for _, num := range numbers {
		for i := 0; i < 1000; i++ {
			_ = math.Sqrt(num)
		}
	}
}

func gooperation(numbers []float64, numGoroutines int) {
	var wg sync.WaitGroup
	chunkSize := len(numbers) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(numbers)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				for i := 0; i < 1000; i++ {
					_ = math.Sqrt(numbers[j])
				}
			}
		}(start, end)
	}

	wg.Wait()
}
