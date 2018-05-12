//switch practice
package main

import "fmt"

func main(){
  word := "dog"

  switch word {
    case "apple":
      fmt.Println("this is a apple")
    case "book":
      fmt.Println("this is a apple")
    case "candy","chocolate":  //注意可以多重條件
      fmt.Println("It's so sweet")
    case "dog":
      fmt.Println("this is a dog")
      fallthrough  //往下執行下一個case. 每個case相當於有一個break
    case "fog":
      fmt.Println("it's fog")
    case "zoo":
      fmt.Println("this is a zoo")
    default:
      fmt.Println("what is this?")
  }

}