package memtable

import (
	"fmt"
	"math"
	"strings"
)

type Node struct {
	key   string
	data  string
	left  *Node
	right *Node
}

type Rotation string

const (
	LEFT_LEFT   string = "LL"
	LEFT_RIGHT  string = "LR"
	RIGHT_LEFT  string = "RL"
	RIGHT_RIGHT string = "RR"
)

// --- PUBLIC APIs ---
// Create an AVL tree
// A design choice here is to create a Node with temp placeholder and then return the address for
// that node -> always be our AVL tree initialized -> or just simply call the create with
// already decided key value pair -> which can be overwritten once we get actual data

// Insert in an AVL tree
func AvlInsert(root *Node, key string, value string) *Node {
	// Create am insertable node
	toInsert := createNode(key, value)

	// Call raw insert
	newRoot := avlInsertUtil(root, toInsert)

	return newRoot
}

// Search in an AVL tree
func AvlSearchKey(root *Node, key string) (bool, string) {
	if root == nil {
		return false, ""
	}

	if root.key == key {
		return true, root.data
	}

	cmpRes := strings.Compare(root.key, key)
	if cmpRes == -1 {
		return AvlSearchKey(root.right, key)
	} else {
		return AvlSearchKey(root.left, key)
	}
}

// -- PUBLIC API end --

// -- PRIVATE API --

// Custom sort to identify smaller larger key
func sortKey(nodeOne *Node, nodeTwo *Node) int {
	cmpRes := strings.Compare(nodeOne.key, nodeTwo.key)
	return cmpRes
}

// Create a node
func createNode(key string, value string) *Node {
	return &Node{
		key:   key,
		data:  value,
		left:  nil,
		right: nil,
	}
}

// Get Height of Node/Tree
func getAvlTreeHeight(root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight := getAvlTreeHeight(root.left)
	rightHeight := getAvlTreeHeight(root.right)

	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}

// Check imbalance
func checkImbalance(root *Node) bool {
	if root == nil {
		return false
	}

	leftHeight := getAvlTreeHeight(root.left)
	rightHeight := getAvlTreeHeight(root.right)

	if int(math.Abs(float64(leftHeight-rightHeight))) > 1 {
		return true
	}

	return false
}

// Identify the type of imbalance
func getTypeImbalance(root *Node) string {
	var rotation string = ""

	if root.left == nil {
		rotation = rotation + "R"
		root = root.right
	} else {
		rotation = rotation + "L"
		root = root.left
	}

	if root.left == nil {
		rotation = rotation + "R"
	} else {
		rotation = rotation + "L"
	}

	return rotation
}

// Free up a node

// Left - Left Rotation
func avlLeftLeftRotation(root *Node) *Node {
	/*
					root											mid
					/												/			\
				mid        ----->    leaf				root
				/
		leaf

	*/
	// store mid -> clear root of all child duties
	// make root right child of mid

	var mid *Node = root.left

	root.left = nil
	mid.right = root

	return mid
}

// Right - Right rotation
func avlRightRightRotation(root *Node) *Node {
	/*
			root												mid
		 			\											/			\
					mid        ----->   root	 leaf
						\
						leaf

	*/

	// store mid -> clear root of all child duties -> make root the left child of mid
	var mid *Node = root.right

	root.right = nil
	mid.left = root

	return mid
}

// Left - Right Rotation
func avlLeftRightRotation(root *Node) *Node {
	/*
			root											leaf
			/												/			\
		mid        ----->     mid				root
			\
			leaf

	*/
	// mid < leaf < root
	// store mid, leaf and root -> free mid and root of all child duties ->
	// make mid left and root right child of leaf -> return leaf

	var mid *Node = root.left
	var leaf *Node = mid.right

	root.left = nil
	mid.right = nil
	leaf.left = mid
	leaf.right = root

	return leaf
}

// Right - Left rotation
func avlRightLeftRotation(root *Node) *Node {
	/*
		root											  leaf
			\												/			\
			mid        ----->    root			mid
			/
		leaf
	*/
	// root < leaf < mid
	// store mid, leaf -> free mid and root of all child duties
	// make root the left node and mid the right node of leaf -> return leaf

	var mid *Node = root.right
	var leaf *Node = mid.left

	root.right = nil
	mid.left = nil
	leaf.left = root
	leaf.right = mid

	return leaf
}

// insert helper
func avlInsertUtil(root *Node, toInsert *Node) *Node {
	// accomodate the new node in empty space on tree
	if root == nil {
		return toInsert
	}

	// find the best spot to insert
	cmpRes := sortKey(root, toInsert)

	// if the same key -> cmpRes = 0 and we just update the value
	if cmpRes == 0 {
		root.data = toInsert.data
		// Free up toInsert
		return root
	}

	// cmpRes == -1 if root < toInsert & cmpRes == 1 if root > toInsert
	if cmpRes == -1 {
		rightRoot := avlInsertUtil(root.right, toInsert)
		root.right = rightRoot
		if checkImbalance(root.right) == true {
			typeImbalance := getTypeImbalance(root.right)
			fmt.Println(typeImbalance)
			switch typeImbalance {
			case LEFT_LEFT:
				root.right = avlLeftLeftRotation(root.right)
			case LEFT_RIGHT:
				root.right = avlLeftRightRotation(root.right)
			case RIGHT_LEFT:
				root.right = avlRightLeftRotation(root.right)
			case RIGHT_RIGHT:
				root.right = avlRightRightRotation(root.right)
			}
		}
	} else {
		leftRoot := avlInsertUtil(root.left, toInsert)
		root.left = leftRoot
		if checkImbalance(root.left) == true {
			typeImbalance := getTypeImbalance(root.left)
			switch typeImbalance {
			case LEFT_LEFT:
				root.left = avlLeftLeftRotation(root.left)
			case LEFT_RIGHT:
				root.left = avlLeftRightRotation(root.left)
			case RIGHT_LEFT:
				root.left = avlRightLeftRotation(root.left)
			case RIGHT_RIGHT:
				root.left = avlRightRightRotation(root.left)
			}
		}
	}

	var newRoot *Node = root
	if checkImbalance(root) == true {
		typeImbalance := getTypeImbalance(root)
		switch typeImbalance {
		case LEFT_LEFT:
			newRoot = avlLeftLeftRotation(root)
		case LEFT_RIGHT:
			newRoot = avlLeftRightRotation(root)
		case RIGHT_LEFT:
			newRoot = avlRightLeftRotation(root)
		case RIGHT_RIGHT:
			newRoot = avlRightRightRotation(root)
		}
	}

	return newRoot
}

func printInorder(root *Node) {
	if root == nil {
		return
	}

	printInorder(root.left)
	fmt.Println(root.key)
	printInorder(root.right)

	return
}

// -- PRIVATE API end --
