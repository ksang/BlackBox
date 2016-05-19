package daemon

import (
	"log"
	"net"
	"strconv"
	"crypto/tls"

	"blackbox/cli"
	"blackbox/worker"
	"blackbox/constants"
)

func AcceptConn(ln net.Listener, cc chan net.Conn) {
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

type Server struct {
	workerNum int	
	listener net.Listener
}

func NewServer() *Server {
	s := Server {
					workerNum 	: constants.WORKER_NUM,
					listener  	: nil,
				}
	return &s
}

func (s *Server) Start(args cli.Args) error {
	log.SetFlags(log.Lshortfile)

	cer, err := tls.LoadX509KeyPair(args.CertFile, args.KeyFile)
	if err != nil {
		log.Println(err)
		return err
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":" + strconv.Itoa(args.Port), config) 
	if err != nil {
		log.Println(err)
		return err
	}
	s.listener = ln
	s.Ready()
	return nil
}

func (s *Server) Ready() error {
	defer s.listener.Close()
	rcache := worker.NewCache()
	wcache := worker.NewCache()

	results := make(chan worker.Pair)
	pending := make(chan net.Conn, s.workerNum)

	for i := 0; i < s.workerNum; i++ {
		w := worker.NewWorker(pending, results, rcache)
		go w.Loop()
	}

	go AcceptConn(s.listener, pending)
	for {	
		select {
		case pair := <- results:
			rcache.Update(pair.Key, pair.Value)
			wcache.Update(pair.Key, pair.Value)
		}		
	}	
}
