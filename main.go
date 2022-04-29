package main

/*
	Atlas Key/Value Store
	Thijs Haker
*/

import (
	"flag"
	"net/http"
	"os"

	"github.com/Atlas-Compute-Platform/lib"
)

var dataSet lib.Dict

func main() {
	lib.SvcName = "Atlas Key/Value Store"
	lib.SvcVers = "1.0"

	var (
		netAddr  = flag.String("p", lib.PORT, "Specify port")
		initData = flag.String("i", "{}", "Specify initial data)")
		err      error
	)
	flag.Usage = lib.Usage
	flag.Parse()

	if dataSet, err = lib.ImportDict([]byte(*initData)); err != nil {
		lib.LogFatal(os.Stderr, "main.main", err)
	}
	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/reset", apiReset)
	http.HandleFunc("/list", apiList)
	http.HandleFunc("/load", apiLoad)
	http.HandleFunc("/store", apiStore)
	http.HandleFunc("/remove", apiRemove)

	http.ListenAndServe(*netAddr, nil)
}
