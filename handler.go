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
	urlPath := path.Clean(r.URL.Path)
	localPath := path.Join(dir, urlPath)
	switch r.Method {
	case "POST":
		if noupload {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if isProtected("upload", auth) {
			u, p, ok := r.BasicAuth()
			if !ok || user != u || password != p {
				w.Header().Set("WWW-Authenticate", "Basic")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}
		if h := r.Header.Get("Content-Type"); !strings.HasPrefix(h, "multipart/form-data") {
			folder := path.Clean(r.PostFormValue("folder"))
			if folder != "" && folder != "/" {
				switch err := os.Mkdir(path.Join(localPath, folder), 0750).(type) {
				case *os.PathError:
					http.Error(w, err.Op+" "+path.Join(urlPath, folder)+": "+
						err.Err.Error(), 500)
					return
				case error:
					http.Error(w, err.Error(), 500)
					return
				}
			}
			http.Redirect(w, r, path.Join(urlPath, folder), 302)
			return
		}
		body, err := r.MultipartReader()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		for part, err := body.NextPart(); err == nil; part, err = body.NextPart() {
			formName := part.FormName()
			if formName != "file" {
				log.Printf("Skipping '%s'", formName)
				continue
			}
			log.Printf("Handling '%s'", formName)
			destFile, err := os.Create(path.Join(localPath, part.FileName()))
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			defer destFile.Close()
			if _, err := io.Copy(destFile, part); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
		http.Redirect(w, r, urlPath, 302)
	case "GET":
		if isProtected("index", auth) {
			u, p, ok := r.BasicAuth()
			if !ok || user != u || password != p {
				w.Header().Set("WWW-Authenticate", "Basic")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}
		entryInfo, err := os.Stat(localPath)
		if err != nil {
			log.Printf("ERROR: os.Stat('%s')", localPath)
			http.Error(w, err.Error(), 500)
			return
		}
		if entryInfo.IsDir() && !strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, r.URL.Path+"/", 302)
			return
		}
		if entryInfo.IsDir() && index != "" {
			if fi, err := os.Stat(path.Join(localPath, index)); err == nil {
				if !fi.IsDir() {
					http.Redirect(w, r, path.Join(urlPath, index), 302)
					return
				}
			}
		}
		if entryInfo.IsDir() {
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
			entries, err := readDir(localPath, sortKey, sortOrder)
			if err != nil {
				log.Printf("ERROR: ReadDir('%s')", localPath)
				http.Error(w, err.Error(), 500)
				return
			}
			ctx := context{urlPath == "/", !noupload, entries, sortOrderMap}
			if err := tmpl.Execute(w, ctx); err != nil {
				log.Println("ERROR: Executing template")
				http.Error(w, err.Error(), 500)
				return
			}
		} else {
			if isProtected("download", auth) {
				u, p, ok := r.BasicAuth()
				if !ok || user != u || password != p {
					w.Header().Set("WWW-Authenticate", "Basic")
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
			}
			f, err := os.Open(localPath)
			if err != nil {
				log.Printf("ERROR: os.Open('%s')", localPath)
				http.Error(w, err.Error(), 500)
				return
			}
			defer f.Close()
			log.Printf("Serving '%s'", localPath)
			http.ServeContent(w, r, entryInfo.Name(), entryInfo.ModTime(), f)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	return
}
