package main

import(
	".."
	"fmt"
)

func main(){

	hashmap,_ := murmur3map.NewHashMap(100)
	hashmap.Set("a","valsfsdfs")
	val, _ := hashmap.Get("a")
	fmt.Println(val.Value)
}