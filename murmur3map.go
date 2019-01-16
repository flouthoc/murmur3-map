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

func (h *HashMap) Get(key string)(*Node, bool){
	index := h.getIndex(key)
	//get bucket for the index
	chain := h.buckets[index]
	for _, node := range chain{
		if node.key == key{
			return &node, true
		}
	}

	return nil, false
}

func (h* HashMap) Set(key string, value interface{}) bool {
	index := h.getIndex(key)
	chain := h.buckets[index]
	found := false

	for i := range chain {
		node := &chain[i]
		if( node.key == key{
			node.Value = value
			found = true
		}
	}

	if found {
		return true
	}

	//key is not thee
	if h.size == h.count {
		return false;
	}

	node := Node{key:key, Value: value}
	chain = append(chain, node)
	h.buckets[index] = chain
	h.count++;
	return true;
}


func (h *HashMap) Delete(key string) (*Node, bool){

	index := h.getIndex(key)
	chain := h.buckets[index]

	found := false
	var location int
	var mapNode *Node

	for loc, node := range chain {
		if node.key == key {
			found = true
			location = loc
			mapNode = &node
		}
	}

	if found {
		h.count--
		N := len(chain) - 1
		chain[location], chain[N] = chain[N], chain[location]
		chain = chain[:N]
		h.buckets[index] = chain
		return mapNode, true
	}

	return nil, false

}
