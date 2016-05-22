package daemon

import (
	"crypto/tls"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"blackbox/cache"
	"blackbox/cli"
	"blackbox/constants"
	"blackbox/worker"
)

func acceptConn(ln net.Listener, cc chan net.Conn) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		//log.Println("New connection accepted:", conn)
		cc <- conn
	}
}

func flushSignal(interval time.Duration) chan bool {
	fc := make(chan bool)
	go func() {
		for {
			time.Sleep(interval)
			fc <- true
		}
	}()
	return fc
}

type Server struct {
	workerNum int
	listener  net.Listener
	cache     *cache.Cache
}

func handleInterrupt(s *Server) {
	intc := make(chan os.Signal)
	// catch interrupt signal
	signal.Notify(intc, os.Interrupt)
	go func() {
		<-intc
		s.listener.Close()
		err := s.cache.Close()
		if err != nil {
			log.Print(err)
		}
		log.Print("Interrupt shutdown complete.")
		os.Exit(2)
	}()
}

func NewServer() *Server {
	s := Server{
		workerNum: constants.WORKER_NUM,
		listener:  nil,
		cache:     nil,
	}
	return &s
}

func (s *Server) Start(args cli.Args) error {

	// init database cache
	cache, err := cache.NewCache(args.DbPath)
	if err != nil {
		log.Println(err)
		return err
	}

	// init server certificate
	cer, err := tls.LoadX509KeyPair(args.CertFile, args.KeyFile)
	if err != nil {
		log.Println(err)
		return err
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	port := strconv.Itoa(args.Port)
	ln, err := tls.Listen("tcp", ":"+port, config)
	if err != nil {
		log.Println(err)
		return err
	}

	// server start done
	s.cache = cache
	s.listener = ln
	go handleInterrupt(s)
	log.Print("Server started, listening at port: ", port)
	s.Ready()
	return nil
}

func (s *Server) Ready() error {
	defer func() {
		s.listener.Close()
		s.cache.Close()
		log.Print("Shutdown complete.")
	}()

	results := make(chan worker.Pair)
	pending := make(chan net.Conn, s.workerNum)
	flush := flushSignal(constants.CACHE_FLUSH_INT * time.Second)

	for i := 0; i < s.workerNum; i++ {
		w := worker.NewWorker(pending, results, s.cache)
		go w.Loop()
	}

	go acceptConn(s.listener, pending)
	for {
		select {
		case pair := <-results:
			err := s.cache.Set([]byte(pair.Key), []byte(pair.Value))
			if err != nil {
				log.Println(err)
			}
		case <-flush:
			log.Print("Flushing db to:", s.cache.StorePath)
			s.cache.Flush()
		}
	}
}
