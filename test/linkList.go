package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(t *Node, v int) {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return
	}

	for ; t.Next != nil; t = t.Next {
		//fmt.Println("")
	}

	t.Next = &Node{v, nil}

}

func contain(t *Node, v int) bool {
	if t == nil {
		return false
	}

	for ; t != nil; t = t.Next {
		if t.Value == v {
			return true
		}
	}

	return false
}

func removeNode(t *Node, v int) *Node {
	if t == nil {
		return t
	}

	//remove the first element
	root := t
	if t.Value == v {
		root = t.Next
		return root
	}

	for root.Next != nil {
		fmt.Println(root.Value)
		if root.Next.Value == v {
			root.Next = root.Next.Next
			return t
		}
		root = root.Next
	}

	return t

}

func main() {
	root = nil
	k := 3

	addNode(root, 3)
	addNode(root, 1)
	addNode(root, 2)
	addNode(root, 3)
	addNode(root, 4)
	addNode(root, 5)

	//fmt.Println(root)

	//l1 := root

	//fmt.Println(contain(root, 8))

	l1 := root

	for contain(root, k) {
		l1 = removeNode(root, k)
	}
	fmt.Println(l1)

	for ; root != nil; root = root.Next {
		fmt.Println(root.Value)
	}

	//	for ; l1 != nil; l1 = l1.Next {
	//		fmt.Println(l1.Value)
	//	}

	/*
		if l1.Value == k {
			root = l1.Next
		}

		fmt.Println(root)
		//fmt.Println(l1)
		// l1 := root
		//l1 = l1.Next

		//fmt.Println(l1.Value, "Next add")

		// test1 = root

			for test1 := root; test1 != nil; test1 = test1.Next {
				fmt.Println(test1.Value)
			}
	*/

}
