package IsMeFunc

import (
	"fmt"
	"reflect"
)

func switchAllType(iV interface{}) {
	switch nameiV := iV.(type) {
	case bool:
		fmt.Printf("bool\n")
	case string:
		fmt.Printf("string\n")
	case int, int64:
		fmt.Printf("int or int 64\n")
	case float32, float64:
		fmt.Printf("float32 or 64\n")
	default:
		fmt.Printf("what Unknown is %v\n", reflect.TypeOf(nameiV))
	}
}
