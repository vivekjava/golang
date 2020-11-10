package main

/*
	Original Author : https://medium.com/swlh/high-performance-string-building-in-go-golang-3fd99b9ca856
*/
import (
	"fmt"
)

func minRemoveToMakeValid(s string) string {

	stack := make([]int, 0)
	toRemove := make([]int, 0)

	for i := 0; i < len(s); i++ {

		if string(s[i]) == "(" {
			stack = append(stack, i)
		} else if string(s[i]) == ")" && len(stack) == 0 {
			toRemove = append(toRemove, i)
		} else if string(s[i]) == ")" && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	newString := ""
	toRemove = append(toRemove, stack...)

	removeMap := make(map[int]bool)

	for i := 0; i < len(toRemove); i++ {
		removeMap[toRemove[i]] = true
	}

	for i := 0; i < len(s); i++ {
		if !removeMap[i] {
			newString += string(s[i])
		}
	}

	return newString
}

func main() {
	s := minRemoveToMakeValid("lee(t(c)o)de)")
	fmt.Println(s)
}
