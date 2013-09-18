Goup
====

Goup is a small one-binary file server written in Go which also allows to upload
files. You can run it standalone or deploy it as FastCGI application.

`go get github.com/dAnjou/goup`

## Binaries (and demo)

http://dump.danjou.de/goup/

## Usage

	Usage of goup:
	  -addr="0.0.0.0:4000": listen on this address
	  -dir=".": directory for storing and serving files
	  -mode="http": run either standalone (http) or as FCGI application (fcgi)
	  -noupload=false: enable or disable uploads
	  -v=false: verbose output (no output at all by default)

	Environment variables (get overridden by command line arguments):
	  GOUP_UPLOAD=false: disable uploads
	  GOUP_DIR=<path>: see -dir
	  GOUP_MODE=(http|fcgi): see -mode
