package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
)

type sortable struct {
	Infos   *[]os.FileInfo
	SortBy  string
	Reverse bool
}

func xnor(a, b bool) bool { return !((a || b) && (!a || !b)) }

func (s sortable) Len() int { return len(*s.Infos) }
func (s sortable) Less(i, j int) bool {
	switch s.SortBy {
	case "mode":
		return xnor((*s.Infos)[i].Mode() > (*s.Infos)[j].Mode(), s.Reverse)
	case "time":
		return xnor((*s.Infos)[i].ModTime().After((*s.Infos)[j].ModTime()), s.Reverse)
	case "size":
		return xnor((*s.Infos)[i].Size() > (*s.Infos)[j].Size(), s.Reverse)
	default:
		return xnor((*s.Infos)[i].Name() > (*s.Infos)[j].Name(), s.Reverse)
	}
}
func (s sortable) Swap(i, j int) { (*s.Infos)[i], (*s.Infos)[j] = (*s.Infos)[j], (*s.Infos)[i] }

func readDir(dirname string, sortby string, reverse bool) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Sort(sortable{&list, sortby, reverse})
	return list, nil
}

var (
	verbose  = flag.Bool("v", false, "verbose output (no output at all by default)")
	noupload = flag.Bool("noupload", false, "disable uploads")
	address  = flag.String("addr", "0.0.0.0:4000", "listen on this address")
	dir      = flag.String("dir", ".", "directory for storing and serving files")
)

func index(w http.ResponseWriter, r *http.Request) {
	url_path := path.Clean(r.URL.Path)
	local_path := path.Join(*dir, url_path)
	switch r.Method {
	case "POST":
		if *noupload {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		body, err := r.MultipartReader()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		for part, err := body.NextPart(); err == nil; part, err = body.NextPart() {
			form_name := part.FormName()
			if form_name != "file" {
				log.Printf("Skipping '%s'", form_name)
				continue
			}
			log.Printf("Handling '%s'", form_name)
			dest_file, err := os.Create(path.Join(local_path, part.FileName()))
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			defer dest_file.Close()
			if _, err := io.Copy(dest_file, part); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
		http.Redirect(w, r, url_path, 302)
	case "GET":
		entry_info, err := os.Stat(local_path)
		if err != nil {
			log.Printf("ERROR: os.Stat('%s')", local_path)
			http.Error(w, err.Error(), 500)
			return
		}
		if entry_info.IsDir() && !strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, r.URL.Path+"/", 302)
			return
		}
		if entry_info.IsDir() {
			reverse := false
			switch r.URL.Query().Get("by") {
			case "asc":
				reverse = false
			case "desc":
				reverse = true
			}
			entries, err := readDir(local_path, r.URL.Query().Get("sort"), reverse)
			if err != nil {
				log.Printf("ERROR: ReadDir('%s')", local_path)
				http.Error(w, err.Error(), 500)
				return
			}
			ctx := context{url_path == "/", !*noupload, entries}
			if err := tmpl.Execute(w, ctx); err != nil {
				log.Println("ERROR: Executing template")
				http.Error(w, err.Error(), 500)
				return
			}
		} else {
			f, err := os.Open(local_path)
			if err != nil {
				log.Printf("ERROR: os.Open('%s')", local_path)
				http.Error(w, err.Error(), 500)
				return
			}
			defer f.Close()
			log.Printf("Serving '%s'", local_path)
			http.ServeContent(w, r, entry_info.Name(), entry_info.ModTime(), f)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	return
}

func main() {
	flag.Parse()

	log.SetOutput(ioutil.Discard)
	if *verbose {
		log.SetOutput(os.Stdout)
	}

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(*address, nil))
}
