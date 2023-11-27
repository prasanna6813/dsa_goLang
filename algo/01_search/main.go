package main

import (
	"fmt"
)

func linearSearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func binarySearch(datalist []int, key int) bool {
	low,median, high := 0,0, len(datalist)-1

	for low <= high {
		median = (low + high) / 2
		// fmt.Println(datalist[median]);
		if datalist[median] < key {
			low = median + 1
		} else if datalist[median] > key {
			high = median - 1
		} else {
			return true
		}
	}
	return false
}

func interpolationSearch(datalist []int, key int) bool {
	// data should be uniformly distributed..
	low, mid, high, min, max := 0, 0, len(datalist)-1, datalist[0], datalist[len(datalist)-1];
	if key < min || key > max {
		return false
	}
	for low <= high {
		mid = low + (key-min)*(high-low) / (max-min);
		// mid = (low + high) / 2;
		// fmt.Println(datalist[mid]);
		if datalist[mid] < key {
			low = mid + 1
		} else if datalist[mid] > key {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	items := []int{1, 2, 4, 6, 9, 20, 22, 26, 31, 33, 34, 38, 40, 43, 45, 63, 64, 65, 66,67,69, 70, 100,}

	fmt.Println(binarySearch(items, 31))

	fmt.Println(linearSearch(items, 334))

	fmt.Println(interpolationSearch(items, 22))
}
