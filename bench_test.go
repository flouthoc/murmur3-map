package murmur3map

import ("testing")

func BenchmarkMurmur(b *testing.B){

	
	for n:=0; n<b.N; n++{
		hashmap,_ := NewMap(1)
		hashmap.Set("key",1)
	}
}