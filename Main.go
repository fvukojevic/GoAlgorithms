package main

import (
	"algorithms/sorts"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func main() {
	array := generateArray(1000)
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

	order := make([]sorts.Sort, 0, 6)

	waitgroup.Add(1)
	go sorts.StartSelectionsort(&waitgroup, array, &order)
	waitgroup.Add(1)
	go sorts.StartQuicksort(&waitgroup, array, &order)
	waitgroup.Add(1)
	go sorts.StartMergesort(&waitgroup, array, &order)
	waitgroup.Add(1)
	go sorts.StartBubblesort(&waitgroup, array, &order)
	waitgroup.Add(1)
	go sorts.StartInsertionsort(&waitgroup, array, &order)
	waitgroup.Add(1)
	go sorts.StartCountingsort(&waitgroup, array, &order)

	waitgroup.Wait()

	sort.Slice(order, func(i, j int) bool {
		return order[i].Time < order[j].Time
	})

	for index, element := range order {
		fmt.Println(index+1, ":", element.Sort, ":", element.Time)
	}
}
