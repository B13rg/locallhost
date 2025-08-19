package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func respIndex(writer http.ResponseWriter, req *http.Request) {
	var tmplFile = "index.tmpl.html"
	tmpl, err := template.New(tmplFile).Parse(IndexTemplateString())
	if err != nil {
		//nolint:errcheck
		fmt.Fprintf(writer, "Error parsing template: %v\n", err)

		return
	}

	reqData := ExtractRequestData(req)

	err = tmpl.Execute(writer, reqData)
	if err != nil {
		//nolint:errcheck
		fmt.Fprintf(writer, "Error executing template: %v\n", err)
	}
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
		Host:       req.Host,
		Header:     req.Header,
	}
}

func respIP(w http.ResponseWriter, req *http.Request) {
	remoteAddr, _ := parseAddress(req.RemoteAddr)
	//nolint:errcheck
	fmt.Fprintf(w, "%v", remoteAddr)
}

func respJson(w http.ResponseWriter, req *http.Request) {
	data := ExtractRequestData(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//nolint:errcheck,gosec,errchkjson
	json.NewEncoder(w).Encode(data)
}

// Start serving on specified port.
func Serve(port int) {
	http.HandleFunc("/", respIndex)
	http.HandleFunc("/ip", respIP)
	http.HandleFunc("/json", respJson)

	//nolint:gosec
	log.Fatal().Err(http.ListenAndServe(":"+strconv.Itoa(port), nil)).Msg("Server encountered error")
}
