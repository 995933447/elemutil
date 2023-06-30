package elemutil

import (
	"github.com/995933447/reflectutil"
	"github.com/995933447/stringhelper-go"
	"reflect"
	"sync"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Struct2MapWithUnderLineKey(s interface{}, skip ...string) map[string]interface{} {
	m := make(map[string]interface{})
	elem := reflect.ValueOf(s).Elem()

	relType := elem.Type()

	for i := 0; i < relType.NumField(); i++ {
		n := relType.Field(i).Name
		if len(skip) > 0 {
			ignore := false
			for _, v := range skip {
				if v == n {
					ignore = true
					break
				}
			}
			if ignore {
				continue
			}
		}
		key := stringhelper.Snake(n)
		m[key] = elem.Field(i).Interface()
	}

	return m
}

func Struct2MapWithUnderLineKeySkipZero(s interface{}, skip ...string) map[string]interface{} {
	m := Struct2MapWithUnderLineKey(s, skip...)

	n := map[string]interface{}{}
	for k, v := range m {
		if !reflectutil.IsBlank(reflect.ValueOf(v)) {
			n[k] = v
		}
	}

	return n
}

var camel2UnderScoreMap = map[string]string{}
var camel2UnderScoreMapMu sync.RWMutex

func Struct2MapLower(s interface{}, skip ...string) map[string]interface{} {
	m := make(map[string]interface{})
	elem := reflect.ValueOf(s).Elem()

	relType := elem.Type()

	var nameUpdate map[string]string

	camel2UnderScoreMapMu.RLock()

	for i := 0; i < relType.NumField(); i++ {
		camelName := relType.Field(i).Name
		n := stringhelper.Snake(camelName)
		if nameUpdate == nil {
			nameUpdate = map[string]string{}
		}
		nameUpdate[camelName] = n

		if len(skip) > 0 {
			ignore := false
			for _, v := range skip {
				if v == n {
					ignore = true
					break
				}
			}
			if ignore {
				continue
			}
		}
		m[n] = elem.Field(i).Interface()
	}

	camel2UnderScoreMapMu.RUnlock()

	if nameUpdate != nil {
		camel2UnderScoreMapMu.Lock()
		for k, v := range nameUpdate {
			camel2UnderScoreMap[k] = v
		}
		camel2UnderScoreMapMu.Unlock()
	}

	return m
}

func Struct2MapLowerSkipZero(s interface{}, skip ...string) map[string]interface{} {
	m := Struct2MapLower(s, skip...)

	n := map[string]interface{}{}
	for k, v := range m {
		if !reflectutil.IsBlank(reflect.ValueOf(v)) {
			n[k] = v
		}
	}

	return n
}
