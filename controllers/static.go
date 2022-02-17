package controllers

import (
	"html/template"
	"net/http"

	"github.com/bawazy/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! we offer a free trial for 30 days",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have support for 12hrs a day weekdays and saturdays.",
		},
		{
			Question: "How do i contact support?",
			Answer:   `You can contact support by sending us an email at <a href="mailto:hbawazy@gmail.com"> hbawazy@gmail.com</a>`,
		},
		{
			Question: "Do you have an office?",
			Answer:   "We are totally remote!",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
