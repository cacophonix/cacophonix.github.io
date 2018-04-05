package main

import (
	"fmt"
)

/*
 */
type Node struct {
	next *Node
	prev *Node
	key  string
}

type LRU struct {
	cache           map[string]string
	linkedListCache map[string]*Node
	maxSize         int
	first           *Node
	last            *Node
}

func (lru *LRU) init(max_size int) {
	lru.cache = make(map[string]string)
	lru.linkedListCache = make(map[string]*Node)
	lru.maxSize = max_size
	lru.first = nil
	lru.last = nil
}

func (lru *LRU) deleteNode(node *Node) {
	var prev, next *Node
	prev = node.prev
	next = node.next
	if prev != nil {
		prev.next = next
	} else {
		lru.first = next
	}
	if next != nil {
		next.prev = prev
	} else {
		lru.last = prev
	}
}

func (lru *LRU) deleteFirst() string {
	key := lru.first.key
	lru.first = lru.first.next
	if lru.first != nil {
		lru.first.prev = nil
	}
	return key
}

func (lru *LRU) appendNode(node *Node) {
	if lru.last != nil {
		lru.last.next = node
	}
	node.prev = lru.last
	node.next = nil
	lru.last = node

	if lru.first == nil {
		lru.first = node
	}
	if lru.last == nil {
		lru.last = node
	}
}

func (lru *LRU) set(key, val string) {
	lru.cache[key] = val
	node, ok := lru.linkedListCache[key]
	if ok {
		lru.deleteNode(node)
		lru.appendNode(node)
	} else {
		node := &Node{nil, lru.last, key}
		lru.linkedListCache[key] = node
		lru.appendNode(node)
	}
	n := len(lru.linkedListCache)
	if n > lru.maxSize {
		key = lru.deleteFirst()
		delete(lru.cache, key)
		delete(lru.linkedListCache, key)
	}
	//printkeys(lru.first)
}

func printkeys(node *Node) {
	for node != nil {
		fmt.Print(node.key + " ")
		node = node.next
	}
	fmt.Println()
}

func (lru *LRU) get(key string) (string, bool) {
	val, ok := lru.cache[key]
	if !ok {
		return "key not found: " + key, false
	}
	lru.set(key, val)
	return val, true
}

func main() {
	lru := &LRU{}
	lru.init(3)
	lru.set("1", "1")
	//state in linkedlist -> 1
	lru.set("2", "2")
	//state in linkedlist -> 1,2
	lru.set("3", "3")
	//state in linkedlist -> 1,2,3
	fmt.Println(lru.get("1")) //success
	//state in linkedlist -> 2,3,1
	lru.set("4", "4")
	//state in linkedlist -> 3,1,4
	fmt.Println(lru.get("4")) //success
	//state in linkedlist -> 3,1,4
	fmt.Println(lru.get("3")) //success
	//state in linkedlist -> 1,4,3
	fmt.Println(lru.get("2")) //fail
	lru.set("2", "2")
	//state in linkedlist -> 4,3,2
	fmt.Println(lru.get("2")) //success
	//state in linkedlist -> 4,3,2
	fmt.Println(lru.get("3")) // success
	//state in linkedlist -> 4,2,3
	fmt.Println(lru.get("1")) // fail
	//state in linkedlist -> 4,2,3
}
