package main
// tags and reflection
import (
	"fmt"
	"reflect"
)
import "encoding/json"

type TagExample struct{
	Name string `json:"name" validation:"only alphabet"`     // it is not single quote and it is `   > very important
	Age  int `json:"age"`
	//This is because only fields starting with a capital letter are exported,
	// or in other words visible outside the current package (and in the json package in this case). > very important
}


func main() {
	dataType := reflect.TypeOf(TagExample{})  // get the tags through reflection
	fieldTag, _ := dataType.FieldByName("Name")  // get the field by name
	println(fieldTag.Tag)
	s, _ := fieldTag.Tag.Lookup("validation")
	println(s)
	s2, _ := fieldTag.Tag.Lookup("json")
	println(s2)


	//real world examples:
	tagExample := TagExample{"Gowtham", 27}
	jsonData, err := json.Marshal(tagExample)
	if err == nil {
		fmt.Printf(string(jsonData))
	}




}
