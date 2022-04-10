package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Atlas-Compute-Platform/lib"
)

func apiReset(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	dataSet = make(lib.Dict)
}

func apiList(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	if err := lib.SendDict(dataSet, w); err != nil {
		lib.LogError(os.Stderr, "main.apiList", err)
	}
}

func apiLoad(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var (
		key string = r.URL.Query().Get(lib.KEY_KEYS)
		val string
		ok  bool
	)

	if val, ok = dataSet[key]; !ok {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "%s", val)
}

func apiStore(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var (
		key string = r.URL.Query().Get(lib.KEY_KEYS)
		val string = r.URL.Query().Get(lib.KEY_VALS)
	)

	dataSet[key] = val
}

func apiRemove(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var key string = r.URL.Query().Get(lib.KEY_KEYS)
	delete(dataSet, key)
}
