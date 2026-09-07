package wal

import (
	"testing"
	"os"
	"lsmtree/util"
	"errors"
	"encoding/csv"
)

// Test create function
func TestCreate(t *testing.T) {

	// Ask to create a WAL file at the OS's tmp dir
	// create temp database dir
	tempDir, err := os.MkdirTemp("", "walTest")
	if err != nil {
		t.Errorf("failed to create temp test dir")
	}
	defer os.RemoveAll(tempDir)
	tempDb := tempDir + "/"
	
	// Temp databse dir created
	// Create the wal file

	walCreateStatus := Create(tempDb)
	if walCreateStatus != util.STATUS_OK {
		t.Errorf("failed to create WAL file")
	}
	
	// A final check using stat
	_, walFileExist := os.Stat(tempDb + "WAL.csv")
	if errors.Is(walFileExist, os.ErrNotExist) {
		t.Errorf("failed to create WAL file even though Create() returned STATUS_OK")
	}
}


// Test WriteEntry
func TestWriteEntry(t *testing.T) {

	// Create db folder and a WAL inside it
	tempDir, err := os.MkdirTemp("", "walTest")
	if err != nil {
		t.Errorf("failed to create temp database")
	}
	defer os.RemoveAll(tempDir)
	
	// Convert the tempDir to DB
	tempDb := tempDir + "/"

	// Create a file WAL.csv inside the database
	walCreateStatus := Create(tempDb)
	if walCreateStatus != util.STATUS_OK {
		t.Errorf("failed to create a WAL file")
	}

	//Try to write a entry
	walWriteStatus := WriteEntry("TestKey1", "TestValue1", tempDb)
	if walWriteStatus != util.STATUS_OK {
		t.Errorf("failed to write a entry in the WAL")
	}

	// Verify by reading the WAL
	csvFile, err := os.Open(GetPath(tempDb))
	if err != nil {
		t.Errorf("failed to read the WAL file")
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	testRecords, err := csvReader.ReadAll()
	if err != nil {
		t.Errorf("failed to read the WAL records")
	}
	if testRecords[0][0] != "TestKey1" && testRecords[0][1] != "TestValue1" {
		t.Errorf("WAL entry do not match with the test value: TestKey1->TestValue1")
	}
}

// Test Clear
func TestClear(t *testing.T) {

	// Create a DB
	tempDir, err := os.MkdirTemp("", "walTest")
	if err != nil {
		t.Errorf("failed to create temp database")
	}
	defer os.RemoveAll(tempDir)

	tempDb := tempDir + "/"

	// Create a WAL
	walCreateStatus := Create(tempDb)
	if walCreateStatus != util.STATUS_OK {
		t.Errorf("failed to create WAL file")
	}

	// Write some entries
	walWriteStatus1 := WriteEntry("TestKey1", "TestKey1", tempDb)
	if walWriteStatus1 != util.STATUS_OK {
		t.Errorf("failed to write record in WAL")
	}
	walWriteStatus2 := WriteEntry("TestKey2", "TestKey2", tempDb)
	if walWriteStatus2 != util.STATUS_OK {
		t.Errorf("failed to write record in WAL")
	}
	walWriteStatus3 := WriteEntry("TestKey3", "TestKey3", tempDb)
	if walWriteStatus3 != util.STATUS_OK {
		t.Errorf("failed to write record in WAL")
	}

	// Clear the WAL
	walClearStatus := Clear(tempDb)
	if walClearStatus != util.STATUS_OK {
		t.Errorf("failed to clear the WAL")
	}

	// Verify by reading the WAL
	csvFile, err := os.Open(GetPath(tempDb))
	if err != nil {
		t.Errorf("failed to read the WAL file")
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	testRecords, err := csvReader.ReadAll()
	if err != nil {
		t.Errorf("failed to read the WAL records")
	}
	if len(testRecords) != 0 {
		t.Errorf("Records still present in WAL even if Clear() returned STATUS_OK")
	}

}

