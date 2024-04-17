package main

import (
	"fmt"
	"log"

	"education/hdf5"
)

func main() {
	// HDF5 파일 생성
	fname := "example.h5"

	file, err := hdf5.CreateFile(fname)
	if err != nil {
		fmt.Println("Failed to create File", file)
	}

	ImageData, err := hdf5.CreateDataGroup(&file.CommonFG, "ImageData")
	ImageData.AddGroupAttributeInt("Type", []int32{20})
	ImageData.AddGroupAttributeString("name", []string{"ImageData"})
	ImageData.AddGroupAttributeInt("byte_length", []int32{8858})
	ImageData.AddGroupAttributeInt("sorted", []int32{1})
	ImageData.AddGroupAttributeInt("closed", []int32{0})
	ImageData.AddGroupAttributeInt("num_tags", []int32{5})
	ImageData.AddGroupAttributeInt("data_offset", []int32{438438})
	// ImageData.Close()
	if err != nil {
		fmt.Println("Failed to create group")
	}

	Data, err := hdf5.CreateDataset(&ImageData.CommonFG, "Data", []uint{1, 1})
	Data.AddDatasetAttributeInt("byte_length", []int32{8828})
	Data.AddDatasetAttributeString("name", []string{"Data"})
	Data.AddDatasetAttributeInt("array_length", []int32{3})
	Data.AddDatasetAttributeInt("data_type_code", []int32{20})
	Data.AddDatasetAttributeInt("header_offset", []int32{438889})
	Data.AddDatasetAttributeInt("data_offset", []int32{438904})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	DataType, err := hdf5.CreateDataset(&ImageData.CommonFG, "DataType", []uint{1, 1})
	DataType.AddDatasetAttributeInt("byte_length", []int32{24})
	DataType.AddDatasetAttributeString("name", []string{"DataType"})
	DataType.AddDatasetAttributeInt("array_length", []int32{1})
	DataType.AddDatasetAttributeInt("data_type_code", []int32{5})
	DataType.AddDatasetAttributeInt("header_offset", []int32{447132})
	DataType.AddDatasetAttributeInt("data_offset", []int32{447151})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	PixelDepth, err := hdf5.CreateDataset(&ImageData.CommonFG, "PixelDepth", []uint{1, 1})
	PixelDepth.AddDatasetAttributeInt("byte_length", []int32{24})
	PixelDepth.AddDatasetAttributeString("name", []string{"PixelDepth"})
	PixelDepth.AddDatasetAttributeInt("array_length", []int32{1})
	PixelDepth.AddDatasetAttributeInt("data_type_code", []int32{5})
	PixelDepth.AddDatasetAttributeInt("header_offset", []int32{447241})
	PixelDepth.AddDatasetAttributeInt("data_offset", []int32{447262})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}

	Calibrations, err := hdf5.CreateDataGroup(&ImageData.CommonFG, "Calibrations")
	Calibrations.AddGroupAttributeInt("Type", []int32{20})
	Calibrations.AddGroupAttributeString("name", []string{"Calibrations"})
	Calibrations.AddGroupAttributeInt("byte_length", []int32{428})
	Calibrations.AddGroupAttributeInt("sorted", []int32{1})
	Calibrations.AddGroupAttributeInt("closed", []int32{0})
	Calibrations.AddGroupAttributeInt("num_tags", []int32{3})
	Calibrations.AddGroupAttributeInt("data_offset", []int32{438471})
	if err != nil {
		fmt.Println("Failed to Add Group")
	}
	Dimensions, err := hdf5.CreateDataGroup(&ImageData.CommonFG, "Dimensions")
	Dimensions.AddGroupAttributeInt("Type", []int32{20})
	Dimensions.AddGroupAttributeString("name", []string{"Dimensions"})
	Dimensions.AddGroupAttributeInt("byte_length", []int32{45})
	Dimensions.AddGroupAttributeInt("sorted", []int32{0})
	Dimensions.AddGroupAttributeInt("closed", []int32{0})
	Dimensions.AddGroupAttributeInt("num_tags", []int32{1})
	Dimensions.AddGroupAttributeInt("data_offset", []int32{447206})
	if err != nil {
		fmt.Println("Failed to Add Group")
	}
	unnamedTags, err := hdf5.CreateDataset(&Dimensions.CommonFG, "unnamedTags", []uint{1, 1})
	unnamedTags.AddDatasetAttributeInt("type", []int32{21})
	unnamedTags.AddDatasetAttributeInt("byte_length", []int32{24})
	unnamedTags.AddDatasetAttributeInt("array_length", []int32{1})
	unnamedTags.AddDatasetAttributeInt("data_type_code", []int32{5})
	unnamedTags.AddDatasetAttributeInt("header_offset", []int32{438502})
	unnamedTags.AddDatasetAttributeInt("data_offset", []int32{447217})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}

	Brightness, err := hdf5.CreateDataGroup(&Calibrations.CommonFG, "Brightness")
	Brightness.AddGroupAttributeInt("Type", []int32{20})
	Brightness.AddGroupAttributeString("name", []string{"Brightness"})
	Brightness.AddGroupAttributeInt("byte_length", []int32{155})
	Brightness.AddGroupAttributeInt("sorted", []int32{1})
	Brightness.AddGroupAttributeInt("closed", []int32{0})
	Brightness.AddGroupAttributeInt("num_tags", []int32{3})
	Brightness.AddGroupAttributeInt("data_offset", []int32{438502})
	if err != nil {
		fmt.Println("Failed to Add Group")
	}

	Origin2, err := hdf5.CreateDataset(&Brightness.CommonFG, "Origin", []uint{1, 1})
	Origin2.AddDatasetAttributeInt("Type", []int32{21})
	Origin2.AddDatasetAttributeInt("byte_length", []int32{24})
	Origin2.AddDatasetAttributeString("name", []string{"Origin"})
	Origin2.AddDatasetAttributeInt("array_length", []int32{1})
	Origin2.AddDatasetAttributeInt("data_type_code", []int32{6})
	Origin2.AddDatasetAttributeInt("header_offset", []int32{438502})
	Origin2.AddDatasetAttributeInt("data_offset", []int32{438519})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	Scale2, err := hdf5.CreateDataset(&Brightness.CommonFG, "Scale", []uint{1, 1})
	Scale2.AddDatasetAttributeInt("type", []int32{21})
	Scale2.AddDatasetAttributeInt("byte_length", []int32{24})
	Scale2.AddDatasetAttributeString("name", []string{"Scale"})
	Scale2.AddDatasetAttributeInt("array_length", []int32{1})
	Scale2.AddDatasetAttributeInt("data_type_code", []int32{6})
	Scale2.AddDatasetAttributeInt("header_offset", []int32{438543})
	Scale2.AddDatasetAttributeInt("data_offset", []int32{438559})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	Units2, err := hdf5.CreateDataset(&Brightness.CommonFG, "Units", []uint{1, 1})
	Units2.AddDatasetAttributeInt("type", []int32{21})
	Units2.AddDatasetAttributeInt("byte_length", []int32{48})
	Units2.AddDatasetAttributeString("name", []string{"Units"})
	Units2.AddDatasetAttributeInt("array_length", []int32{3})
	Units2.AddDatasetAttributeInt("data_type_code", []int32{20})
	Units2.AddDatasetAttributeInt("header_offset", []int32{438583})
	Units2.AddDatasetAttributeInt("data_offset", []int32{438599})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}

	Dimension, err := hdf5.CreateDataGroup(&Calibrations.CommonFG, "Dimension")
	Dimension.AddGroupAttributeInt("Type", []int32{20})
	Dimension.AddGroupAttributeString("name", []string{"Dimension"})
	Dimension.AddGroupAttributeInt("byte_length", []int32{168})
	Dimension.AddGroupAttributeInt("sorted", []int32{0})
	Dimension.AddGroupAttributeInt("closed", []int32{0})
	Dimension.AddGroupAttributeInt("num_tags", []int32{1})
	Dimension.AddGroupAttributeInt("data_offset", []int32{438677})
	if err != nil {
		fmt.Println("Failed to Add Group")
	}
	unnamed, err := hdf5.CreateDataGroup(&Dimension.CommonFG, "unnamed")
	unnamed.AddGroupAttributeInt("Type", []int32{20})
	unnamed.AddGroupAttributeInt("byte_length", []int32{147})
	unnamed.AddGroupAttributeInt("sorted", []int32{1})
	unnamed.AddGroupAttributeInt("closed", []int32{0})
	unnamed.AddGroupAttributeInt("num_tags", []int32{3})
	unnamed.AddGroupAttributeInt("data_offset", []int32{438698})
	if err != nil {
		fmt.Println("Failed to Add Group")
	}

	Origin, err := hdf5.CreateDataset(&unnamed.CommonFG, "Origin", []uint{1, 1})
	Origin.AddDatasetAttributeInt("type", []int32{21})
	Origin.AddDatasetAttributeInt("byte_length", []int32{24})
	Origin.AddDatasetAttributeInt("array_length", []int32{1})
	Origin.AddDatasetAttributeString("name", []string{"Origin"})
	Origin.AddDatasetAttributeInt("array_length", []int32{1})
	Origin.AddDatasetAttributeInt("data_type_code", []int32{6})
	Origin.AddDatasetAttributeInt("header_offset", []int32{438698})
	Origin.AddDatasetAttributeInt("data_offset", []int32{438715})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	Scale, err := hdf5.CreateDataset(&unnamed.CommonFG, "Scale", []uint{1, 1})
	Scale.AddDatasetAttributeInt("type", []int32{21})
	Scale.AddDatasetAttributeInt("byte_length", []int32{24})
	Scale.AddDatasetAttributeInt("array_length", []int32{1})
	Scale.AddDatasetAttributeString("name", []string{"Scale"})
	Scale.AddDatasetAttributeInt("data_type_code", []int32{6})
	Scale.AddDatasetAttributeInt("header_offset", []int32{438739})
	Scale.AddDatasetAttributeInt("data_offset", []int32{438755})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}
	Units, err := hdf5.CreateDataset(&unnamed.CommonFG, "Units", []uint{1, 1})
	Units.AddDatasetAttributeInt("type", []int32{21})
	Units.AddDatasetAttributeInt("byte_length", []int32{40})
	Units.AddDatasetAttributeString("name", []string{"Units"})
	Units.AddDatasetAttributeInt("array_length", []int32{3})
	Units.AddDatasetAttributeInt("data_type_code", []int32{20})
	Units.AddDatasetAttributeInt("header_offset", []int32{438779})
	Units.AddDatasetAttributeInt("data_offset", []int32{438795})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}

	DisplayCalibratedUnits, err := hdf5.CreateDataset(&Calibrations.CommonFG, "DisplayCalibratedUnits", []uint{1, 1})
	DisplayCalibratedUnits.AddDatasetAttributeInt("type", []int32{21})
	DisplayCalibratedUnits.AddDatasetAttributeInt("byte_length", []int32{24})
	DisplayCalibratedUnits.AddDatasetAttributeString("name", []string{"DisplayCalibratedUnits"})
	DisplayCalibratedUnits.AddDatasetAttributeInt("array_length", []int32{1})
	DisplayCalibratedUnits.AddDatasetAttributeInt("data_type_code", []int32{8})
	DisplayCalibratedUnits.AddDatasetAttributeInt("header_offset", []int32{438835})
	DisplayCalibratedUnits.AddDatasetAttributeInt("data_offset", []int32{438868})
	if err != nil {
		fmt.Println("Failed to Add Dataset")
	}

	log.Println("HDF5 file, group, and dataset with data created successfully.")

	fmt.Println(file.GetFileName(), file.GetType())
}
