[![Go Report Card](https://goreportcard.com/badge/github.com/hatorikibble/dirserv)](https://goreportcard.com/report/github.com/hatorikibble/dirserv)
# dirserv
small webserver for serving directories written in go

     ./dirserv --help
    Usage of ./dirserv:
      --dir string
        	directory (default ".")
      --port int
        	port number (default 3000)


# Example

    $ ./dirserv --dir=/tmp/html_files --port=8080
    2017/12/04 22:03:31 Serving directory '/tmp/html_files' on port 8080
