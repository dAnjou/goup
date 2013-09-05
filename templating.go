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
}

var (
	default_template = `
<!DOCTYPE html>
<title>goup</title>
<meta charset='utf-8'>
<style type="text/css">
table {
	font-family: monospace;
}
td {
	text-align: right;
	padding: 0 15px;
}
td.name {
	text-align: left;
}
.grosskursiv { font-style:italic; font-size:200%; }
</style>
{{ if .Upload }}
<div>
	<form action='.' method='post' enctype='multipart/form-data'>
		<input type='file' name='file'>
		<input type='submit'>
	</form>
	<form action='.' method='post'>
		<input type='text' name='folder' placeholder='folder'>
		<input type='submit'>
	</form>
</div>
<hr>
{{ end }}
<table>
	<tr>
		<th>mode</th>
		<th>last modified</th>
		<th>size</th>
		<th>name</th>
	</tr>
	{{ if not .IsRoot }}
	<tr>
		<td></td><td></td><td></td>
		<td class="name"><a href="..">..</a></td>
	</tr>
	{{ end }}
	{{ range $e := .DirEntries }}
	<tr>
		<td>{{ $e.Mode }}</td>
		<td>{{ $e.ModTime.Format "2006-01-02T15:04:05" }}</td>
		{{ if $e.IsDir }}
			<td></td>
			<td class="name"><a href="{{ $e.Name }}/">{{ $e.Name }}/</a></td>
		{{ else }}
			<td>{{ size $e.Size }}</td>
			<td class="name"><a href="{{ $e.Name }}">{{ $e.Name }}</a></td>
		{{ end }}
	</tr>
	{{ end }}
</table>
`
	funcMap = template.FuncMap{
		"size": func(b int64) string {
			return humanize.Bytes(uint64(b))
		},
		//"time": humanize.Time,
	}
	tmpl = template.Must(template.New("index").Funcs(funcMap).Parse(default_template))
)
