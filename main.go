package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1> Welcome to my awesone website </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1> <p>To get in touch, email me at <a href=\"mailto:hbawazy@gmail.com\">hbawazy@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, `
	<h1>FAQ PAGE</h1> 
	<ul>
	<li><b>Is there a free versiion?</b> Yes! we offer a free trial for 30 days</li>
	<li><b>What are your support hours?</b> We have support for 12hrs a day weekdays and saturdays.</li>
	<li><b>How do i contact support?</b> You can contact support by sending us an email at <a href="mailto:hbawazy@gmail.com"> hbawazy@gmail.com</a></li>
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
// 		//TODO: handle page not found error
// 		http.Error(w, "Page Not found", http.StatusNotFound)
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
		//TODO: handle page not found error
		http.Error(w, "Page Not found", http.StatusNotFound)
	}

}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
