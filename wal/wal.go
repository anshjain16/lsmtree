package wal

import (
	"os"
	"errors"
	"fmt"
	"encoding/csv"
	"lsmtree/util"
)

const wal_filename string = "WAL.csv"

// Create a WAL
/* 
	- create the CSV at the database location
	- nothing much to do
*/
func Create(dest string) util.Status {

	// Need to check if the destination exist or not
	// What if the destination does not exist -> do we create it? No right? -> this is a helper class -> 
	// this should not take care if dir exist or not
	// so we do not create the folder

	// TODO: create a new status type to return if dir not found
	// check if the dest exist or not
	// DEBUG LOG
	fmt.Println(dest)
	_, check_dest_path := os.Lstat(dest)
	if errors.Is(check_dest_path, os.ErrNotExist) {
		// DEBUG
		fmt.Println("databse does not exist", check_dest_path)
		return util.STATUS_NOT_OK
	}
	
	// the database path is existing, goood to create the WAL file
	file_path := dest + wal_filename
	wal_file, wal_file_create_err := os.Create(file_path) 
	
	// if the file creation failed, retunr NOT_OK
	if wal_file_create_err != nil {
		fmt.Println("file creation failed")
		return util.STATUS_NOT_OK
	}
	defer wal_file.Close()
	return util.STATUS_OK

}

// Write on a WAL
func WriteEntry(key string, value string, dest string) util.Status {

	wal_filepath := dest + wal_filename

	// Check if the WAL file exist or not
	_, wal_file_exist := os.Stat(wal_filepath)
	if errors.Is(wal_file_exist, os.ErrNotExist) {
		// DEBUG LOG
		fmt.Println("WAL file does not exist", wal_file_exist)
		return util.STATUS_NOT_OK
	}

	// The file exist -> open it(IN WRITE MODE) and append the entry
	wal_file, open_err := os.OpenFile(wal_filepath, os.O_WRONLY|os.O_APPEND, 0644)
	if open_err != nil {
		// DEBUG LOG
		fmt.Println("WAL file failed to open")

		return util.STATUS_NOT_OK
	}
	defer wal_file.Close()

	file_writer := csv.NewWriter(wal_file)
	defer file_writer.Flush()
	entry := []string{key, value}
	fmt.Println(entry)
	write_err := file_writer.Write(entry)
	if write_err != nil {
		// DEBUG LOG
		fmt.Println("Write failure to the WAl file:", write_err)

		return util.STATUS_NOT_OK
	}

	return util.STATUS_OK

}



// Clear a WAL
func Clear(dest string) util.Status {

	// A hacky way to do this is just open the file using os.Create -> this would create a new file
	// even if it already exist

	// check if the wal file exist or not
	wal_filepath := dest + wal_filename
	_, wal_file_exist := os.Stat(wal_filename)
	if errors.Is(wal_file_exist, os.ErrNotExist) {
		// DEBUG LOG
		fmt.Println("WAL file does not exist", wal_file_exist)

		return util.STATUS_NOT_OK
	}	

	wal_file, create_err := os.Create(wal_filepath)
	if create_err != nil {
		// DEBUG LOG
		fmt.Println("failed to run the create op on existing file")

		return util.STATUS_NOT_OK
	}
	defer wal_file.Close()

	return util.STATUS_OK
}




// Recover
//TODO: need to implement this function after memtable implementation is complete
