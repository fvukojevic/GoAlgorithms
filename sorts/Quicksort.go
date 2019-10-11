package sorts

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func StartQuicksort(waitgroup *sync.WaitGroup, a []int, order *[]string) {
	start := time.Now()
	quicksort(a)
	elapsed := time.Since(start)
	fmt.Println("Quick Sort took: ", elapsed, "\n--- Sorted by Quick Sort ---\n", a, "\n")
	*order = append(*order, "Quick Sort")
	waitgroup.Done()
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
