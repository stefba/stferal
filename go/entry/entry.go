package entry

import (
	"sacer/go/entry/file"
	"sacer/go/entry/info"
	"time"
)

type Entry interface {
	Parent() Entry
	File() *file.File

	Id() int64
	Type() string
	Section() string

	Hash() string
	Timestamp() string

	Info() info.Info
	Date() time.Time

	Title(string) string
	Path(string) string
	Perma(string) string

	MediaObject() bool 
	ObjectType() string

	SetParent(Entry)
	SetInfo(info.Info)
}

type Collection interface {
	Entries() Entries
}

type Blob interface {
	Location(string) (string, error)
}

func IsBlob(e Entry) bool {
	_, ok := e.(Blob)
	return ok
}
