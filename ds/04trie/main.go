package main

import "fmt"

// Node represents each node in the trie

type Node struct {
	children [26]*Node
	isEnd    bool
}

// Trie represents a trie a trie and has a pointer the node
type Trie struct{ root *Node }

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie

func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and RETURN true is that word is included in the

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true{
		return true 
	}
	return false
}

func main() {
	myTrie := InitTrie()

	toAdd := []string{
		"aragorn","aragon","eragon","oreio", "oregano",
	}

	for _,v:=range toAdd{
		myTrie.Insert(v)
	}


	myTrie.Insert("aragorn")
	fmt.Println(myTrie.Search("aron"))
	// fmt.Println(testTrie.root)

}
