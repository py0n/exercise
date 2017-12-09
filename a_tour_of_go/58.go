package tour

import (
	"fmt"
	"net/http"
)

// String ...
type String string

// SearveHTTP ...
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

// Struct ...
type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

// ServeHTTP ...
func (st *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, st.Greeting+st.Punct+st.Who)
}

func startWebServer2() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}

// C58 ...
func C58(opts Options, args []string) int {
	startWebServer2()
	return 0
}
