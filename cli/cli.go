// Package cli implements command line arguments.
// It will also raise errors when user provide
// incorrect arguments.
package cli

import (
	"flag"
)

type Args struct {
	Daemon   bool
	Port     int
	CertFile string
	KeyFile  string
	CaCert   string
	DbPath   string
}

// Parse arguments from command line and raise error if necessary.
func Parse() Args {
	var daemon = flag.Bool("d", false, "Daemon mode.")
	var port = flag.Int("p", 23333, "Listen port.")
	var certFile = flag.String("c", "server.pem", "Server cert file.")
	var keyFile = flag.String("k", "server.key", "Server key file.")
	var caCert = flag.String("a", "ca.pem", "CA cert file.")
	var dbPath = flag.String("f", "/tmp", "Database cache location.")
	flag.Parse()
	return Args{
		Daemon:   *daemon,
		Port:     *port,
		CertFile: *certFile,
		KeyFile:  *keyFile,
		DbPath:   *dbPath,
		CaCert:   *caCert,
	}
}
