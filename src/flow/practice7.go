//goto practice
package main
 
import "fmt"

func main(){
  
    stoneHealth := 50
  Health:
    stoneHealth -= 10
    fmt.Println("hit Stone")
    if  stoneHealth > 25 {
      goto Health
    } else if stoneHealth < 25 && stoneHealth > 0 {
      fmt.Println("Stone Hurt")
      goto Hurt
    } else{
      goto Dead
    }

  Hurt:
    stoneHealth -= 10
    if stoneHealth > 0 {
      fmt.Println("hit Stone")
      goto Hurt
    } else {
      goto Dead
    }


  Dead:
    fmt.Println("Stone dead")

}