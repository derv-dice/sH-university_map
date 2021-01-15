package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func adminHandler() http.Handler {
	mux := httprouter.New()
	mux.GET("/admin/", hello)
	// TODO admin router

	return adminAuth(mux)
}

func privateHandler() http.Handler {
	mux := httprouter.New()
	mux.GET("/private/", hello)
	// TODO private router

	return privateAuth(mux)
}

func publicHandler() http.Handler {
	mux := httprouter.New()
	mux.GET("/", hello)
	// TODO public router

	return mux
}

func MainHandler() (handler http.Handler) {
	adminHandler := adminHandler()
	privateHandler := privateHandler()
	publicHandler := publicHandler()

	mux := http.NewServeMux()
	mux.Handle("/admin/", adminHandler)
	mux.Handle("/private/", privateHandler)
	mux.Handle("/", publicHandler)

	handler = accessLog(mux)
	handler = recovery(handler)
	return
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("hello, from: %s", r.URL.Path)))
	w.WriteHeader(200)
}

func adminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO auth
		next.ServeHTTP(w, r)
	})
}

func privateAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO auth
		next.ServeHTTP(w, r)
	})
}

func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				w.Write([]byte("Internal Server Error"))
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func accessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("access_log: {method: %s, ip: %s, url: %s, time: %s}", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}
