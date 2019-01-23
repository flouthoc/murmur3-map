package murmur3map

func BenchmarkMurmur(b *testing.B){

	for n:=0; n<b.N; n++{
		hashmap,_ := NewMap(100)
		hashmap.Set("keyforstring","Sample Text Value")
	}
}