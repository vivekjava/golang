package main

/*
Arrays
The type [n]T is an array of n values of type T.
ex : var a [10]int

*/

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

func main() {
	var arrayStruct [5]MyProfile
	arrayStruct[0] = MyProfile{Age: 0, Gender: "Male", FirstName: "Profile", LastName: "0"}
	arrayStruct[1] = MyProfile{Age: 1, Gender: "Male", FirstName: "Profile", LastName: "1"}
	arrayStruct[2] = MyProfile{Age: 02, Gender: "Male", FirstName: "Profile", LastName: "2"}
	arrayStruct[3] = MyProfile{Age: 03, Gender: "Male", FirstName: "Profile", LastName: "3"}
	arrayStruct[4] = MyProfile{Age: 04, Gender: "Male", FirstName: "Profile", LastName: "4"}
	for a := 0; a < 5; a++ {
		fmt.Println("Data ", arrayStruct[a])
		var temp = arrayStruct[a]
		temp.Name = temp.FirstName + temp.LastName
		arrayStruct[a] = temp
	}
	fmt.Println("Fetched Object", arrayStruct) //parser(arrayStruct)
}

/*
OUtput :
PS D:\workspace\golang\golang> go run .\ArrayStruct.go
Data  { 0 Male Profile 0}
Data  { 1 Male Profile 1}
Data  { 2 Male Profile 2}
Data  { 3 Male Profile 3}
Data  { 4 Male Profile 4}
Fetched Object [{Profile0 0 Male Profile 0} {Profile1 1 Male Profile 1} {Profile2 2 Male Profile 2} {Profile3 3 Male Profile 3} {Profile4 4 Male Profile 4}]

*/
