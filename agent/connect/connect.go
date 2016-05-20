package connect

import (
	"blackbox/agent/cli"
	"blackbox/constants"
	"crypto/tls"
	"io"
	"log"
	"net"
)

func Connect(arg cli.Args) (net.Conn, error) {

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	log.Print("Connecting to: ", arg.Target)
	conn, err := tls.Dial("tcp", arg.Target, conf)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func RequestSecret(arg cli.Args, key string) ([]byte, error) {
	conn, err := Connect(arg)
	defer conn.Close()

	req := constants.KEY_REQUEST_HEADER + " " + key + "\n"
	log.Print("Sending request: ", req)
	n, err := conn.Write([]byte(req))
	if err != nil {
		log.Println(n, err)
		return nil, err
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	// If read without error or session closed.
	if err != nil && err != io.EOF {
		log.Println(n, err)
		return nil, err
	}
	return buf[:n], nil
}
