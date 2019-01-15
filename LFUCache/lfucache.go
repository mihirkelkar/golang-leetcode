package main

import (
	"container/list"
	"time"
)

type LfuCache struct {
	capacity int
	cache    map[int]Node
	//We are going to use the container/list data structure to build this.
	//We will only pass keys to the eviction list.
	//our least freequently used key will be at the back of the list.
	//new elements will therefore be added to the back of the list.
	//every new access moves these to the front.
	evictionList *list.List
}

type Node struct {
	Value       int
	Accesscount int
	Timestamp   time.Time
}

func NewLfuCache(capacity int) LfuCache {
	l := LfuCache{
		capacity:     capacity,
		cache:        make(map[int]Node),
		evictionList: list.New(),
	}
	return l
}

func (l *LfuCache) Get(key int) int {
	var node Node
	var ok bool
	if node, ok = l.cache[key]; !ok {
		return -1
	}
	//set access count higher
	node.Accesscount++
	node.Timestamp = time.Now()
	//move the key of this node to the front of the list.
	return node.Value
}

func (l *LfuCache) Put(key, value int) {
	//check if that key already exists. if it does, just update it
	//and move to the front of the queue.
	// if capacity is not full and key does not exist.
	//add new key to map. Add new key to backoflist.
	//if capacity is full and key does not exist.
	//find the key at the back of the evictList
	//Evict that key from the list and from the map.
	//Add new key at backoflist.
}

//func (l *LfuCache) Put(key, value) {
//create a temp node that
//}
