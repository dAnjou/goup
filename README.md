Goup
====

Goup is a small one-binary file server written in Go which also allows to upload
files. You can run it standalone or deploy it as FastCGI application.

## Installation for Go users

`go get -u github.com/dAnjou/goup`

## Binary and packages

https://github.com/dAnjou/goup/releases

## Usage

	Usage of goup:
	  -addr="0.0.0.0:4000": listen on this address
	  -auth="": comma-separated list of what will be protected by HTTP Basic authentication (index,download,upload)
	  -dir=".": directory for storing and serving files
	  -index="": serve this file if it exists in the current directory instead of a listing
	  -mode="http": run either standalone (http) or as FCGI application (fcgi)
	  -noupload=false: enable or disable uploads
	  -password="": password for HTTP Basic authentication (-auth needs to be set)
	  -user="": user for HTTP Basic authentication (-auth needs to be set)
	  -v=false: verbose output (no output at all by default)
	  -version=false: show version and exit
