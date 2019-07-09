package main

import (
	"log"
	"path/filepath"
	"plugin"
)

var v, f plugin.Symbol

func init() {
	dir := "./plugin"
	p, err := plugin.Open(filepath.Join(dir, "plugin.so"))
	if err != nil {
		log.Fatal(err)
	}

	v, err = p.Lookup("V")
	if err != nil {
		log.Fatal(err)
	}

	f, err = p.Lookup("F")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	*v.(*int) = 7
	f.(func())()
}
