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
	assetIndexHTML, err = Asset("assets/index.html")
	defaultTemplate     = string(assetIndexHTML)
	funcMap             = template.FuncMap{
		"size": func(b int64) string {
			return humanize.Bytes(uint64(b))
		},
		//"time": humanize.Time,
	}
	tmpl = template.Must(template.New("index").Funcs(funcMap).Parse(defaultTemplate))
)
