# BlackBox
Encryption Key Management System

BlackBox enables you keep file encrytion key on the server, agent never keeps the key and only gets it's key on-demand.
System using client side certificate authentication, following extension is required:
"X509v3 Extended Key Usage: TLS Web Client Authentication, TLS Web Server Authentication"

NOTE:
    For the certificate signature hash algorithm, md5 is not supported, please use sha256/384/512 instead.

### BlackBox
	Server side program, generate and keep keys.
	  -a string
	    	CA cert file. (default "ca.pem")
	  -c string
	    	Server cert file. (default "server.pem")
	  -d	Daemon mode.
	  -f string
	    	Database cache location. (default "/tmp")
	  -k string
	    	Server key file. (default "server.key")
	  -p int
	    	Listen port. (default 23333)

### Agent
	Client side program, request key and encrypt/decrypt files.
	  -a string
	    	CA cert file. (default "ca.pem")
	  -c string
	    	Agent cert file. (default "agent.pem")
	  -d	Decrypt mode.
	  -e	Encrypt mode.
	  -f string
	    	Path to the file.
	  -k string
	    	Agent key file. (default "agent.key")
	  -p string
	    	Path to the folder.
	  -r	Remove the original file.(encrypt only)
	  -s string
	    	Encrypt file suffix. (default "blackbox")
	  -t string
	    	Target address with port. (default "127.0.0.1:23333")
	