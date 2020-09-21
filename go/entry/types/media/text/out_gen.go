// Code generated by go generate; DO NOT EDIT.

package text

import (
	"fmt"

	"sacer/go/entry"
	"sacer/go/entry/helper"
	"sacer/go/entry/parts/file"
	"sacer/go/entry/parts/info"
	"time"
)

func (e *Text) Type() string {
	return "text"
}

func (e *Text) Parent() entry.Entry {
	return e.parent
}

func (e *Text) File() *file.File {
	return e.file
}

func (e *Text) Id() int64 {
	return e.date.Unix()
}

func (e *Text) Timestamp() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Text) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Text) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Text) Date() time.Time {
	return e.date
}

func (e *Text) Info() info.Info {
	return e.info
}

func (e *Text) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Text) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.info.Title(lang))
}

func (e *Text) IsBlob() bool {
	return entry.IsBlob(e)
}

func (e *Text) MediaType() string {
	if e.Type() == "audio" {
		return "tob"
	}
	if entry.IsBlob(e) {
		return "blob"
	}
	return "tob"
}

func (e *Text) SetParent(parent entry.Entry) {
	e.parent = parent
}

func (e *Text) SetInfo(inf info.Info) {
	e.info = inf
}

func (e *Text) Path(lang string) string {
	return fmt.Sprintf("%v/%v", e.parent.Path(lang), e.Slug(lang))
}

// This recursive function call will be caught by a Tree type. For now, all
// further up parent entries are exclusively of type Tree.
func (e *Text) Section() string {
	return e.Parent().Section()
}

func (e *Text) Perma(lang string) string {
	if e.parent.Type() == "set" {
		return e.parent.Perma(lang)
	}

	name := e.Hash()
	if slug := e.Slug(lang); slug != "" {
		name = fmt.Sprintf("%v-%v", slug, e.Hash())
	}

	switch e.Section() {
	case "index":
		return fmt.Sprintf("%v#%v", e.parent.Perma(lang), helper.Normalize(e.Title(lang)))
	case "kine":
		return fmt.Sprintf(
			"/%v/%v/%v",
			helper.KineName[lang],
			e.Date().Format("06-01"),
			fmt.Sprintf("%v-%v", e.Date().Format("02"), name),
		)
	}

	return fmt.Sprintf("%v/%v", e.parent.Path(lang), name)
}
