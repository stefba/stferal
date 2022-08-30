package extra

import (
	"net/http"
	"sacer/go/server/head"
	"sacer/go/server"
	"sacer/go/server/users"
)

func Manifest(s *server.Server, w http.ResponseWriter, r *http.Request, a *users.Auth) {
	lang := head.Lang(r.Host)
	head := &head.Head{
		Title:   head.SiteName(lang),
		Desc:    s.Vars.Lang("site", lang),
		Auth:    a,
		Options: head.GetOptions(r),
		Host:    r.Host,
	}
	err := head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "manifest", head)
	if err != nil {
		s.Log.Println(err)
	}
}
