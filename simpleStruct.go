package main

import (
	"fmt"
)

/* Simple MyProfile Object */
type MyProfile struct {
	Name      string
	Age       int
	Gender    string
	FirstName string
	LastName  string
}

func parser(inp MyProfile) MyProfile {
	var result = inp
	result.Name = inp.FirstName + inp.LastName
	return result
}

func main() {
	var mp = MyProfile{Name: "Your Name", Age: 24, Gender: "Male", FirstName: "vivek", LastName: "Ravichandran"}
	fmt.Println("Fetched Object", parser(mp))
}

/*
Output :
PS D:\workspace\golang\golang> go run .\simpleStruct.go
Fetched Object {vivekRavichandran 24 Male vivek Ravichandran}

*/
