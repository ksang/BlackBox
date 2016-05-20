package worker

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCacheUpdate(t *testing.T) {
	var c ReadCache
	c.New()
	for i := 0; i < 10; i++ {
		key := make([]byte, 16)
		_, err := rand.Read(key)
		if err != nil {
			t.Errorf("error:", err)
		}
		c.Update(hex.EncodeToString(key), "Password String")
	}
	c.PrintPair()
	fmt.Println("Size:", c.Size)
}
