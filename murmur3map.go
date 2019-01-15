package murmur3map

import "errors"

type Node{
	key string
	Value interface{}
}


type HashMap struct{
	size int
	count int
	buckets [][]Node
}


func (h* HashMap) getIntex(key string) int{
	return int(hash(key)) % h.size
}

func hash(key string) uint32{

	//call murmurhash
}


func NewHashMap(size int) (*HashMap, error){
	h := new(HashMap)
	if size < 1 {
		return h, errors.New("Size cannot be less than 0")
	}
	h.size = size
	h.count = 0
	h.buckets = make([][]Node, size)
	for i := range h.buckets{
		h.buckets[i] = make([]Node, 0)
	}
	return h, nil
}