package sorts

import (
	"fmt"
	"sync"
	"time"
)

func StartCountingsort(waitgroup *sync.WaitGroup, a []int, order *[]Sort) {
	start := time.Now()
	countingSort(a)
	elapsed := time.Since(start)
	fmt.Println("Counting Sort took: ", elapsed, "\n--- Sorted by Counting Sort ---\n", a, "\n")
	*order = append(*order, Sort{"Counting Sort", elapsed})
	waitgroup.Done()
}

func makeRange(min, max int) []int {
	arr := make([]int, max-min+1)
	for i := range arr {
		arr[i] = 0
	}
	return arr
}

func countingSort(a []int) []int {
	counter := makeRange(0, 9999)

	for _, e := range a {
		counter[e] += 1
	}

	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}

	res := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		e := a[i]           // elem to add
		t := counter[e] - 1 // target pos

		res[t] = e
		counter[e] = counter[e] - 1
	}

	return res
}
