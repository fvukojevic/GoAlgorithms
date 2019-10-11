package sorts

import (
	"fmt"
	"sync"
	"time"
)

func StartMergesort(waitgroup *sync.WaitGroup, a []int, order *[]string) {
	start := time.Now()
	mergeSort(a)
	elapsed := time.Since(start)
	fmt.Println("Merge Sort took: ", elapsed, "\n--- Sorted by Merge Sort ---\n", a, "\n")
	*order = append(*order, "Merge Sort")
	waitgroup.Done()
}

func mergeSort(a []int) []int {
	var num = len(a)

	if num == 1 {
		return a
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = a[i]
		} else {
			right[i-middle] = a[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
