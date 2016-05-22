package cache

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cache, err := NewCache("/tmp")
	if err != nil {
		fmt.Println(err)
		return
	}
	cache.Set([]byte("11111"), []byte("1"))
	cache.Set([]byte("22222"), []byte("2"))
	cache.Set([]byte("11111"), []byte("11"))

	v, err := cache.Get([]byte("11111"))
	fmt.Println("Get 11111:", string(v))
	err = cache.Flush()
	if err != nil {
		fmt.Println(err)
	}
	cache.Close()
	cache2, err := NewCache("/tmp")
	if err != nil {
		fmt.Println(err)
		return
	}
	v2, err := cache2.Get([]byte("11111"))
	fmt.Println("Get 11111:", string(v2))
}
