package main

import (
  "fmt"
  "math"
)

func sayGreeting(n string){
  fmt.Printf("Good morning %v\n", n);
}

func sayBye(n string){
  fmt.Printf("Goodbye %v\n", n);
}

func cycleNames(n []string, f func(string))  {
  for _, v := range n {
    f(v);
  }
}

func circleArea(r float64) float64 {
  return math.Pi * r * r;
}

func main() {
  /* sayGreeting("Mario");
  sayGreeting("Luigi");
  sayBye("Mario");
  sayBye("Luigi");
  */
  cycleNames([]string{"Mario", "Luigi", "Peach"}, sayGreeting);
  cycleNames([]string{"Mario", "Luigi", "Peach"}, sayBye);
  
  a1 := circleArea(10.5);
  a2 := circleArea(15);

  fmt.Println(a1, a2);
  fmt.Printf("Circle 1 is %0.3f and Circle 2 is %0.3f\n", a1, a2);
}
