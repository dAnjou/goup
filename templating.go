package main

import (
	"github.com/dustin/go-humanize"
	"html/template"
	"os"
)

type context struct {
	IsRoot     bool
	Upload     bool
	DirEntries []os.FileInfo
	SortOrder  map[string]string
}

var (
	asset__index_html, err = Asset("assets/index.html")
	default_template = string(asset__index_html)
	funcMap = template.FuncMap{
		"size": func(b int64) string {
			return humanize.Bytes(uint64(b))
		},
		//"time": humanize.Time,
	}
	tmpl = template.Must(template.New("index").Funcs(funcMap).Parse(default_template))
)
