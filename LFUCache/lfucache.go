package main

/* Full implemenation of a lFU cache.
based on this blog post :
https://medium.com/@epicshane/a-python-implementation-of-lfu-least-frequently-used-cache-with-o-1-time-complexity-e16b34a3c49b
*/
import (
	"container/list"
	"errors"
	"sort"
)

type Node struct {
	value       interface{}
	frequency   int
	listElement *list.Element
}

type FreqMap map[int]*list.List

type LfuCache struct {
	Capacity     int
	UsedCapacity int
	KeyMap       map[int]Node
	FreqMap      FreqMap
}

func InitLfuCache(capacity int) *LfuCache {
	return &LfuCache{
		Capacity:     capacity,
		UsedCapacity: 0,
		KeyMap:       make(map[int]Node),
		FreqMap:      make(FreqMap),
	}
}

//RemoveFromMap : Remove the provided element from the frequency map's list
//and returns the value of the removed element. The value is acutally the key in the cache map
//The function still works if the value that you provide does not exist.
//it throws an exception if the key you're deleting from does not exist.
//it throws an exception if the el passed is nil.
func (f FreqMap) RemoveFromMap(freqKey int, el *list.Element) (interface{}, error) {
	if el == nil {
		return nil, errors.New("Element to remove cannot be nil")
	}

	if _, ok := f[freqKey]; !ok {
		return nil, errors.New("Key does not exist to remove element from")
	}

	freqList := f[freqKey]
	value := freqList.Remove(el)
	if value != nil {
		return value, nil
	}
	return nil, errors.New("The provided element could not be found")
}

// RemoveOldest : Removes the Oldest Element from a provided frequency.
//Retuns a value which is the key in actual cache map.
func (f FreqMap) RemoveOldest(freqKey int) (interface{}, error) {
	//check if the frequency exist. if not return nil, error.
	if _, ok := f[freqKey]; !ok {
		return nil, errors.New("Source Frequency Does not exist")
	}

	//check if the length of the list is not 0
	if f[freqKey].Len() == 0 {
		return nil, errors.New("Source Frequency has nothing to remove")
	}

	el := f[freqKey].Back()
	value := f[freqKey].Remove(el)
	if value != nil {
		return value, nil
	}
	return nil, errors.New("The provided element could not be found")
}

//AddToMap : Add an entry for a specific frequency to the front of the map.
//The logic to check whether this value was already on this frequency should
//not be present in this function.
//The value being passed here is actually the key in the cache map
func (f FreqMap) AddToMap(freqKey int, value interface{}) *list.Element {
	if _, ok := f[freqKey]; !ok {
		f[freqKey] = list.New()
	}
	el := f[freqKey].PushFront(value)
	if el != nil {
		return el
	}
	return nil
}

//SwapFrequency :  provide a source frequency, a target frequency and the el to move.
//This removes the el from the source frequency,
// adds it to the front of the target frequency.
//Returns nil if the swap is successful or corresponding errors.
func (f FreqMap) SwapFrequency(source int, target int, el *list.Element) error {
	//return if the element to delete is nil
	if el == nil {
		return errors.New("Element to Swap cannot be empty")
	}
	//check if the source freq exists and has any elements in the list.
	//if not return error
	if _, ok := f[source]; !ok {
		return errors.New("Source Frequency does not exist")
	}
	//check if it has any elements in the list.
	if f[source].Len() == 0 {
		return errors.New("The list at the source appears empty")
	}

	//if the source checks have passed, go ahead and delete the element
	//from the source list
	value, err := f.RemoveFromMap(source, el)
	if err != nil {
		return err
	}

	el = f.AddToMap(target, value)
	if el == nil {
		return errors.New("Error in adding value to target frequency")
	}
	return nil
}

//FindLeastFreq : Find the least frequency currently available in the frequency map with
// an active length greater than 0.
func (f FreqMap) FindLeastFreq() int {
	keys := []int{}
	for k := range f {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		if f[key].Len() > 0 {
			return key
		}
	}
	return -1
}

func (lfu *LfuCache) Get(key int) (interface{}, error) {
	//check if the entry exists in the Keymap.
	var node Node
	var ok bool
	if node, ok = lfu.KeyMap[key]; !ok {
		return nil, errors.New("The following key does not exist in the Cache")
	}

	//if the value does indeed exist in the keymap.
	//find the frequency and swap the node from one part to the other.
	err := lfu.FreqMap.SwapFrequency(node.frequency, node.frequency+1, node.listElement)
	if err != nil {
		return nil, err
	}
	return node.value, nil
}

func (lfu *LfuCache) Put(key int, value interface{}) error {
	var node Node
	var ok bool
	//check if the key exists.
	//if it does, then get the most recent entries.Capacity contraints dont apply
	if _, ok = lfu.KeyMap[key]; ok {
		//update the value of the node's frequency.
		node.value = value
		node.frequency++
		node.listElement.Value = key
		//re-assign the node with updated values to the key.
		lfu.KeyMap[key] = node
		//now move this from its older frequency queue to a new queue.
		err := lfu.FreqMap.SwapFrequency(node.frequency-1, node.frequency, node.listElement)
		if err != nil {
			return err
		}
		return nil
	}

	//the key does not exist in the map.
	//check if the UsedCapcity is less thatn the capcaity.
	if lfu.UsedCapacity < lfu.Capacity {
		//now add the list to the FreqMap.
		//set the value of the list element to be the key of the cache map
		el := lfu.FreqMap.AddToMap(1, key)
		if el == nil {
			return errors.New("Error Adding element to frequency map. Aborting")
		}
		lfu.KeyMap[key] = Node{
			value:       value,
			frequency:   1,
			listElement: el,
		}
		lfu.UsedCapacity++
		return nil
	}

	//the key does not exist in the map, but we need to remove
	//the least frequency used element from the map to be able to add the
	//element here.
	if lfu.UsedCapacity >= lfu.Capacity {
		//find the smallest key from the FreqMap that has zero element.
		val := lfu.FreqMap.FindLeastFreq()
		if val != -1 {
			value, err := lfu.FreqMap.RemoveOldest(val)
			//get the key for the map from the oldest element.
			if err == nil {
				delete(lfu.KeyMap, value.(int))
			}
			le := lfu.FreqMap.AddToMap(1, key)
			if le == nil {
				return errors.New("Error Adding element to frequency map. Aborting")
			}
			lfu.KeyMap[key] = Node{
				value:       value,
				frequency:   1,
				listElement: le,
			}

		}
	}

	return nil
}
