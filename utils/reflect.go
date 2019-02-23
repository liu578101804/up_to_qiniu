package utils

import (
	"fmt"
	"reflect"
)

func setField(obj interface{}, name string, value interface{}) error{

	structData := reflect.ValueOf(obj).Elem()
	filedValue := structData.FieldByName(name)

	if !filedValue.IsValid() {
		return fmt.Errorf("utils.setField() No such field: %s in obj ", name)
	}

	if !filedValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value ", name)
	}

	fieldType := filedValue.Type()
	val := reflect.ValueOf(value)

	valTypeStr := val.Type().String()
	fieldTypeStr := fieldType.String()

	//转换类型
	if valTypeStr == "float64" && fieldTypeStr == "int" {
		val = val.Convert(fieldType)
	//	类型不匹配
	}else if fieldType != val.Type() {
		return fmt.Errorf("utils.setField() "+ name +" Provided value type "+ valTypeStr + " didn't match obj field type " + fieldTypeStr)
	}

	filedValue.Set(val)
	return nil
}


func SetStructByJSON(obj interface{}, mapData map[string]interface{}) error {

	for key,value := range mapData {
		if err := setField(obj, key, value); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}