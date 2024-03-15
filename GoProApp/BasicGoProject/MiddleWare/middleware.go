package main

import (
	"log"
	"net/http"
	"time"
)

// method -- 1
// Custom Logging
type Logger struct { //step -- 2
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) { //step -- 3  //It is called Automatically by the HTTPServer
	//Implementing interface
	start := time.Now()

	//calling next Handler or original Handler.
	l.handler.ServeHTTP(w, r)
	log.Printf("%v", time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger { // step --1
	return &Logger{handlerToWrap}
}

// // method -- 2
// // basic logging middleware
// func logging(f http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) { //this the --- actual Handler Request
// 		start := time.Now()
// 		log.Println(r.URL.Path)
// 		log.Println(start)

// 		// Call the next handler in the middlewarechain -- (calling the original Handler Request)
// 		f(w, r)
// 	}
// }

func main() {

	//creating http router/mux
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)       //logging(helloHandler)) //-- NewLogger(helloHandler)
	mux.HandleFunc("/abstract", abstractHandler) //logging(abstractHandler))

	//wrap entire mux into NewLogger
	wrapMux := NewLogger(mux) //I can send multiple HandlerFunc at a time or single HandlerFunc

	//activation the server
	http.ListenAndServe(":8081", wrapMux) //mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello this Home Page"))
}
func abstractHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello this Abstract Page"))
}
