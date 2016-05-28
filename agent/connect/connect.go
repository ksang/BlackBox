/* 
Package connect implements connection and protocol
functions to blackbox servrer.
*/
package connect

import (
	"blackbox/agent/cli"
	"blackbox/constants"
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net"
)

// Connects to blackbox server, using client certificate authentication.
func Connect(args cli.Args) (net.Conn, error) {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(args.CaCertData)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{args.AgentCert},
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

// Get a connection to blackbox server then request secret.
func RequestSecretConn(conn net.Conn, key string) ([]byte, error) {
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

// Connects to blackbox server and get encryption key, finally close the connection.
func RequestSecret(arg cli.Args, key string) ([]byte, error) {
	conn, err := Connect(arg)
	if err != nil{
		log.Print("Connection failed.")
		return nil, err
	}
	defer conn.Close()
	return RequestSecretConn(conn, key)
}
