package tour

import (
	"fmt"
	"net/http"
)

type hello struct{}

// ServeHTTP ...
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func startWebServer() {
	var h hello
	http.ListenAndServe("localhost:4000", h)
}

// C57 ...
func C57(opts Options, args []string) int {
	startWebServer()
	return 0
}
