//指針運用 practice
//method繼承

package main

import "fmt"

type Human struct {
  name string
  gender string
}

//繼承練習用
type tester struct {
  Human //繼承Human的method
  item string
}

//這個function會失敗 因為傳進去的p 是一個copy 本體沒修改到
func (p Human) changeGender1(c string) Human {
  p.gender = c
  return p
}

//這個才會成功
func (p *Human) changeGender2(c string) *Human { //注意輸入輸出要一致
  p.gender = c
  return p
}

func main() {
  //指針練習
  tom := Human{name: "Tom",gender: "male" }
  tom.changeGender1("female")
  fmt.Println(tom)

  ben := Human{name: "Ben",gender: "male" }
  ben.changeGender2("female")
  fmt.Println(ben)

  //繼承練習
  mike := tester{Human:Human{name: "Mike", gender: "male"}, item: "test"} //嘗試拿掉item: 會error 看來宣告方式要一致
  mike.changeGender2("female") //可以適用繼承來的method
  fmt.Println(mike) 




}