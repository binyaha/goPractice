//pass index or value
package main

import "fmt"

func add1(a int)int {
  a = a+1
  return a
}

//要改變傳入的值
func add2(a *int)int {
    *a = *a+1 
    return *a 
}

func main(){
  //注意x並沒有被改變
  x := 3
  fmt.Println("x = ",x)
  x1 := add1(x)
  fmt.Println("x+1 = ",x1)
  fmt.Println("x = ", x)


  //要改變y
  y := 3
  fmt.Println("y = ",y)
  y1 := add2(&y) //注意這邊
  fmt.Println("y+1 = ",y1)
  fmt.Println("y = ", y)
}