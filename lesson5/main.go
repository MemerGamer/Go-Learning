package main

import (
  "fmt"
)

func main() {
  // arrays
  // var ages [3]int = [3]int{20, 25, 30}
  var ages = [3]int{20, 25, 30}

  names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
  names[1] = "Luigi"

  fmt.Println(ages, len(ages))
  fmt.Println(names, len(names))

  // slices (use arrays under the hood)
  var scores = []int{100, 50, 60}
  scores[2] = 25

  // you can append to slices but not arrays
  scores = append(scores, 85)

  fmt.Println(scores, len(scores))

  // slice ranges
  rangeOne := names[1:3]
  rangeTwo := names[2:]
  // the last index is not included

  rangeThree := names[:3]

  fmt.Println(rangeOne)
  fmt.Println(rangeTwo)
  fmt.Println(rangeThree)

  rangeOne = append(rangeOne, "Koopa")
  fmt.Println(rangeOne)

}

