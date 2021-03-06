package murmur3map

/*
	Uses: murmur implementation from github.com/spaolacci/murmur3
*/

import ("errors"
	"./pkg/xxhash"
	"./pkg"
	"hash"
)


type Node struct{
	key string
	Value interface{}
}


type MurmurMap struct{
	size int
	count int
	hasher hash.Hash32
	buckets [][]Node
}


func (h* MurmurMap) getIndex(key string) int{
	return int(xxhash.Sum64([]byte(key))) % h.size
}


func (h* MurmurMap) hash(key string) uint64{

	//call murmurhash
	//----h.hasher.Write([]byte(key))
	//---return h.hasher.Sum32()

	return xxhash.Sum64([]byte(key))
	//return memhash32([])

	//return hasher.Sum([]byte(key))

}


func NewMap(size int) (*MurmurMap, error){
	h := new(MurmurMap)
	h.hasher = murmur3.New32WithSeed(0x01)
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

func (h *MurmurMap) Get(key string)(*Node, bool){
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

func (h* MurmurMap) Set(key string, value interface{}) bool {
	index := h.getIndex(key)
	chain := h.buckets[index]
	found := false

	/* remove sequential*/

	for i := range chain {
		node := &chain[i]
		if node.key == key {
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


func (h *MurmurMap) Delete(key string) (*Node, bool){

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
