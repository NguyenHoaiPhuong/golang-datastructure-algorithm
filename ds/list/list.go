package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("List is a set of element with the same data type")

	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)
	// intList.PushBack("Akagi") // This line will cause error

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
