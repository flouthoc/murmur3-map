# mumur3-map

Go Fast Object Map based on upon Murmur3


## Usage

Example 1

```go
	hashmap,_ := murmur3map.NewMap(100)
	hashmap.Set("keyforstring","Sample Text Value")
	val, _ := hashmap.Get("keyforstring")
	fmt.Println(val.Value)
```

Example 2

```go
type sampletype struct{

	name string
}

func main(){

	hashmap,_ := murmur3map.NewMap(100)

	sampletype_instance := new(sampletype)
	sampletype_instance.name = "my name is flouthoc"

	hashmap.Set("keyforobject", sampletype_instance)

	valtwo,_ := hashmap.Get("keyforobject")
	obj2,_ := valtwo.Value.(*sampletype)
	fmt.Println(obj2.name)
}

```


## Docs

#### func NewMap(size int) (*MurmurMap, error)
Allocates a new mumurmap with  given size. Returns struct MurmurMap

#### func (h* MurmurMap) Set(key string, value interface{}) bool

Sets k-v pair in map. Keys are supposed to be strings. Values can be anything just be careful when you are fetching them back from the map , see example for usage.

#### func (h *MurmurMap) Get(key string)(*Node, bool)

Returns back the value corresponding to the specified key otherwise returns false.
Just be sure to cast value back to struct type when you are done with fetching. See Example for usage.



### References
* https://en.wikipedia.org/wiki/MurmurHash
* https://github.com/spaolacci/murmur3

### Feel free to fork or create pull request

