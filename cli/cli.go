package cli

import (
	"flag"
)

type Args struct {
	Daemon   bool
	Port     int
	CertFile string
	KeyFile  string
}

func Parse() Args {
	var daemon = flag.Bool("d", false, "Daemon mode.")
	var port = flag.Int("p", 23333, "Listen port.")
	var certFile = flag.String("c", "server.pem", "Server cert file.")
	var keyFile = flag.String("k", "server.key", "Server key file.")
	flag.Parse()
	return Args{
		Daemon:   *daemon,
		Port:     *port,
		CertFile: *certFile,
		KeyFile:  *keyFile,
	}
}
