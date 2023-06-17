package main 

import (
  "fmt"
)

func main() {
  // x := 0

  // it is like a while loop but with the for keyword

  // for x < 5 {
  //   fmt.Println("Value of x is: ", x)
  //   x++
  // }

  // traditional for loop
  // for i := 0; i < 5; i++ {
  //   fmt.Println("Value of i is: ", i)
  // }

  names := []string{"John", "Paul", "George", "Ringo"}

  // for i := 0; i < len(names); i++ {
  //   fmt.Println(names[i])
  // }

  // for in loop style loop
  for index, value := range names {
    fmt.Printf("The value at index %v is %v\n", index, value)
  }

  fmt.Println("")

  // if you don't want to use the index
  for _, value := range names {
    fmt.Printf("The value at is %v\n", value)
    value = "new string"
  }
  // value is a copy of the value at the index
  // so it will not change the value in the array
  fmt.Println("")
  fmt.Println(names)

}
