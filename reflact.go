package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `required:"true" max:"100"`
	Age  int    `required:"true" max:"3"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {

	Sample := User{"Kresna", 19}

	SampleType := reflect.TypeOf(Sample)

	fmt.Println(SampleType)
	fmt.Println(SampleType.Field(1).Name)
	fmt.Println(SampleType.Field(0).Tag.Get("required"))
	fmt.Println(SampleType.Field(0).Tag.Get("max"))

	fmt.Println(IsValid(Sample))

}
