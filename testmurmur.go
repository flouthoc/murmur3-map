package main

import (
	"./pkg"
	"fmt"
)


func main(){
	//fmt.Println(murmur3.New32())
	//fmt.Println(murmur3.New32().Sum32())
	hasher := murmur3.New32()
	hasher.Write([]byte("sssdfsd.."))
	fmt.Println(hasher.Sum32())
}


