//slice practice

package main

import "fmt"

func main() {
  slice := []byte{'a', 'b', 'c', 'd'}
  fmt.Println(slice)
  slice2 := slice[:3]
  fmt.Println(slice2)
  slice3 := slice[:3:4]
  fmt.Println(slice3)
}