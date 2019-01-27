package main

import (
	"container/list"
	"fmt"
)

//Node  : Is the struct that Gets stored as a value pair in the key map
type Node struct {
	Key   int
	Value interface{}
	el    *list.Element
}

type LRUCache struct {
	capacity int
	cachemap map[int]Node
	prique   *list.List
}

func InitCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cachemap: make(map[int]Node),
		prique:   list.New(),
	}
}

// Put : Inserts a value into the cache.
func (lru *LRUCache) Put(key int, value interface{}) {
	//check if that key already exists.
	if _, ok := lru.cachemap[key]; ok {
		//update the value
		node := lru.cachemap[key]
		//set update value for this node.
		node.Value = value
		//move the element to the front.
		lru.prique.MoveToFront(node.el)
		return
	}
	//check the capacity of the cache.
	if lru.prique.Len() < lru.capacity {
		//since the capcaity isn't full, we can add more elements here.
		//create a list Element
		el := lru.prique.PushFront(key)
		//add to map.
		lru.cachemap[key] = Node{Key: key, Value: value, el: el}
	} else {
		//evict the key that is at the back of the queue.
		el := lru.prique.Back()
		lru.prique.Remove(el)
		//remove from the cache map too.
		delete(lru.cachemap, el.Value.(int))
	}
	return
}

func (lru *LRUCache) Get(key int) interface{} {
	//check if the value exists in the map.
	if _, ok := lru.cachemap[key]; ok {
		node := lru.cachemap[key]
		lru.prique.MoveToFront(node.el)
		return node.Value
	}
	return -1
}

func main() {
	cache := InitCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}
