package hdf5

import (
	"log"

	"gonum.org/v1/hdf5"
)

type File struct {
	*hdf5.File
}

func CreateFile(fileName string) (*File, error) {
	file, err := hdf5.CreateFile(fileName, hdf5.F_ACC_TRUNC)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
		return nil, err
	}
	// defer file.Close()

	return &File{file}, nil
}

func OpenFile(fileName string) (*File, error) {
	file, err := hdf5.OpenFile("example.h5", hdf5.F_ACC_RDONLY)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	// defer file.Close()

	return &File{file}, nil
}

func (f *File) GetFileName() string {
	return f.FileName()
}

func (f *File) GetType() string {
	return f.Type().String()
}
