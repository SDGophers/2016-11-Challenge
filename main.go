package main

import (
	"html/template"
	"log"
	"net/http"
)

type UserData struct {
	FullName string `name:"fullname" desc:"Full name"`
	Email    string `name:"email" desc:"Email address"`
	Phone    string `name:"phone" desc:"Phone number"`
}

type Field struct {
	InputName string
	Desc      string
}

type Form struct {
	Heading string
	Fields  []Field
}

func (f *Form) ParseFields(s interface{}) {

}

func main() {

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		output := "parse user data here\n"
		w.Write([]byte(output))
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		d := UserData{}
		form := &Form{Heading: "User Information"}
		form.ParseFields(d)

		t, err := template.ParseFiles("tpl/form.html")
		if nil != err {
			log.Println(err)
		}
		t.Execute(w, form)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
