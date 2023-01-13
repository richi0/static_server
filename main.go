package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

var USERNAME string
var PASSWORD string

func basicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			if username == USERNAME && password == PASSWORD {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func teardown(runtime int) {
	if runtime == 0 {
		return
	}
	go func() {
		time.Sleep(time.Minute * time.Duration(runtime))
		os.Exit(0)
	}()
}

func main() {
	port := flag.String("port", "4000", "Port number")
	folder := flag.String("folder", ".", "Folder to serve from")
	username := flag.String("username", "username", "BasicAuth username")
	password := flag.String("password", "password", "BasicAuth password")
	runtime := flag.Int("runtime", 0, "Time the serve will be available in minutes (0 == always)")
	flag.Parse()

	USERNAME = *username
	PASSWORD = *password
	teardown(*runtime)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(*folder))
	mux.Handle("/", fs)

	log.Printf("Serving folder %s on :%s...", *folder, *port)
	err := http.ListenAndServe(":"+*port, basicAuthMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}

}
