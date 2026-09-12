package memtable

import (
)

type Memtable struct {
	root *Node
	threshold int
}

// Create a memtable
func Create(threshold int) *Memtable {
	return &Memtable{
		root: nil,
		threshold: threshold,
	}
}



// Insert a key value pair in memtable
func Insert(memtable *Memtable, key string, value string) *Memtable{
	updatedMemtableRoot := AvlInsert(memtable.root, key, value)
	memtable.root = updatedMemtableRoot
	return memtable
}



// Search a key in memtable
func Search(memtable *Memtable, key string) (string, string){

	isPresent, value := AvlSearchKey(memtable.root, key)

	if isPresent {
		return key, value
	}

	return key, "KEY_NOT_PRESENT"

}

// Delete a key
func Delete(memtable *Memtable, key string) *Memtable{

	updatedMemtableRoot := AvlInsert(memtable.root, key, "TOMBSTONED")
	memtable.root = updatedMemtableRoot

	return memtable

}


// Flush the memtable to persistent storage (SSTABLE)
