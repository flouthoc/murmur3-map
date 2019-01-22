package main

import(
	".."
	"fmt"
)

type sampletype struct{

	name string
}

func main(){


	/** Example one **/
	hashmap,_ := murmur3map.NewMap(100)
	hashmap.Set("keyforstring","Sample Text Value")
	val, _ := hashmap.Get("keyforstring")
	fmt.Println(val.Value)

	/** Example Two **/
	sampletype_instance := new(sampletype)
	sampletype_instance.name = "my name is flouthoc"

	hashmap.Set("keyforobject", sampletype_instance)

	valtwo,_ := hashmap.Get("keyforobject")
	obj2,_ := valtwo.Value.(*sampletype)
	fmt.Println(obj2.name)
}