package murmur3map

import ("testing")

func BenchmarkMurmur(b *testing.B){

	/*var m map[string]int
	m = make(map[string]int)
	for n:=0; n<b.N; n++{
		m["keyfsdfsfsdfsdfds"] = 1
	}*/



	hashmap,_ := NewMap(1)
	for n:=0; n<b.N; n++{
		
		hashmap.Set("keyfsdfsfs",1)
	}
}