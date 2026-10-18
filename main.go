package main

import (
	"html/template"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, "Could not load login page", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.Execute(w, "")
		if err != nil {
			log.Println("Template error:", err)
		}
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	message := "Invalid username or password!"

	if username == "barzan" && password == "1234" {
		message = "Login successful!"
	}

	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Could not load login page", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, message)
	if err != nil {
		log.Println("Template error:", err)
	}
}

func main() {
	http.HandleFunc("/", loginHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
