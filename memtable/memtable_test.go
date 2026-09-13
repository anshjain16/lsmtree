package memtable

import(
	"testing"
	"fmt"
)

// Test memtable creation
func TestCreate(t *testing.T) {

	// create a memtable
	memtable := Create(10)
	if memtable == nil {
		t.Errorf("failed to create memtable")
	}
}


// Test
func TestMemtable(t *testing.T) {

	var testMemtable *Memtable
	testMemtable = Create(10)

	testMemtable = Insert(testMemtable, "A", "A")
	testMemtable = Insert(testMemtable, "B", "B")
	testMemtable = Insert(testMemtable, "AAA", "AAA")

	var testKey, testValue string
	testKey, testValue = Search(testMemtable, "A")
	if testKey != "A" && testValue != "A" {
		t.Errorf("failed to insert key value pair in memtable")
	}
	testKey,testValue = Search(testMemtable, "AAA")
	if testKey != "AAA" && testValue != "AAA" {
		t.Errorf("failed to insert key value pair in memtable")
	}
	testKey, testValue = Search(testMemtable, "B")
	if testKey != "B" && testValue != "B" {
		t.Errorf("failed to insert key value pair in memtable")
	}

	testKey, testValue = Search(testMemtable, "Z")
	if testKey != "Z" && testValue != "KEY_NOT_PRESENT" {
		t.Errorf("Search returns true on a key not present in the database")
	}

	testMemtable = Delete(testMemtable, "AAA")
	testKey, testValue = Search(testMemtable, "AAA")
	if testKey != "AAA" && testValue != "TOMBSTONED" {
		t.Errorf("memtable delete failed")
	}
	fmt.Println("memtable test completed")

}


