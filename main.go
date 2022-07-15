package main

import (
	"fmt"
	zfs "github.com/bicomsystems/go-libzfs"
)

func getList() {
	list, err := zfs.DatasetOpenAll()
	if err != nil {
		fmt.Errorf("test:", err)
	}
	print(list)
}

func main() {
	getList()
}
