package tools

import (
	"fmt"
	"os"
	"reflect"
)

func LoadEnvByStruct(s interface{}) {
	// 确保传入的是一个指针
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		fmt.Println("Error: Expected a pointer to a struct")
		return
	}

	// 获取指针指向的值
	elem := v.Elem()

	// 确保这个值是一个结构体
	if elem.Kind() != reflect.Struct {
		fmt.Println("Error: Expected a pointer to a struct")
		return
	}

	for i := 0; i < elem.NumField(); i++ {

		typeField := elem.Type().Field(i)
		valueField := elem.Field(i)

		filed := typeField.Tag.Get("envField")

		valueField.SetString(os.Getenv(filed))
	}
}
