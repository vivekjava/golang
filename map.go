/*
	Reference : https://gobyexample.com/maps
*/
package main
import (
	"fmt"
)
func main(){
	mymap := make(map[string]int32)
	mymap["age"] = 21 
	mymap["age"] = 24
	//mymap["detail"] = make(map[string]string)
	delete(mymap,"age")
	fmt.Println("Map",len(mymap))
	fmt.Println("Map",mymap["age"])
}