package mysql

import (
	"reflect"
)

// reflectStruct represent a structure which stores all the values return from reflect.ValueOf()
type reflectStruct struct {
	value    reflect.Value
	valuePtr reflect.Value
	types    reflect.Type
	numField int
}

// newReflectStruct create a reflectStruct object
func newReflectStruct(structure interface{}) reflectStruct {
	valuePtr := reflect.ValueOf(structure)
	elements := valuePtr.Elem()

	return reflectStruct{
		valuePtr: valuePtr,
		value:    reflect.Indirect(valuePtr),
		types:    elements.Type(),
		numField: elements.NumField(),
	}
}

func (rs reflectStruct) isValidIndex(index int) bool {
	return index >= 0 && index < rs.value.NumField()
}

func (rs reflectStruct) validateIndex(index int) {
	if !rs.isValidIndex(index) {
		panic("array index is out of boundry")
	}
}

func (rs reflectStruct) getFiledNameByIndex(index int) string {
	rs.validateIndex(index)
	return rs.types.Field(index).Name
}

func (rs reflectStruct) getFieldByIndex(index int) reflect.Value {
	rs.validateIndex(index)
	return rs.value.FieldByIndex([]int{index})
}

func (rs reflectStruct) getFieldAddressByIndex(index int) interface{} {
	field := rs.getFieldByIndex(index)
	if !field.CanAddr() {
		panic("the passed field is not addressable")
	}
	return field.Addr().Interface()
}
