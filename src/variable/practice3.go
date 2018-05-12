//array pracitce
package main

import "fmt"

func main(){
  var arr [10]int
  arr[0] = 10
  arr[1] = 11

  fmt.Println("first element is %d\n", arr[0])
  fmt.Println("second element is %d\n", arr[1])
  fmt.Println("third element is %d\n", arr[2])

  arr2 := [...]int{4, 5, 6}
  fmt.Println("first element is %d\n", arr2[0])
  fmt.Println("second element is %d\n", arr2[1])
  fmt.Println("third element is %d\n", arr2[2])
  //fmt.Println("third element is %d\n", arr[3]) 加這行會error

  doubleArray := [2][4]int{[4]int{1,2,3,4}, [4]int{5,6,7,8}}
  fmt.Println("first element is %d\n", doubleArray[0])
}