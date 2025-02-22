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

// TODO:Int,Float,String 합칠 것
func (ds *DataSet) AddDatasetAttributeInt(attributeName string, attrData []int32) error {
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

func (ds *DataSet) AddDatasetAttributeFloat(attributeName string, attrData []float64) error {
	space, err := hdf5.CreateSimpleDataspace([]uint{1}, nil)
	if err != nil {
		log.Fatalf("Failed to create dataspace: %s", err)
		return err
	}
	defer space.Close()

	attr, err := ds.CreateAttribute(attributeName, hdf5.T_NATIVE_DOUBLE, space)
	if err != nil {
		log.Fatalf("Failed to create attribute: %s", err)
		return err
	}
	defer attr.Close()

	// Preparing data for writing based on type

	err = attr.Write(&attrData[0], hdf5.T_NATIVE_DOUBLE)
	if err != nil {
		log.Fatalf("Failed to write attribute: %s", err)
		return err
	}

	return nil
}
func (ds *DataSet) AddDatasetAttributeString(attributeName string, attrData []string) error {
	space, err := hdf5.CreateSimpleDataspace([]uint{1}, nil)
	if err != nil {
		log.Fatalf("Failed to create dataspace: %s", err)
		return err
	}
	defer space.Close()

	attr, err := ds.CreateAttribute(attributeName, hdf5.T_GO_STRING, space)
	if err != nil {
		log.Fatalf("Failed to create attribute: %s", err)
		return err
	}
	defer attr.Close()

	// Preparing data for writing based on type

	err = attr.Write(&attrData[0], hdf5.T_GO_STRING)
	if err != nil {
		log.Fatalf("Failed to write attribute: %s", err)
		return err
	}

	return nil
}

func (ds *DataGroup) GetDatasetAttribute(attributeName string) string {
	// 속성 열기 (속성의 이름을 알고 있어야 합니다)
	attr, err := ds.OpenAttribute(attributeName)
	if err != nil {
		log.Fatalf("Failed to open attribute "+attributeName+": %s", err)
	}
	defer attr.Close()

	// 속성 타입에 따른 적절한 변수 타입 선언
	var value string

	// 속성 읽기
	if err := attr.Read(&value, hdf5.T_GO_STRING); err != nil {
		log.Fatalf("Failed to read attribute "+attributeName+": %s", err)
	}
	return value
}
