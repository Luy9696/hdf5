package hdf5

import (
	"errors"

	"github.com/go-hdf5/hdf5"
	"gonum.org/v1/hdf5"
)

// getHDF5DataType는 Go 타입을 HDF5 데이터 타입으로 매핑하는 함수
func getHDF5DataType(data interface{}) (*hdf5.Datatype, error) {

	switch data.(type) {
	case int:
		return hdf5.T_NATIVE_INT, nil
	case float32, float64:
		return hdf5.T_NATIVE_FLOAT, nil
	// case reflect.Float64:
	// 	return hdf5.T_NATIVE_DOUBLE, nil
	// 추가 타입 지원은 여기에 구현
	default:
		return nil, errors.New("unsupported data type")
	}
}

// // getHDF5DataType는 Go 타입을 HDF5 데이터 타입으로 매핑하는 함수
// func getHDF5DataType(data interface{}) (*hdf5.Datatype, error) {
// 	switch reflect.TypeOf(data).Kind() {
// 	case reflect.Int, reflect.Int32:
// 		return hdf5.T_NATIVE_INT, nil
// 	case reflect.Int64:
// 		return hdf5.T_NATIVE_INT64, nil
// 	case reflect.Float32:
// 		return hdf5.T_NATIVE_FLOAT, nil
// 	case reflect.Float64:
// 		return hdf5.T_NATIVE_DOUBLE, nil
// 	// 추가 타입 지원은 여기에 구현
// 	default:
// 		return nil, errors.New("unsupported data type")
// 	}
// }
