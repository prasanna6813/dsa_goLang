package main

import "fmt"

const ArraySize = 7

// HashTable structure
// HashTable will hold an array..
type HashTable struct{
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the
type bucket struct{
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key string
	next *bucketNode
}

// Insert
// search
// Delete

// Insert
// search
// Delete

// hashFunc

// Init - initallizes the hash table

func main(){
	testHashTable := &HashTable{}
	fmt.Println(testHashTable)
}