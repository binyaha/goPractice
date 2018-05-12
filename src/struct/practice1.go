//struct practice (物件導向)
package main

import "fmt"

type human struct {
  name string
  age int

}

func compareAge(p1,p2 human)(r1 human,r2 human,y int){
  if p1.age > p2.age {
    return p1,p2, p1.age-p2.age
  } else {
    return p2,p1, p2.age-p1.age
  }
}



func main(){
  var hank human
  hank.name = "Hank"
  hank.age = 30
  stone := human{"Stone", 38}//不同的initial方法
  cherry :=human{age:18,name:"Cherry"}

  older1, younger1, age := compareAge(hank,stone) //物件就是一個參數
  older2, younger2, age := compareAge(hank,cherry)
  fmt.Printf("%s is %d years older than %s\n", older1.name, age, younger1.name)
  fmt.Printf("%s is %d years older than %s\n", older2.name, age, younger2.name)
}





