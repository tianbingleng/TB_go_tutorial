package server3

import (
	"fmt"
	"log"
	"net/http"

	lissajous "github.com/Unified/TB_go_tutorial/ch1/lissajous"
)

func main() {
	http.HandleFunc("/info", info) // each request calls handler
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	}
	http.HandleFunc("/lissajous", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
