package main

import (
	hdf5 "education/hdf5"
	"fmt"
	"log"
)

func main() {
	// HDF5 파일 생성
	fname := "example.h5"

	file, err := hdf5.CreateFile(fname)
	if err != nil {
		fmt.Println("Failed to create File", file)
	}

	ImageData, err := hdf5.CreateDataGroup(&file.CommonFG, "ImageData")
	ImageData.AddGroupAttribute("byte_length", []int32{8858})
	ImageData.AddGroupAttribute("Type", []int32{42})
	ImageData.AddGroupAttribute("sorted", []int32{1})
	ImageData.AddGroupAttribute("closed", []int32{0})
	ImageData.AddGroupAttribute("num_tags", []int32{5})
	ImageData.AddGroupAttribute("data_offset", []int32{438438})
	// ImageData.Close()
	if err != nil {
		fmt.Println("Failed to create group")
	}

	Data, err := hdf5.CreateDataset(&ImageData.CommonFG, "Data", []uint{5, 5})
	Data.AddDatasetAttribute("byte_length", []int32{8828})
	Data.AddDatasetAttribute("array_length", []int32{3})
	Data.AddDatasetAttribute("data_type_code", []int32{20})
	Data.AddDatasetAttribute("header_offset", []int32{438889})
	Data.AddDatasetAttribute("data_offset", []int32{438904})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	DataType, err := hdf5.CreateDataset(&ImageData.CommonFG, "DataType", []uint{5, 5})
	DataType.AddDatasetAttribute("byte_length", []int32{24})
	DataType.AddDatasetAttribute("array_length", []int32{1})
	DataType.AddDatasetAttribute("data_type_code", []int32{5})
	DataType.AddDatasetAttribute("header_offset", []int32{447132})
	DataType.AddDatasetAttribute("data_offset", []int32{447151})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	PixelDepth, err := hdf5.CreateDataset(&ImageData.CommonFG, "PixelDepth", []uint{5, 5})
	PixelDepth.AddDatasetAttribute("byte_length", []int32{24})
	PixelDepth.AddDatasetAttribute("array_length", []int32{1})
	PixelDepth.AddDatasetAttribute("data_type_code", []int32{5})
	PixelDepth.AddDatasetAttribute("header_offset", []int32{447241})
	PixelDepth.AddDatasetAttribute("data_offset", []int32{447262})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	Calibrations, err := hdf5.CreateDataGroup(&ImageData.CommonFG, "Calibrations")
	Calibrations.AddGroupAttribute("Type", []int32{20})
	Calibrations.AddGroupAttribute("byte_length", []int32{428})
	Calibrations.AddGroupAttribute("sorted", []int32{1})
	Calibrations.AddGroupAttribute("closed", []int32{0})
	Calibrations.AddGroupAttribute("num_tags", []int32{3})
	Calibrations.AddGroupAttribute("data_offset", []int32{438471})
	if err != nil {
		fmt.Println("Failed to Add Group")
	}
	DisplayCalibratedUnits, err := hdf5.CreateDataset(&Calibrations.CommonFG, "DisplayCalibratedUnits", []uint{5, 5})
	DisplayCalibratedUnits.AddDatasetAttribute("byte_length", []int32{24})
	DisplayCalibratedUnits.AddDatasetAttribute("array_length", []int32{1})
	DisplayCalibratedUnits.AddDatasetAttribute("data_type_code", []int32{6})
	DisplayCalibratedUnits.AddDatasetAttribute("header_offset", []int32{438502})
	DisplayCalibratedUnits.AddDatasetAttribute("data_offset", []int32{438519})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	Brightness, err := hdf5.CreateDataGroup(&ImageData.CommonFG, "Brightness")
	Brightness.AddGroupAttribute("Type", []int32{20})
	Brightness.AddGroupAttribute("byte_length", []int32{155})
	Brightness.AddGroupAttribute("sorted", []int32{1})
	Brightness.AddGroupAttribute("closed", []int32{0})
	Brightness.AddGroupAttribute("num_tags", []int32{3})
	Brightness.AddGroupAttribute("data_offset", []int32{438502})
	if err != nil {
		fmt.Println("Failed to Add Group")
	}
	Origin, err := hdf5.CreateDataset(&Brightness.CommonFG, "Origin", []uint{5, 5})
	Origin.AddDatasetAttribute("byte_length", []int32{24})
	Origin.AddDatasetAttribute("array_length", []int32{1})
	Origin.AddDatasetAttribute("data_type_code", []int32{6})
	Origin.AddDatasetAttribute("header_offset", []int32{438502})
	Origin.AddDatasetAttribute("data_offset", []int32{438519})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	Scale, err := hdf5.CreateDataset(&Brightness.CommonFG, "Scale", []uint{5, 5})
	Scale.AddDatasetAttribute("byte_length", []int32{24})
	Scale.AddDatasetAttribute("array_length", []int32{1})
	Scale.AddDatasetAttribute("data_type_code", []int32{6})
	Scale.AddDatasetAttribute("header_offset", []int32{438543})
	Scale.AddDatasetAttribute("data_offset", []int32{438559})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	Units, err := hdf5.CreateDataset(&Brightness.CommonFG, "Units", []uint{5, 5})
	Units.AddDatasetAttribute("byte_length", []int32{48})
	Units.AddDatasetAttribute("array_length", []int32{3})
	Units.AddDatasetAttribute("data_type_code", []int32{20})
	Units.AddDatasetAttribute("header_offset", []int32{438583})
	Units.AddDatasetAttribute("data_offset", []int32{438599})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}

	log.Println("HDF5 file, group, and dataset with data created successfully.")
}
