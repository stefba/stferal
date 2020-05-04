package entry

import (
	"time"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
)

type Entry interface{
	Parent() Entry
	File()   *file.File

	Id()      int64
	Type()    string
	Section() string

	Hash()      string
	Timestamp() string

	Info()  info.Info
	Date()  time.Time

	Title(string) string
	Perma(string) string
}

type Entries []Entry
