# BlackBox
Encryption Key Management System

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
	