package main

import (
	"education/hdf5"
	"fmt"
	"log"
)

func main() {
	// HDF5 파일 열기

	fname := "example.h5"
	file, _ := hdf5.OpenFile(fname)

	// 'ImageData' 그룹 열기
	group, err := hdf5.OpenDataGroup(&file.CommonFG, "ImageData")
	if err != nil {
		log.Fatalf("Failed to open group 'ImageData': %s", err)
	}
	defer group.Close()

	attr1 := group.GetDataGroupAttribute("name")

	fmt.Println(attr1)

}
