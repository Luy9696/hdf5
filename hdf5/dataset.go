package hdf5

import (
	"log"

	"gonum.org/v1/hdf5"
)

type DataSet struct {
	*hdf5.Dataset
}

func CreateDataset(group *hdf5.CommonFG, datasetName string, space []uint) (*DataSet, error) {
	// 30x30 데이터 공간 생성
	dspace, err := hdf5.CreateSimpleDataspace(space, nil)
	if err != nil {
		log.Fatalf("Failed to create dataspace: %s", err)
		return nil, err
	}
	defer dspace.Close()

	// 데이터셋 생성
	dset, err := group.CreateDataset(datasetName, hdf5.T_NATIVE_INT, dspace)
	if err != nil {
		log.Fatalf("Failed to create dataset: %s", err)
		return nil, err
	}
	// defer dset.Close()
	return &DataSet{dset}, err
}

func (ds *DataSet) AddDatasetAttribute(attributeName string, attrData []int32) error {
	space, err := hdf5.CreateSimpleDataspace([]uint{1}, nil) // 1차원 데이터 공간
	if err != nil {
		log.Fatalf("Failed to create dataspace: %s", err)
		return err
	}
	defer space.Close()

	// 속성 생성
	attr, err := ds.CreateAttribute(attributeName, hdf5.T_NATIVE_INT, space)
	if err != nil {
		log.Fatalf("Failed to create attribute: %s", err)
		return err
	}
	defer attr.Close()

	// 속성 데이터 쓰기
	err = attr.Write(&attrData[0], hdf5.T_NATIVE_INT)
	if err != nil {
		log.Fatalf("Failed to write attribute: %s", err)
		return err
	}
	return nil
}
