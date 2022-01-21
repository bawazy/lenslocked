package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	userId := chi.URLParam(r, "userId")

	fmt.Fprintf(w, `<h1> User:%v </h1>`, userId)

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/users/{userId}", userHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
