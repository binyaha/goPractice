//map practice
package main

import "fmt"

func main() {

diction := make(map[string]int)

diction["one"]=1
diction["two"]=2
diction["three"]=3


fmt.Println(diction["one"])

rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
csharpRating, ok := rating["C#"]
if ok {
    fmt.Println("C# is in the map and its rating is ", csharpRating)
} else {
    fmt.Println("We have no rating associated with C# in the map")
}

hash1 := make(map[string]string)
hash1["one"]="i am one"
hash2 := hash1
hash2["one"] = "hash2 change key one value"

fmt.Println(hash2["one"], hash1["one"])

}
