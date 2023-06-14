package main

import "fmt"

func main() {
	age := 20
	name := "Hunor"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")

	fmt.Print("new line \n")

	// Println
	fmt.Println("hello baby")
	fmt.Println("goodbye baby")
	fmt.Println("my age is", age, "and my name is", name)

	// Formatted string
	// Printf
	// %_ - format specifier
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", age, name)

  fmt.Printf("age is of type %T\n", age)
  fmt.Printf("you scored %f points! \n", 255.55)
  fmt.Printf("you scored %0.2f points! \n", 255.55)

  // Sprintf (save formatted strings)
  var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)

  fmt.Println("the saved string is:", str)
}
