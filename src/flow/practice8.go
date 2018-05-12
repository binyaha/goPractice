//for practice
package main

import "fmt"

func main(){
  //basic for loop
  for i := 0; i < 4 ; i++ {
    fmt.Println(i)
  }

  //break
  for i := 0; i < 4 ; i++ {
    if i == 2 {
      break
    }
    fmt.Println(i)
  }

  //continue
  for i := 0; i < 10 ; i++ {
    if i == 3 {
      continue
    }
    fmt.Println(i)
  }
}