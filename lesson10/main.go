package main 

import (
  "fmt"
  "strings"
)

func getInitials(name string) (string, string) {
  capitalizedName := strings.ToUpper(name)
  names := strings.Split(capitalizedName, " ")

  var initials []string

  for _, currentName := range names {
    initials = append(initials, string(currentName[:1]))
  }

  if len(initials) > 1 {
    return initials[0], initials[1]
  }

  return initials[0], ""
}

func main() {
  fn1, sn1 := getInitials("john doe")
  fmt.Println(fn1, sn1)


  fn2, sn2 := getInitials("claude shannon")
  fmt.Println(fn2, sn2)

  fn3, sn3 := getInitials("alan")
  fmt.Println(fn3, sn3)

}
