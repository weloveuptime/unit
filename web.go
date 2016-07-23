package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	http.HandleFunc("/", hello)
	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, %s world from %s", runtime.Version(), stringValue(os.Hostname, "localhost"))
}

func stringValue(f func() (string, error), defaultValue string) string {
    s, err := f()
    if s == "" && err != nil {
        s = defaultValue
    }
    return s
}
