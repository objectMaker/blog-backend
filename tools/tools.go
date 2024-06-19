package tools

import (
	"fmt"
	"os"
	"reflect"
)

func LoadEnvByStruct(s interface{}) {
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		valueField := val.Field(i)
		filed := typeField.Tag.Get("envField")
		fmt.Println(filed)
		valueField.SetString(os.Getenv(filed))
		fmt.Println(os.Getenv(filed))
		// 根据字段类型设置值
		valueField.SetString(os.Getenv(filed))

	}
}
