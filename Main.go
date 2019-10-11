package main

import (
	"algorithms/sorts"
	"math/rand"
	"sync"
	"time"
)

func main() {
	array := generateArray(10000)
	startSorting(array)
}

// Generates a slice of size, size filled with random numbers
func generateArray(size int) []int {

	array := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(9999)
	}
	return array
}

func startSorting(array []int) {
	var waitgroup sync.WaitGroup

	waitgroup.Add(1)
	go sorts.StartSelectionsort(&waitgroup, array)
	waitgroup.Add(1)
	go sorts.StartQuicksort(&waitgroup, array)
	waitgroup.Add(1)
	go sorts.StartMergesor(&waitgroup, array)
	waitgroup.Add(1)
	go sorts.StartBubblesort(&waitgroup, array)
	waitgroup.Add(1)
	go sorts.StartInsertionsort(&waitgroup, array)
	waitgroup.Add(1)
	go sorts.StartCountingsort(&waitgroup, array)

	waitgroup.Wait()
}
