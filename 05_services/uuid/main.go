package main

import (
	"fmt"
	"net/http"

	"strings"

	"github.com/satori/go.uuid"
)

func init() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	u := uuid.NewV4()
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "<h2>%v</h2><br>", u.String())
	fmt.Fprintf(w, "<a href=\"http://%v\">Link to Main</a>",
		strings.Replace(r.Host, "uuid.", "", 1))
}
