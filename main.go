//go:generate go/entry/types/generate

package main

import (
	"net/http"
	"stferal/go/routes"
	"stferal/go/server"
)

func main() {
	s := server.NewServer()

	err := s.Load()
	if err != nil {
		s.Log.Fatal(err)
	}

	http.Handle("/", routes.Router(s))
	http.ListenAndServe(":8013", nil)
}
