package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubblesort(datalist []int) []int {
	l := len(datalist) - 1
	if l <= 0 {
		return datalist
	}
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < l; i++ {
			if datalist[i] > datalist[i+1] {
				datalist[i], datalist[i+1] = datalist[i+1], datalist[i]
				swapped = true
			}
		}
		l--
	}
	return datalist
}

func quickSort(a []int) []int {
	// It uses the idea of divide and conquer
	// it finds the pivot such that, elements on its left are smaller then pivot and elements in its right are greater then the pivot
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := right / 2
	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

func generateSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	return rand.Perm(size)
}

func main() {
	items := generateSlice(13)
	fmt.Println("\n--- Unsorted --- \n\n", items)
	// fmt.Println("\n--- bubblesort items --- \n\n", bubblesort(items))
	fmt.Println("\n --- quicksorted items -- \n\n", quickSort(items))
}
