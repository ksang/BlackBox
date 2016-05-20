package worker

import "fmt"

type Pair struct {
	Key   string
	Value string
}

type Caches interface {
	Update(string, string)
	Get(string) (string, bool)
}

type Cache struct {
	Pairs map[string]string
	Size  int
}

func NewCache() *Cache {
	c := &Cache{
		Pairs: make(map[string]string),
		Size:  0,
	}
	return c
}

func (c *Cache) Update(key string, value string) {
	c.Pairs[key] = value
	c.Size = len(c.Pairs)
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.Pairs[key]
	return value, ok
}

func (c *Cache) PrintPair() {
	for key, value := range c.Pairs {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println("Total Size:", c.Size)
}
