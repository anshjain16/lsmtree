package main

import (
	"fmt"
	"lsmtree/wal"
)

func main() {
	fmt.Println("hello")
	tmp_db := "/home/parzival/code/personal_projects/lsmtree/"
	write_status := wal.WriteEntry("name", "parzival", tmp_db)
	clear_status := wal.Clear(tmp_db)
	fmt.Println("the write WAL output is", write_status)
	fmt.Println("the clear WAL output is", clear_status)
}
