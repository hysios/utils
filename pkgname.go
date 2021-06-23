package utils

import (
	"errors"
	"strings"
)

type PkgName struct {
	Fullname string
	Path     string
	Domain   string
}

func ParsePkg(name string) (*PkgName, error) {
	var (
		domain, path string
		pos          = strings.IndexAny(name, "/")
	)

	if pos == 0 {
		return nil, errors.New("head or tail not \"/\"")
	} else if pos == len(name)-1 {
		return nil, errors.New("head or tail not \"/\"")
	}

	if pos > 0 { // has domain
		first := name[:pos]
		if strings.IndexAny(first, ".") > 0 {
			domain = first
			path = name[pos+1:]
		} else {
			path = name
		}
	} else {
		path = name
	}

	return &PkgName{
		Fullname: name,
		Domain:   domain,
		Path:     path,
	}, nil
}
