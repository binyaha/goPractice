//variable practice
package main

import (
"fmt"
)

func main() {
  var a int
  var b, c int = 1,2
  var d,e = 1,2
  f,g := 1,2
  i := "ddc"
  j := []byte(i)
  j[0] = 'c'
  k := string(j)
  l := i+k
  m := "a"+i[1:]
  err := errors.New("emit macho dwarf: elf header corrupted")
    if err != nil {
    fmt.Print(err)
  }
  fmt.Println(a,b,c,d,e,f,g,i,j,k,l,m,err)
}