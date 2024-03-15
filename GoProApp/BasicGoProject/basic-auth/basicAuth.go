/*Using Basic Authentication using ----  http mux*/
package main

import (
	"errors"
	"fmt"
	"net/http"
)

// declaring gobal var using maping
var (
	credentails = map[string]string{
		"user": "Dhaneshwar"}
)

// func main() {
// 	//routing http
// 	mux := http.NewServeMux()

// 	//Requests
// 	mux.HandleFunc("/login", homePage)
// 	mux.HandleFunc("/home", loginPage)

// 	//Activating the Server
// 	http.ListenAndServe(":8081", mux)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	username, password, ok := r.BasicAuth() //returns the username,pwd,and error
// 	if !ok {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Unauthorized 401")
// 	}

// 	if credPasswd, found := credentails[username]; found && password == credPasswd {
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprint(w, "Welocome to the homePage")
// 	} else {
// 		//Need to add the Headers.
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "UnAuthorized")
// 	}
// }

// func loginPage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "UnAuthorized")
// }

// Middle Basic Authentication
// func GetAuthenticated(r *http.Request) error {
// 	username, password, ok := r.BasicAuth()
// 	if !ok {
// 		//errors is pkg
// 		return errors.New("authentication credentials not provided")
// 	}
// 	if credPasswd, found := credentails[username]; found && password == credPasswd {
// 		return nil
// 	} else {
// 		return errors.New("InValid Credentails")
// 	}
// }

// func main() {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/home", homePage)
// 	mux.HandleFunc("/login", loginPage)

// 	http.ListenAndServe(":8081", mux)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	err := GetAuthenticated(r)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "UnAuthorized")
// 		// http.Error(w, "please sign-in", http.StatusUnauthorized)
// 		return
// 	}
// 	fmt.Fprint(w, "welcome to homePage")
// }

// func loginPage(w http.ResponseWriter, r *http.Request) {
// 	err := GetAuthenticated(r)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "UnAuthorized")
// 		return
// 	}
// 	fmt.Fprint(w, "Welcome to loginPage")
// }

// middleWare Auth
type Basic struct {
	handler http.Handler
}

// For Authentication
func GetAuthenticated(r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if !ok {
		//errors is pkg
		return errors.New("authentication credentials not provided")
	}
	if credPasswd, found := credentails[username]; found && password == credPasswd {
		return nil
	} else {
		return errors.New("InValid Credentails")
	}
}

// Implementing the http.Hander method ServeHttp
func (b *Basic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := GetAuthenticated(r)
	if err != nil {
		http.Error(w, "Please... Enter the Valid Details to Login.", http.StatusUnauthorized)
		return
	}

	//calling the original Handler or Request
	b.handler.ServeHTTP(w, r)
}

func NewLogger(handlerToWrap http.Handler) *Basic { // step --1
	return &Basic{handlerToWrap}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", homePage)
	mux.HandleFunc("/login", loginPage)

	//Wrapping
	wrapMux := NewLogger(mux)

	http.ListenAndServe(":8081", wrapMux)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welome to homePage")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welome to loginPage")
}
