package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Eko"}

	var SampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(SampleType.NumField())
	fmt.Println(SampleType.Field(0).Name)
	fmt.Println(SampleType.Field(0).Tag.Get("required"))
	fmt.Println(SampleType.Field(0).Tag.Get("max"))
	fmt.Println(SampleType.Field(0).Tag.Get("min"))

	sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := ContohLagi{"Eko", "Oke"}
	fmt.Println(IsValid(contoh))
}
