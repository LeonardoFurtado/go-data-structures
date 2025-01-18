package main

import (
	"fmt"
	"sync"
)

func sumSlice(slice []int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, num := range slice {
		sum += num
	}
	resultChan <- sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	resultChan := make(chan int)

	var wg sync.WaitGroup

	mid := len(numbers) / 2
	firstHalf := numbers[:mid]
	secondHalf := numbers[mid:]

	wg.Add(2)
	go sumSlice(firstHalf, resultChan, &wg)
	go sumSlice(secondHalf, resultChan, &wg)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	finalSum := 0
	for partialSum := range resultChan {
		finalSum += partialSum
	}

	fmt.Printf("Result: %d\n", finalSum)
}
