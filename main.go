package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type SettingHandler struct {
	config   map[string][]MethodModel
	notFound NotFoundModel
}

func (sh *SettingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	methods := sh.config[r.URL.Path]

	if len(methods) == 0 {
		responseBodyNotFound, err := json.Marshal(sh.notFound)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(responseBodyNotFound)
		}
		return
	}

	var isResponse = false
	for _, method := range methods {
		if strings.EqualFold(method.Method, r.Method) {

			responseHeaders := method.Response.Headers

			if len(responseHeaders) == 0 {
				w.Header().Set("Content-Type", "application/json")
			} else {
				for k, v := range responseHeaders {
					w.Header().Set(k, v)
				}
			}

			responseBody := method.Response.Body
			statusCode := method.Response.Status
			w.WriteHeader(statusCode)
			if err := json.NewEncoder(w).Encode(responseBody); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			isResponse = true
			return
		}
	}

	if !isResponse {
		responseBodyNotFound, err := json.Marshal(sh.notFound)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(responseBodyNotFound)
		}
	}
}

func startServer(setting Settings) {

	config := map[string][]MethodModel{}

	for _, paths := range setting.Path {

		config[paths.Path] = append(config[paths.Path], paths.Methods...)
	}

	handlerFn := &SettingHandler{config: config, notFound: setting.NotFound}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlerFn.ServeHTTP)
	addr := fmt.Sprintf(":%d", setting.Port)
	fmt.Println(setting.Name + " Server is running on " + addr)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	server.ListenAndServe()
}

func main() {

	var file = "servers.json"

	if len(os.Args) >= 2 {
		file = os.Args[1]
	}

	settings, err := GetSettings(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	//run multiples servers

	for _, setting := range settings {
		go startServer(setting)
	}

	select {}
}
