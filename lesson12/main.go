package main

import (
  "fmt"
)

func main() {

        // key: string, value: float64
  menu := map[string]float64{
    "eggs":  1.25,
    "bacon": 2.99,
    "toast": 0.99,
    "coffee": 1.50,
    "soup": 3.99,
    "salad": 5.99,
    "pizza": 7.99,
    "pie": 3.99,
    "toffee pudding": 4.99,
  }

  fmt.Println(menu)
  fmt.Println(menu["pie"])

  // looping maps 

  for key, value := range menu{
    fmt.Println(key,"-", value)
  }

  // ints as key type
  phonebook := map[int]string{
    1234567890: "John",
    1234567891: "Jane",
    1234567892: "Jill",
  }

  fmt.Println(phonebook)
  fmt.Println(phonebook[1234567890])

  // update value 

  phonebook[1234567890] = "John Doe"
  fmt.Println(phonebook)

  phonebook[1234567893] = "Jack"
  fmt.Println(phonebook)
}
