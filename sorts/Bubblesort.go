package sorts

import (
	"fmt"
	"sync"
	"time"
)

func StartBubblesort(waitgroup *sync.WaitGroup, a []int, order *[]Sort) {
	start := time.Now()
	bubblesort(a)
	elapsed := time.Since(start)
	fmt.Println("Bubble Sort took: ", elapsed, "\n--- Sorted by Bubble Sort ---\n", a, "\n")
	*order = append(*order, Sort{"Bubble Sort", elapsed})
	waitgroup.Done()
}

func bubblesort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}
