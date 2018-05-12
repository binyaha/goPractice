//if practice
package main

import "fmt"

func main(){
  num := 2

  //basic if
  if num%2 == 0 {
    fmt.Println("number is even")
  } else {
    fmt.Println("number is odd")
  }

  //在條件是聲明變量 只在這個if內
  //go的else if語法: else if {}
  if num2 := 3; num2%2 == 0 {
    fmt.Println("number2 is even")
  } else {
    fmt.Println("number2 is odd")
  }

}