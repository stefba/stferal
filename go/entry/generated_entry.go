// Code generated by go generate; DO NOT EDIT.

package entry

import (
	"fmt"
	"stferal/go/entry/helper"
)

type Entry struct {
	Parent *Struct
	Object interface{}
}

func (e *Entry) O() interface{} {
	return e.Object
}

type Entries []*Entry

func NewEntry(path string, parent *Struct) (*Entry, error) {
	obj, err := NewEntryObject(path)
	if err != nil {
		return nil, err
	}
	return &Entry{
		Parent: parent,
		Object: obj,
	}, nil
}

func NewEntryObject(path string) (interface{}, error) {
	switch helper.FileType(path) {
	/*
		case "text":
			return text.New(path)
		case "dir":
			return set.New(path)
	*/
	}
	return nil, fmt.Errorf("invalid entry type: %v", path)
}
