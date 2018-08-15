package main

import (
	"fmt"
	"math"
)

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807
*/
type Numnode struct {
	value int
	next  *Numnode
}

func createNodeChain(num int, prev_node *Numnode) *Numnode {
	newNode := Numnode{num, nil}
	if prev_node != nil {
		prev_node.next = &newNode
	}
	return &newNode
}

func convertNumberToNodes(num int) *Numnode {
	/*
	  Takes an input number and converts it into a node chain as
	  defined in the example above
	*/
	nodesCounter := 0
	var firstNode *Numnode
	var prevNode *Numnode
	for num > 0 {
		if nodesCounter == 0 {
			firstNode = createNodeChain(num%10, nil)

		} else {
			if prevNode == nil {
				prevNode = firstNode
			}
			prevNode = createNodeChain(num%10, prevNode)
		}
		num = num / 10
		nodesCounter += 1
	}
	return firstNode
}

func createNodesToNumber(firstNode *Numnode) int {
	curNode := firstNode
	var indexcount int = 0
	var numsum int
	for curNode != nil {
		numsum += curNode.value * int(math.Pow10(indexcount))
		indexcount += 1
		curNode = curNode.next
	}
	return numsum
}

func main() {
	ip1 := convertNumberToNodes(342)
	ip2 := convertNumberToNodes(465)
	num := createNodesToNumber(ip1) + createNodesToNumber(ip2)
	op := convertNumberToNodes(num)
	for op != nil {
		fmt.Println(op.value)
		op = op.next
	}
}
