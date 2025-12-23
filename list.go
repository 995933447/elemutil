package elemutil

import (
	"errors"
	"fmt"
	"reflect"
)

func InUint32List(val uint32, list []uint32) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InUint64List(val uint64, list []uint64) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InUint16List(val uint16, list []uint16) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InUintList(val uint, list []uint) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InInt32List(val int32, list []int32) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InInt64List(val int64, list []int64) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InInt16List(val int16, list []int16) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InIntList(val int, list []int) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InFloat32List(val float32, list []float32) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InFloat64List(val float64, list []float64) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func InStringList(val string, list []string) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func PluckUint64(list interface{}, fieldName string) ([]uint64, error) {
	var result []uint64
	vo := reflect.ValueOf(list)
	switch vo.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			elem := vo.Index(i)

			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			// 如果某一个元素的nil，跳过从这个元素中获取数据
			if !elem.IsValid() {
				continue
			}

			if elem.Kind() != reflect.Struct {
				panic("element not struct")
			}

			f := elem.FieldByName(fieldName)
			if !f.IsValid() {
				return nil, fmt.Errorf("struct missed field %s", fieldName)
			}

			if f.Kind() != reflect.Uint64 {
				return nil, fmt.Errorf("struct element %s type required uint64", fieldName)
			}

			result = append(result, f.Uint())
		}
	default:
		return nil, fmt.Errorf("required list of struct type")
	}

	return result, nil
}

func PluckInt64(list interface{}, fieldName string) ([]int64, error) {
	var result []int64
	vo := reflect.ValueOf(list)
	switch vo.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			elem := vo.Index(i)

			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			// 如果某一个元素的nil，跳过从这个元素中获取数据
			if !elem.IsValid() {
				continue
			}

			if elem.Kind() != reflect.Struct {
				return nil, fmt.Errorf("element not struct")
			}

			f := elem.FieldByName(fieldName)
			if !f.IsValid() {
				return nil, fmt.Errorf("struct missed field %s", fieldName)
			}

			if f.Kind() != reflect.Int64 {
				return nil, fmt.Errorf("struct element %s type required int64", fieldName)
			}

			result = append(result, f.Int())
		}
	default:
		return nil, fmt.Errorf("required list of struct type")
	}

	return result, nil
}

func PluckUint64Set(list interface{}, fieldName string) (map[uint64]struct{}, error) {
	out, err := PluckUint64(list, fieldName)
	if err != nil {
		return nil, err
	}
	res := map[uint64]struct{}{}
	for _, v := range out {
		res[v] = struct{}{}
	}
	return res, nil
}

func PluckUint32(list interface{}, fieldName string) ([]uint32, error) {
	var result []uint32
	vo := reflect.ValueOf(list)
	switch vo.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			elem := vo.Index(i)

			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			// 如果某一个元素的nil，跳过从这个元素中获取数据
			if !elem.IsValid() {
				continue
			}

			if elem.Kind() != reflect.Struct {
				return nil, fmt.Errorf("element not struct")
			}

			f := elem.FieldByName(fieldName)
			if !f.IsValid() {
				return nil, fmt.Errorf("struct missed field %s", fieldName)
			}

			if f.Kind() != reflect.Uint32 {
				return nil, fmt.Errorf("struct element %s type required uint32", fieldName)
			}
			result = append(result, uint32(f.Uint()))
		}
	default:
		return nil, fmt.Errorf("required list of struct type")
	}
	return result, nil
}

func PluckUint32Set(list interface{}, fieldName string) (map[uint32]struct{}, error) {
	out, err := PluckUint32(list, fieldName)
	if err != nil {
		return nil, err
	}
	res := map[uint32]struct{}{}
	for _, v := range out {
		res[v] = struct{}{}
	}

	return res, nil
}

func PluckString(list interface{}, fieldName string) ([]string, error) {
	var result []string
	vo := reflect.ValueOf(list)
	switch vo.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			elem := vo.Index(i)

			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			// 如果某一个元素的nil，跳过从这个元素中获取数据
			if !elem.IsValid() {
				continue
			}

			if elem.Kind() != reflect.Struct {
				return nil, errors.New("element not struct")
			}

			f := elem.FieldByName(fieldName)
			if !f.IsValid() {
				return nil, fmt.Errorf("struct missed field %s", fieldName)
			}

			if f.Kind() != reflect.String {
				return nil, fmt.Errorf("struct element %s type required string", fieldName)
			}
			result = append(result, f.String())
		}
	default:
		return nil, errors.New("required list of struct type")
	}
	return result, nil
}

func PluckStringSet(list interface{}, fieldName string) (map[string]struct{}, error) {
	out, err := PluckString(list, fieldName)
	if err != nil {
		return nil, err
	}

	res := map[string]struct{}{}
	for _, v := range out {
		res[v] = struct{}{}
	}

	return res, nil
}

// KeyBy 根据KeyFieldName 取出 map[key]val
// 允许list长度为0
func KeyBy(list interface{}, fieldName string) (interface{}, error) {
	lv := reflect.ValueOf(list)

	switch lv.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return nil, errors.New("list required slice or array type")
	}

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	if evs.Kind() != reflect.Struct {
		return nil, errors.New("element not struct")
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		return nil, fmt.Errorf("field %s not found", fieldName)
	}

	m := reflect.MakeMapWithSize(reflect.MapOf(field.Type, ev), lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		// 如果是nil的，意味着key和value同时不存在，所以跳过不处理
		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			return nil, errors.New("element not struct")
		}

		m.SetMapIndex(elemStruct.FieldByIndex(field.Index), elem)
	}

	return m.Interface(), nil
}
