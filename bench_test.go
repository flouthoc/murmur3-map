package murmur3map

import ("testing")

	/*var m map[string]int
	m = make(map[string]int)
	for n:=0; n<b.N; n++{
		m["keyfsdfsfsdfsdfds"] = 1
	}*/

func BenchmarkInbuitMapSet(b *testing.B){

	var m map[string]int
	m = make(map[string]int)
	for n:=0; n<b.N; n++{
		m["sdfsdfsd"] = 1
	}
}

func BenchmarkInbuitMapGet(b *testing.B){


	var m map[string]int
	m = make(map[string]int)

	m["sdfsdfsd"] = 1

	for n:=0; n<b.N; n++{
		if m["sdfsdfsd"] > 0 {

		}
	}
}

func BenchmarkMurmurOnlyHash(b *testing.B){

	//this will take 16 ns/op

	hashmap,_ := NewMap(1)
	for n:=0; n<b.N; n++{
		hashmap.hash("sdfsdfsd")
	}
}

func BenchmarkMurmurOnlyGetIndex(b *testing.B){

	//this will take 32 ns/op

	hashmap,_ := NewMap(1)
	for n:=0; n<b.N; n++{
		hashmap.getIndex("sdfsdfsd")
	}
}

func BenchmarkMurmurSet(b *testing.B){

	//this will take 32 ns/op

	hashmap,_ := NewMap(1)
	for n:=0; n<b.N; n++{
		hashmap.Set("sdfsdfsd", 1)
	}
}

func BenchmarkMurmurGet(b *testing.B){

	//this will take 32 ns/op

	hashmap,_ := NewMap(1)
	hashmap.Set("sdfsdfsd", 1)
	for n:=0; n<b.N; n++{
		hashmap.Get("sdfsdfsd")
	}
}