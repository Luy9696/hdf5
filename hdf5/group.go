package hdf5

import (
	"log"

	"gonum.org/v1/hdf5"
)

type DataGroup struct {
	*hdf5.Group
}

func CreateDataGroup(commonFg *hdf5.CommonFG, groupName string) (*DataGroup, error) {
	// 그룹 생성
	group, err := commonFg.CreateGroup(groupName)
	if err != nil {
		log.Fatalf("Failed to create group: %s", err)
	}
	// defer group.Close()

	return &DataGroup{group}, err
}

func (dg *DataGroup) AddGroupAttributeInt(attributeName string, attrData []int32) error {
	space, err := hdf5.CreateSimpleDataspace([]uint{1}, nil) // 1차원 데이터 공간
	if err != nil {
		log.Fatalf("Failed to create dataspace: %s", err)
		return err
	}
	defer space.Close()

	// 속성 생성
	attr, err := dg.CreateAttribute(attributeName, hdf5.T_NATIVE_INT, space)
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

func (dg *DataGroup) AddGroupAttributeFloat(attributeName string, attrData []float64) error {
	space, err := hdf5.CreateSimpleDataspace([]uint{1}, nil)
	if err != nil {
		log.Fatalf("Failed to create dataspace: %s", err)
		return err
	}
	defer space.Close()

	attr, err := dg.CreateAttribute(attributeName, hdf5.T_NATIVE_DOUBLE, space)
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
func (dg *DataGroup) AddGroupAttributeString(attributeName string, attrData []string) error {
	space, err := hdf5.CreateSimpleDataspace([]uint{1}, nil)
	if err != nil {
		log.Fatalf("Failed to create dataspace: %s", err)
		return err
	}
	defer space.Close()

	attr, err := dg.CreateAttribute(attributeName, hdf5.T_GO_STRING, space)
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
