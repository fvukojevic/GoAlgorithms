package sorts

import (
	"fmt"
	"sync"
	"time"
)

func StartInsertionsort(waitgroup *sync.WaitGroup, a []int) {
	start := time.Now()
	insertionsort(a)
	elapsed := time.Since(start)
	fmt.Println("Insertion Sort took: ", elapsed, "\n--- Sorted by Insertion Sort ---\n", a, "\n")
	waitgroup.Done()
}

func insertionsort(a []int) []int {
	var n = len(a)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j = j - 1
		}
	}

	return a
}
