package memtable

import (
	"testing"
	"fmt"
)

// test getAvlTreeHeight
// TODO: implement a random height test
func TestGetAvlTreeHeight(t *testing.T)  {

	// test 0 height tree
	zeroHeight := getAvlTreeHeight(nil)
	if zeroHeight != 0 {
		t.Errorf("Test for zero height tree failed")
	}

	// test 1 height tree
	oneHeightTree := createNode("testKey", "testValue2")
	oneHeight := getAvlTreeHeight(oneHeightTree)
	if oneHeight != 1 {
		t.Errorf("Test for one height tree failed")
	}

}


// test Imbalance
func TestAvlTreeImbalanceCheck(t *testing.T) {

	// create a test nodes
	nodeOne := createNode("keyOne", "valueOne")
	nodeTwo := createNode("keyTwo", "valueTwo")
	nodeThree := createNode("keyThree", "valueThree")


	// create a balanced tree
	nodeTwo.left = nodeOne
	nodeTwo.left = nodeThree

	imbalancedNode := checkImbalance(nodeTwo)
	if imbalancedNode != false {
		t.Errorf("checkImbalanc failed to detect a balanced tree")
	}
	fmt.Println("balanced node check passed")

	// clear all relations to reuse existing nodes
	nodeTwo.left = nil
	nodeTwo.right = nil

	nodeOne.right = nodeTwo
	nodeTwo.right = nodeThree

	iNode := checkImbalance(nodeOne)
	if iNode != true {
		t.Errorf("checkImbalance failed to detect an imbalanced tree")
	}
	fmt.Println("Imbalanced Node check passed")
}



// Test Insert
func TestAvlInsert(t *testing.T) {
	
	var memtable *Node = nil
	//node 1
	memtable = AvlInsert(memtable, "KeyOne", "ValueOne")
	printInorder(memtable)

	memtable = AvlInsert(memtable, "KeyTwo", "ValueTwo")
	printInorder(memtable)

	memtable = AvlInsert(memtable, "KeyX", "ValueX")
	printInorder(memtable)

	memtable = AvlInsert(memtable, "KeyOne", "NewValue")
	fmt.Println(memtable.data)

	memtable = AvlInsert(memtable, "KeyY", "ValueY")
	memtable = AvlInsert(memtable, "KeyZ", "ValueZ")
	printInorder(memtable)
	memtable = AvlInsert(memtable, "1", "1")
	printInorder(memtable)

	_, value := AvlSearchKey(memtable, "KeyOne")
	fmt.Println(value)

}



