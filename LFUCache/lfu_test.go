package main

import (
	"testing"
)

func TestAddToMap(t *testing.T) {
	fremap := make(FreqMap)
	//check what happens when you try to add a freq and a value
	if fremap.AddToMap(1, 1) == nil {
		t.Errorf("TestAddToMap : Expected *list.Element{1} on freq 1, returned nil")
	}
	//now check what happens if you add a new value to the same freq
	if fremap.AddToMap(1, 2) == nil {
		t.Errorf("TestAddtoMap : Expected *list.Element{2} on freq 1, returned nil")
	}
	//now check what happend if you add an all new frequency
	if fremap.AddToMap(2, 1) == nil {
		t.Errorf("TestAddtoMap : Expected *list.Element{2} on freq 2, returned nil")
	}
}

func TestRemoveFromMap(t *testing.T) {
	fremap := make(FreqMap)
	el := fremap.AddToMap(1, 1)
	//the element cannot be nil before you call RemoveFromMap
	if el != nil {
		value, _ := fremap.RemoveFromMap(1, el)
		if value != 1 {
			t.Errorf("TestRemoveFromMap : Expected 1 as the return value. 1 not returned")
		}

		//now test to see what happens if you try to remove value from a freq that does not exist
		_, err := fremap.RemoveFromMap(2, el)
		if err == nil {
			t.Errorf("TestRemoveFromMap : Expected error while trying to delete value from freq that doesn't exist. Returned nil")
		}
	} else {
		t.Errorf("TestRemoveFromMap : Error creating mock value to test")
	}
}

func TestSwapFrequency(t *testing.T) {
	fremap := make(FreqMap)
	el := fremap.AddToMap(1, 1)
	err := fremap.SwapFrequency(1, 2, el)
	if err != nil {
		t.Errorf("TestSwapFrequency : Expected nil. Returned Error")
	}

	//now check what happens if you try to make the same swap again
	err = fremap.SwapFrequency(1, 2, el)
	if err == nil {
		t.Errorf("TestSwapFrequency : Expected error frequency 1 list empty. Returned nil")
	}

	//check if the source frequency does not exist
	err = fremap.SwapFrequency(5, 2, el)
	if err == nil {
		t.Errorf("TestSwapFrequency : Expected error source frequency does not exist. Returned nil")
	}

	//check error with nil element.
	err = fremap.SwapFrequency(1, 2, nil)
	if err == nil {
		t.Errorf("TestSwapFrequency : Expected error nil element. No error returned")
	}
}

func TestRemoveOldest(t *testing.T) {
	fremap := make(FreqMap)
	fremap.AddToMap(1, 1)
	fremap.AddToMap(1, 2)
	fremap.AddToMap(1, 3)

	//removing oldest from freq 1
	value, err := fremap.RemoveOldest(1)
	if err != nil || value != 1 {
		t.Errorf("TestRemoveOldest: Expected value 1 and err nil. Returned error")
	}
	//remove the next oldest
	value, err = fremap.RemoveOldest(1)
	if err != nil || value != 2 {
		t.Errorf("TestRemoveOldest: Expected value 2 and err nil. Returned error")
	}

	//remove the next oldest
	value, err = fremap.RemoveOldest(1)
	if err != nil || value != 3 {
		t.Errorf("TestRemoveOldest: Expected value 3 and err nil. Returned error")
	}

	//check that it cannot remove elements from an empty source
	_, err = fremap.RemoveOldest(1)
	if err == nil {
		t.Errorf("TestRemoveOldest: Expected err list empty, returned nil")
	}

	//check that it cannot remvoe elements from a frequency that does not exist
	_, err = fremap.RemoveOldest(2)
	if err == nil {
		t.Errorf("TestRemoveOldest: Expected err freq not found, returned nil")
	}
}

func TestFindLeastFreq(t *testing.T) {
	fremap := make(FreqMap)
	fremap.AddToMap(1, 1)
	fremap.AddToMap(2, 2)
	value := fremap.FindLeastFreq()
	if value != 1 {
		t.Errorf("TestLeastFreq: Expected 1")
	}
	fremap.RemoveOldest(1)
	value = fremap.FindLeastFreq()
	if value != 2 {
		t.Errorf("TestLeastFreq: Expected 2, returned")
	}
	fremap.AddToMap(1, 1)
	value = fremap.FindLeastFreq()
	if value != 1 {
		t.Errorf("TestLeastFreq: Expected 1")
	}
}

func TestGetPut(t *testing.T) {
	lfu := InitLfuCache(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	val, err := lfu.Get(1) // returns 1
	if val != 1 || err != nil {
		t.Errorf("TestGetPut: Expected 1")
	}
	lfu.Put(3, 3)         // evicts key 2
	val, err = lfu.Get(2) // returns nil (not found)
	if val != nil {
		t.Errorf("TestGetPut: Expected nil")
	}
	lfu.Get(3)    // returns 3.
	lfu.Put(4, 4) // evicts key 1.
	lfu.Get(1)    // returns -1 (not found)
	if val != nil {
		t.Errorf("TestGetPut: Expected nil")
	}
	lfu.Get(3) // returns 3
	lfu.Get(4) // returns 4
}
