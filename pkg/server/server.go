package server

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func respIndex(w http.ResponseWriter, req *http.Request) {
	var tmplFile = "index.tmpl.html"
	tmpl, err := template.New(tmplFile).Parse(IndexTemplateString)
	if err != nil {
		fmt.Fprintf(w, "Error parsing template\n")
		return
	}
	err = tmpl.Execute(w, req)
	if err != nil {
		fmt.Fprintf(w, "Error executing template\n")
		return
	}
	// for name, headers := range req.Header {
	// 	for _, h := range headers {
	// 		fmt.Fprintf(w, "%v: %v\n", name, h)
	// 	}
	// }
}

func respIP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v", req.RemoteAddr)
}

// Start serving on specified port.
// If port is <= 0, defaults to 8080.
func Serve(port int) {
	if port <= 0 {
		port = 8080
	}
	http.HandleFunc("/", respIndex)
	http.HandleFunc("/ip", respIP)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
