package uttils

import (
	"fmt"
	"reflect"
)

// Greet takes argument to greet each one of them if no parameters are passed, it will greet to the world
// Capitalize the first letter to make function public
func Greet(params ...interface{}) {
	greet := "Hello "
	if len(params) > 0 {
		for _, param := range params {
			if reflect.TypeOf(param).Kind() == reflect.String {
				greet += fmt.Sprintf("%v, ", param)
			}
		}
		greet = greet[:len(greet)-2] + "!"
	} else {
		greet += "World!"
	}
	fmt.Println(greet)
}