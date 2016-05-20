package worker

import (
	"log"
	"net"
)

type Worker struct {
	stopping chan chan error
	pending  chan net.Conn
	result   chan Pair
	rcache   *Cache
}

func NewWorker(p chan net.Conn, r chan Pair, rc *Cache) *Worker {
	w := &Worker{
		stopping: make(chan chan error),
		pending:  p,
		result:   r,
		rcache:   rc,
	}
	return w
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
			err = doWork(conn, w.rcache, w.result)
			if err != nil {
				log.Println(err)
			}
		case errc := <-w.stopping:
			errc <- err
			return
		}
	}
}

func doWork(conn net.Conn, rc *Cache, result chan Pair) error {
	defer conn.Close()
	key, err := ParseRequest(conn)
	if err != nil {
		log.Println(err)
		return err
	}

	if value, ok := rc.Get(key); ok {
		log.Print("Old Pair, Key: ", key, " Value: ", value)
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
