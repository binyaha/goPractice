//struct 繼承
package main

import "fmt"


type human struct {
  name string
  age int
  phone string
}

type hero struct {
  human
  name string //注意 有重複的屬性
  phone string //注意 有重複的屬性
  skill string

}

func main(){
  hero1 := hero{human:human{"Tom",24,"0911123456"},name: "superman",phone: "123123123" ,skill: "fly"}
  fmt.Println("hero1 can", hero1.skill)
  fmt.Println("hero1's name is", hero1.name)
  fmt.Println("hero1's real name is", hero1.human.name)
}