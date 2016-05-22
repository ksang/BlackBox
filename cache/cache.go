package cache

import (
	"blackbox/constants"
	"github.com/steveyen/gkvlite"
	"log"
	"os"
	"path/filepath"
)

type Cache struct {
	fd         *os.File
	StorePath  string
	Store      *gkvlite.Store
	Collection *gkvlite.Collection
}

func NewCache(dbPath string) (*Cache, error) {
	path := filepath.Join(dbPath, constants.CACHE_FILE)
	dbf, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	store, err := gkvlite.NewStore(dbf)
	if err != nil {
		return nil, err
	}
	collection := store.SetCollection(constants.COLLECTION_NAME, nil)
	log.Print("Database open: ", path)
	return &Cache{
		StorePath:  path,
		Store:      store,
		Collection: collection,
	}, nil
}

func (c *Cache) Set(key []byte, value []byte) error {
	return c.Collection.Set(key, value)
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	return c.Collection.Get(key)
}

func (c *Cache) Flush() error {
	return c.Store.Flush()
}

func (c *Cache) Delete(key []byte) (bool, error) {
	return c.Collection.Delete(key)
}

func (c *Cache) Close() error {
	err := c.Flush()
	if err != nil {
		return err
	}
	c.Store.Close()
	c.fd.Close()
	return nil
}
