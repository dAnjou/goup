package main

import (
	"os"
	"sort"
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
	return xnor((*s.Infos)[i].Name() > (*s.Infos)[j].Name(), s.Reverse)
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
