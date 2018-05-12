//func practice
package main 

import "fmt"

//單純的func使用
func main() {
  a := "Hank"

  //Q1
  fmt.Println(addString(a))
  //Q2
  p1, p2 := addString2("hi","bye")
  fmt.Println(p1,"\n",p2)
  //Q3
  fmt.Println(addString3("Hank","30"))
}

func addString(b string)string{
   
  return b+b+" gogogo!"
}

//可多輸入輸出func

func addString2(str1 string, str2 string)(combine1 string, combine2 string){
  result1 := str1+" friend "+str2
  result2 := str2+" friend "+str1
  return result1,result2
}
//可不同類型input&int to string
func addString3(name string, age string)string{
  s := fmt.Sprint(age)
    result := name+ " is "+ s+" years old"
    return result
}
