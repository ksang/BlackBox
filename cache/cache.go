/*
Package cache implements a interface to use gkvlite.
It simplifies the usage for blackbox server.
*/

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

// Initialize the cache with dbPath location specified in argument.
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

// Set cache item, insert for non-exist ones, update for exist ones.
func (c *Cache) Set(key []byte, value []byte) error {
	return c.Collection.Set(key, value)
}

// Get value using key.
func (c *Cache) Get(key []byte) ([]byte, error) {
	return c.Collection.Get(key)
}

// Save cache in memory to disk.
func (c *Cache) Flush() error {
	return c.Store.Flush()
}

// Delete item.
func (c *Cache) Delete(key []byte) (bool, error) {
	return c.Collection.Delete(key)
}

// Cleanup cache.
func (c *Cache) Close() error {
	err := c.Flush()
	if err != nil {
		return err
	}
	c.Store.Close()
	c.fd.Close()
	return nil
}
