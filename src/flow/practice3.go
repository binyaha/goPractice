//defer practice
//解釋: func結束後 開始執行defer 最後的先執行  
package main

import "fmt"

//這裡是解釋用的程式碼

func main(){
  defer fmt.Println("defer1")
  fmt.Println("2")
  fmt.Println("3")
  defer fmt.Println("defer2")
  fmt.Println("4")
  fmt.Println("5")
  defer fmt.Println("defer3") //注意 defer中 這個最先print
}