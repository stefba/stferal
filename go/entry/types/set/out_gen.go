// Code generated by go generate; DO NOT EDIT.

package set

import (
	"fmt"

	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func (e *Set) Type() string {
	return "set"
}

func (e *Set) Parent() entry.Entry {
	return e.parent
}

func (e *Set) File() *file.File {
	return e.file
}

func (e *Set) Id() int64 {
	return e.date.Unix()
}

func (e *Set) Timestamp() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Set) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Set) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Set) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Set) Date() time.Time {
	return e.date
}

func (e *Set) Info() info.Info {
	return e.info
}

func (e *Set) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.Title(lang))
}

func (e *Set) IsBlob() bool {
	return entry.IsBlob(e)
}

func (e *Set) Entries() entry.Entries {
	return e.entries
}

// This recursive function call will be caught by a Tree type. For now, all
// further up parent entries are exclusively of type Tree.
func (e *Set) Section() string {
	return e.Parent().Section()
}

func (e *Set) Perma(lang string) string {
	return fmt.Sprintf("%v--not-implemented--%v", e.parent.Perma(lang), e.Slug(lang))
}
