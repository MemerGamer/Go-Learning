package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
  greeting := "Hello everyone!"
  
  // check if a string contains a substring
  fmt.Println(strings.Contains(greeting, "Hello!"))

  // replaces all instances of a substring  
  fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

  // original string is not modified
  fmt.Println("Original string value: ", greeting)

  // uppercase 
  fmt.Println(strings.ToUpper(greeting))  
  
  // index of a substring
  fmt.Println(strings.Index(greeting, "eve"))

  // split a string
  fmt.Println(strings.Split(greeting, " "))

  ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

  // sort an array
  sort.Ints(ages) 

  fmt.Println(ages)

  // search for an element in a sorted array
  index := sort.SearchInts(ages, 30)

  fmt.Println(index)

  names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
  
  // sort an array
  sort.Strings(names)
  fmt.Println(names)

  // search for an element in a sorted array
  fmt.Println(sort.SearchStrings(names, "bowser"))

}


