package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch, email me at <a 
	href=\"mailto:huseyinozantemel@gmail.com\">huseyinozantemel@gmail.com</a>`)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1><p>Some questions and answers here
	<ul>
	  test 
	</ul>
	`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		// http.Error(w, "Page not found", http.StatusNotFound)
// 		// http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "Page not found")
// 		http.NotFound(w, r)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	// var router http.HandlerFunc = pathHandler
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	var router Router
	fmt.Println("Starting th server on :3000...")
	http.ListenAndServe(":3000", router)
	// http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
