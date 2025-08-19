package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"text/template"
)

func respIndex(w http.ResponseWriter, req *http.Request) {
	var tmplFile = "index.tmpl.html"
	tmpl, err := template.New(tmplFile).Parse(indexTemplateString)
	if err != nil {
		fmt.Fprintf(w, "Error parsing template\n")
		return
	}

	reqData := ExtractRequestData(req)

	err = tmpl.Execute(w, reqData)
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

func parseAddress(remoteAddr string) (string, string) {
	addrParts := strings.Split(remoteAddr, ":")

	// extract and remove port
	slices.Reverse(addrParts)
	retPort := addrParts[0]
	addrParts = slices.Delete(addrParts, 0, 1)
	// recombine
	slices.Reverse(addrParts)
	retAddr := strings.Join(addrParts, ":")

	return retAddr, retPort
}

func ExtractRequestData(req *http.Request) *RequestResponse {
	remoteAddr, remotePort := parseAddress(req.RemoteAddr)

	return &RequestResponse{
		RemoteAddr: remoteAddr,
		RemotePort: remotePort,
		Method:     req.Method,
		Proto:      req.Proto,
		Header:     req.Header,
	}
}

func respIP(w http.ResponseWriter, req *http.Request) {
	remoteAddr, _ := parseAddress(req.RemoteAddr)
	fmt.Fprintf(w, "%v", remoteAddr)
}

func respJson(w http.ResponseWriter, req *http.Request) {
	data := ExtractRequestData(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

// Start serving on specified port.
func Serve(port int) {
	http.HandleFunc("/", respIndex)
	http.HandleFunc("/ip", respIP)
	http.HandleFunc("/json", respJson)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
