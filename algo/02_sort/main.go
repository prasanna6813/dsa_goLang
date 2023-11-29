package main

import (
	"fmt"
	"math/rand"
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

	// fmt.Println("left:", left, "right:", right, "pivot:", pivot)

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

func selectionsort(datalist []int) []int {
	l := len(datalist)
	for i := 0; i < l; i++ {
		var minIdex = i
		for j := i; j < l; j++ {
			if datalist[j] < datalist[minIdex] {
				minIdex = j
			}
		}
		datalist[i], datalist[minIdex] = datalist[minIdex], datalist[i]
	}
	return datalist
}

func insertionSort(datalist []int) []int {
	l := len(datalist)
	for i := 1; i < l; i++ {
		j := i
		for j > 0 {
			if datalist[j-1] > datalist[j] {
				datalist[j-1], datalist[j] = datalist[j], datalist[j-1]
			}
			j = j - 1
		}
	}
	return datalist
}

func shellSort(datalist []int) []int {
	// Shell sort is a highly efficient sorting algorithm that is based on the insertion sort algorithm. It works by repeatedly breaking the array down into smaller subarrays and sorting them using insertion sort. This process gradually produces a sorted array.
	gap := len(datalist) / 2

	for gap > 0 {
		for i := gap; i < len(datalist); i++ {
			temp := datalist[i]
			j := i

			for ; j >= gap && datalist[j-gap] > temp; j -= gap {
				fmt.Println("j:", j, "  gap:", gap, "   temp:", temp)
				datalist[j] = datalist[j-gap]
			}
			datalist[j] = temp
		}
		gap /= 2
	}
	return datalist
}

func main() {
	items := rand.Perm(16)
	fmt.Println("\n--- Unsorted --- \n\n", items)
	// fmt.Println("\n--- bubblesort items --- \n\n", bubblesort(items))
	fmt.Println("\n --- quicksorted items -- \n\n", quickSort(items))
	// fmt.Println("\n --- selectionsorted items -- \n\n", selectionsort(items))
	// fmt.Println("\n --- insertionSorted items -- \n\n", insertionSort(items))
	// fmt.Println("\n --- shellSorted items -- \n\n", shellSort(items))

}
