// 把func定義為一種type
package main

import "fmt"



//把出入輸出相同的func 定義為一種type
type move func(string)string

func sing(name string)string{
  return "sing " + name
}
func beat(name string)string{
  return " beat " + name
}
func eat(name string)string{
  return " eat " + name
}

func say(name string, do move) string {
  return "I want to "+ do(name)
}

//把func當作value傳入
func main(){
  fmt.Println(say("Rolling in the deep", sing))
  fmt.Println(say("stone", beat))
  fmt.Println(say("steak", eat ))
} 