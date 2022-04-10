package main

/*
	Atlas Key/Value Store
	Thijs Haker
*/

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Atlas-Compute-Platform/lib"
)

var dataSet lib.Dict

func usage() {
	fmt.Fprintf(os.Stderr, "Atlas Key/Value Store %s\n", lib.VERS)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var (
		netAddr  = flag.String("p", lib.PORT, "Specify port")
		initData = flag.String("i", "{}", "Specify initial data)")
		err      error
	)
	flag.Usage = usage
	flag.Parse()

	if dataSet, err = lib.ImportDict([]byte(*initData)); err != nil {
		lib.LogError(os.Stderr, "main.main", err)
		os.Exit(1)
	}

	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/reset", apiReset)
	http.HandleFunc("/list", apiList)
	http.HandleFunc("/load", apiLoad)
	http.HandleFunc("/store", apiStore)
	http.HandleFunc("/remove", apiRemove)

	http.ListenAndServe(*netAddr, nil)
}
