package main

import (
	"fmt"
	"log"
	"net/http"
)

// cette fonction doit respecter cette signature pour être un handlerfunc
func hello(w http.ResponseWriter, r *http.Request) {
	// Ecriture d'un message en réponse http
	fmt.Printf("%v %v\n", r.Method, r.URL)
	fmt.Fprintf(w, "Hello tout le monde")
}

// une requête ressemble à cela
// http://localhost:8585/goodbye
func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\n", r.Method, r.URL)
	fmt.Fprintf(w, "Goodbye gys")
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "login.html")
		return
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() failed. err=%v", err)
			return
		}

		// Recuperation des données du formulaire
		fmt.Fprintf(w, "Go login POST. value=%v\n", r.PostForm)
		username := r.FormValue("username")
		password := r.FormValue("password")
		// verification des informations
		if username == "Go" && password == "Aissa" {
			fmt.Fprintf(w, "You are logged\n")
		} else {
			fmt.Fprint(w, "Wrong username / password\n ")
		}
	}
}

func main() {
	//associe un chemin à une fonction
	http.HandleFunc("/", hello)
	http.HandleFunc("/goodbye", goodbye)
	http.HandleFunc("/login", login)

	// création du serveur http écoutant sur le port 8585
	if err := http.ListenAndServe(":8585", nil); err != nil {
		log.Fatal(err)
	}
}
