package worker

import (
	"blackbox/cache"
	"log"
	"net"
)

type Pair struct {
	Key   string
	Value string
}

type Worker struct {
	stopping chan chan error
	pending  chan net.Conn
	result   chan Pair
	cache    *cache.Cache
}

func NewWorker(p chan net.Conn, r chan Pair, c *cache.Cache) *Worker {
	return &Worker{
		stopping: make(chan chan error),
		pending:  p,
		result:   r,
		cache:    c,
	}
}

func (w *Worker) Stop() error {
	errc := make(chan error)
	w.stopping <- errc
	return <-errc
}

func (w *Worker) Loop() {
	var err error
	for {
		select {
		case conn := <-w.pending:
			err = doWork(conn, w.cache, w.result)
			if err != nil {
				log.Println(err)
			}
		case errc := <-w.stopping:
			errc <- err
			return
		}
	}
}

func doWork(conn net.Conn, cache *cache.Cache, result chan Pair) error {
	defer conn.Close()
	key, err := ParseRequest(conn)
	if err != nil {
		log.Println(err)
		return err
	}

	value, err := cache.Get([]byte(key))
	if value != nil {
		log.Print("Old Pair, Key: ", key, " Value: ", string(value))
		n, err := conn.Write([]byte(value))
		if err != nil {
			log.Println(n, err)
			return err
		}
	} else {
		value, err := GenerateValue(16)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Print("New Pair, Key: ", key, " Value: ", value)
		result <- Pair{key, value}
		n, err := conn.Write([]byte(value))
		if err != nil {
			log.Println(n, err)
			return err
		}
	}
	return nil
}
