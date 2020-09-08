package front

import (
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/server"
)

type frontMain struct {
	Head     *head.Head
	Index    entry.Entries
	Graph    entry.Entries
	Video    entry.Entries
	Featured entry.Entry
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	lang := head.Lang(r.Host)

	head := &head.Head{
		Title:   "",
		Section: "home",
		Path:    "/",
		Host:    r.Host,
		Entry:   nil,
		Desc:    s.Vars.Lang("site", lang),
		Options: head.GetOptions(r),
	}
	err := head.Process()
	if err != nil {
		return
	}

	index := s.Recents["index"].Local(s.Flags.Local)[lang]
	graph := s.Recents["graph"].Local(s.Flags.Local)[lang]

	video := s.Recents["video"].Local(s.Flags.Local)[lang]

	e, err := s.Trees["graph"].Local(s.Flags.Local)[lang].LookupEntryHash(s.Vars.FrontSettings.Featured)
	if err != nil {
		s.Log.Println(err)
	}

	err = s.ExecuteTemplate(w, "front", &frontMain{
		Head:     head,
		Index:    index.Offset(0, s.Vars.FrontSettings.Index),
		Graph:    graph.Offset(0, s.Vars.FrontSettings.Graph),
		Video:    video.Offset(0, 10),
		Featured: e,
	})
	if err != nil {
		log.Println(err)
	}
}
