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
	dir      string = "."
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nEnvironment variables (get overridden by command line arguments):")
		fmt.Fprintln(os.Stderr, "  GOUP_UPLOAD=false: disable uploads")
		fmt.Fprintln(os.Stderr, "  GOUP_DIR=<path>: see -dir")
		fmt.Fprintln(os.Stderr, "  GOUP_MODE=(http|fcgi): see -mode")
	}
	if os.Getenv("GOUP_UPLOAD") == "false" {
		noupload = true
	}
	if d := os.Getenv("GOUP_DIR"); d != "" {
		dir = d
	}
	mode := "http"
	if m := os.Getenv("GOUP_MODE"); m != "" {
		mode = m
	}
	flag.StringVar(&dir, "dir", dir, "directory for storing and serving files")
	flag.StringVar(&mode, "mode", mode, "run either standalone (http) or as FCGI application (fcgi)")
	flag.BoolVar(&noupload, "noupload", noupload, "enable or disable uploads")
	verbose := flag.Bool("v", false, "verbose output (no output at all by default)")
	address := flag.String("addr", "0.0.0.0:4000", "listen on this address")
	flag.Parse()

	log.SetOutput(ioutil.Discard)
	if *verbose {
		log.SetOutput(os.Stdout)
		flag.VisitAll(func(f *flag.Flag) {
			log.Printf("SETTINGS: %s = %s", f.Name, f.Value)
		})
	}

	http.HandleFunc("/", index)

	switch mode {
	case "http":
		log.Fatal(http.ListenAndServe(*address, nil))
	case "fcgi":
		log.Fatal(fcgi.Serve(nil, nil))
	default:
		log.Fatalf("Unknown mode '%s'!", mode)
	}
}
