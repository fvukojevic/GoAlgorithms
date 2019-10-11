package sorts

import (
	"fmt"
	"sync"
	"time"
)

func StartSelectionsort(waitgroup *sync.WaitGroup, a []int) {
	start := time.Now()
	selectionsort(a)
	elapsed := time.Since(start)
	fmt.Println("Selection Sort took: ", elapsed, "\n--- Sorted by Selection Sort ---\n", a, "\n")
	waitgroup.Done()
}

func selectionsort(a []int) []int {
	var n = len(a)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}

	return a
}
