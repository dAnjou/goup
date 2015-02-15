package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	url_path := path.Clean(r.URL.Path)
	local_path := path.Join(dir, url_path)
	switch r.Method {
	case "POST":
		if noupload {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if h := r.Header.Get("Content-Type"); !strings.HasPrefix(h, "multipart/form-data") {
			folder := path.Clean(r.PostFormValue("folder"))
			if folder != "" && folder != "/" {
				switch err := os.Mkdir(path.Join(local_path, folder), 0750).(type) {
				case *os.PathError:
					http.Error(w, err.Op+" "+path.Join(url_path, folder)+": "+
						err.Err.Error(), 500)
					return
				case error:
					http.Error(w, err.Error(), 500)
					return
				}
			}
			http.Redirect(w, r, path.Join(url_path, folder), 302)
			return
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
		if entry_info.IsDir() && index != "" {
			if fi, err := os.Stat(path.Join(local_path, index)); err == nil {
				if !fi.IsDir() {
					http.Redirect(w, r, path.Join(url_path, index), 302)
					return
				}
			}
		}
		if entry_info.IsDir() {
			sortKey := r.URL.Query().Get("key")
			sortOrder := asc
			switch r.URL.Query().Get("order") {
			case "asc":
				sortOrder = asc
				sortOrderMap[sortKey] = "desc"
			case "desc":
				sortOrder = desc
				sortOrderMap[sortKey] = "asc"
			}
			entries, err := readDir(local_path, sortKey, sortOrder)
			if err != nil {
				log.Printf("ERROR: ReadDir('%s')", local_path)
				http.Error(w, err.Error(), 500)
				return
			}
			ctx := context{url_path == "/", !noupload, entries, sortOrderMap}
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
