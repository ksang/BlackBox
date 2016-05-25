package connect

import (
	"blackbox/agent/cli"
	"blackbox/constants"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func Connect(args cli.Args) (net.Conn, error) {

	//init ca certificate
	caCert, err := ioutil.ReadFile(args.CaCert)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// init client certificate
	cer, err := tls.LoadX509KeyPair(args.CertFile, args.KeyFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cer},
		RootCAs: caCertPool,
	}
	log.Print("Connecting to: ", args.Target)
	conn, err := tls.Dial("tcp", args.Target, tlsConfig)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func RequestSecret(arg cli.Args, key string) ([]byte, error) {
	conn, err := Connect(arg)
	if err != nil{
		log.Print("Connection failed.")
		return nil, err
	}
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
