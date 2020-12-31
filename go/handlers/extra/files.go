package extra

import (
	"fmt"
	"log"
	"net/http"
	p "path/filepath"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/set"
	"sacer/go/entry/types/video"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/server/paths"
	"strings"
)

/*

	TODO: This package should be simplified.

*/

func ServeFile(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth, path *paths.Path) {
	err := serveFile(s, w, r, a, path)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}
}

func serveFile(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth, path *paths.Path) error {
	section := path.Section()
	lang := head.Lang(r.Host)
	tree := s.Trees[section].Access(a.Subscriber)[lang]
	e, err := tree.LookupEntryHash(path.Hash)
	if err != nil {
		return err
	}

	col, ok := e.(entry.Collection)

	if !ok {
		return serveSingleBlob(w, r, e, path)
	}

	return serveCollectionBlob(w, r, col, path)
}

func serveSingleBlob(w http.ResponseWriter, r *http.Request, e entry.Entry, path *paths.Path) error {
	blob, ok := e.(entry.Blob)
	if !ok {
		return fmt.Errorf("File to serve (%v) is no blob.", e.File().Name())
	}
	v, ok := e.(*video.Video)
	if ok {
		switch p.Ext(path.SubFile.Name) {
		case ".vtt":
			serveStatic(w, r, v.SubtitleLocation(path.SubFile.Size))
			return nil
		}
	}
	location, err := blob.Location(path.SubFile.Size)
	if err != nil {
		return err
	}
	serveStatic(w, r, location)
	return nil
}

func serveCollectionBlob(w http.ResponseWriter, r *http.Request, col entry.Collection, path *paths.Path) error {
	name := baseName(path.SubFile.Name)
	for _, e := range col.Entries() {
		if baseName(e.File().Name()) == name {
			return serveSingleBlob(w, r, e, path)
		}
	}

	if name := path.SubFile.Name; len(name) > 5 && name[:5] == "cover" {
		set, ok := col.(*set.Set)
		if ok && set.Cover != nil {
			return serveSingleBlob(w, r, set.Cover, path)
		}
		return fmt.Errorf("serveCollectionBlob: Cover %v not found.", path.SubFile.Name)
	}

	return fmt.Errorf("serveCol !N!,onBlob: File %v not found.", path.SubFile.Name)
}

func baseName(name string) string {
	name = stripBlur(name)
	return stripSize(name)
}

func stripSize(name string) string {
	i := strings.LastIndex(name, "-")
	if i > 0 {
		return name[:i]
	}
	return name
}

func stripBlur(name string) string {
	name = tools.StripExt(p.Base(name))
	if l := len(name); l > 4 && name[l-4] == '_' {
		return name[:l-4]
	}
	return name
}
