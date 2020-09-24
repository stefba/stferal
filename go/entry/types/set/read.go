package set

import (
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/tools/read"
	"sacer/go/entry/tools/sort"
	"sacer/go/entry/types"
)

func readEntries(path string, parent entry.Entry) (entry.Entries, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "readEntries",
	}

	files, err := read.GetFiles(path, false)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	entries, err := readEntryFiles(files, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	return sort.SortEntries(path, entries)
}

func readEntryFiles(files []*read.FileInfo, parent entry.Entry) (entry.Entries, error) {
	entries := entry.Entries{}
	for _, fi := range files {
		entry, err := media.NewMediaEntry(fi.Path, parent)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

