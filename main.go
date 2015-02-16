package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/fcgi"
	"os"
)

var (
	noupload bool   = false
	user     string = ""
	password string = ""
	auth     string = ""
	dir      string = "."
	mode     string = "http"
	index    string = ""
	VERSION  string = ""
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nEnvironment variables (get overridden by command line arguments):")
		fmt.Fprintln(os.Stderr, "  GOUP_UPLOAD=false: disable uploads")
		fmt.Fprintln(os.Stderr, "  GOUP_DIR=<path>: see -dir")
		fmt.Fprintln(os.Stderr, "  GOUP_MODE=(http|fcgi): see -mode")
		fmt.Fprintln(os.Stderr, "  GOUP_INDEX=<filename>: see -index")
	}
	if os.Getenv("GOUP_UPLOAD") == "false" {
		noupload = true
	}
	if d := os.Getenv("GOUP_DIR"); d != "" {
		dir = d
	}
	if m := os.Getenv("GOUP_MODE"); m != "" {
		mode = m
	}
	if i := os.Getenv("GOUP_INDEX"); i != "" {
		index = i
	}
	address := flag.String("addr", "0.0.0.0:4000", "listen on this address")
	flag.BoolVar(&noupload, "noupload", noupload, "enable or disable uploads")
	flag.StringVar(&user, "user", user, "user for HTTP Basic authentication (-auth needs to be set)")
	flag.StringVar(&password, "password", password, "password for HTTP Basic authentication (-auth needs to be set)")
	flag.StringVar(&auth, "auth", auth, "comma-separated list of what will be protected by HTTP Basic authentication (index,download,upload)")
	flag.StringVar(&dir, "dir", dir, "directory for storing and serving files")
	flag.StringVar(&mode, "mode", mode, "run either standalone (http) or as FCGI application (fcgi)")
	flag.StringVar(&index, "index", index, "serve this file if it exists in the current directory instead of a listing")
	verbose := flag.Bool("v", false, "verbose output (no output at all by default)")
	version := flag.Bool("version", false, "show version and exit")
	flag.Parse()

	if *version {
		fmt.Println(VERSION)
		return
	}

	log.SetOutput(ioutil.Discard)
	if *verbose {
		log.SetOutput(os.Stdout)
		flag.VisitAll(func(f *flag.Flag) {
			log.Printf("SETTINGS: %s = %s", f.Name, f.Value)
		})
	}

	http.HandleFunc("/", handler)

	switch mode {
	case "http":
		log.Fatal(http.ListenAndServe(*address, nil))
	case "fcgi":
		log.Fatal(fcgi.Serve(nil, nil))
	default:
		log.Fatalf("Unknown mode '%s'!", mode)
	}
}
