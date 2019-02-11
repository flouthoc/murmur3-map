package murmur3map

import ("testing")

	/*var m map[string]int
	m = make(map[string]int)
	for n:=0; n<b.N; n++{
		m["keyfsdfsfsdfsdfds"] = 1
	}*/

func BenchmarkMurmur(b *testing.B){




	//this will take 16 ns/op

	hashmap,_ := NewMap(1)
	for n:=0; n<b.N; n++{
		
		hashmap.hash("sdfsdfsd")
	}
}

func BenchmarkMurmur2(b *testing.B){

	//this will take 32 ns/op

	hashmap,_ := NewMap(1)
	for n:=0; n<b.N; n++{
		
		hashmap.getIndex("sdfsdfsd")
	}
}