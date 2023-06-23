package main

import (
  "fmt"
)

func updateName(x *string) {
  *x = "David"
}  


func main() {
  name := "John"

  // updateName(name)

  // fmt.Println("memory location of name is ", &name)

  m := &name

  fmt.Println("memory address ", m)
  fmt.Println("value at memory address ", *m)

  updateName(m)

  fmt.Println("memory address ", m)
  fmt.Println("value at memory address ", *m)


  
}
