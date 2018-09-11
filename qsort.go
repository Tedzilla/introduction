package main

import (
	"fmt"
	"math/rand"
)

func main() {
	myArray := []int{1, 12, 4, 5, 11, 7, 9}
	fmt.Println("my array is === ", qsort(myArray))
}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
