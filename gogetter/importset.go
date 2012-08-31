package main

import "fmt"

type ImportSet struct {
	set []string
}

func NewImportSet() *ImportSet {
	return &ImportSet{make([]string, 0)}
}

func (is *ImportSet) Put(s string) {
	if !is.contains(s) {
		is.set = append(is.set, s)
	}
}

func (is ImportSet) contains(s string) bool {
	for _, st := range is.set {
		if st == s {
			return true
		}
	}
	return false
}

func (is ImportSet) String() string {
	return fmt.Sprint(is.set)
}
