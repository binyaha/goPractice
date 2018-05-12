// struct method  計算正多邊行面積
package main

import "fmt"
import "math"

type polygonal struct {
  number float64  //幾個邊
  length float64  //邊長
}

//計算多邊形的一個小三角形面積
func (t polygonal) subArea() float64 {
  subDegree := math.Pi/math.Floor(t.number)  //
  len1 := t.length
  len2 := 0.5*t.length/math.Tan(subDegree)
  return 0.5*len1*len2
}

//單純練習多寫一個有點多餘的method
func (a polygonal) area() float64 {

  if a.number != math.Floor(a.number) {
    fmt.Println("delete float number because number should be integer")
  }

  return a.subArea()*math.Floor(a.number)
}



func main(){

  object := polygonal{5,2}  //這裡設定多邊形面積和邊長
  if object.number > 3 {
  fmt.Println("Object's area is",object.area())
  } else {
    fmt.Println("number need bigger or equal 3")
  } 

}