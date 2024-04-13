package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Définition du gestionnaire pour la route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Envoie d'un message de bienvenue au client
		fmt.Fprintf(w, "Bienvenue sur notre serveur web !")
	})

	// Définition du gestionnaire pour la route "/hello"
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Envoie d'un message spécifique pour la route /hello au client
		fmt.Fprintf(w, "Bonjour ! Vous avez atteint la route /hello.")
	})

	// Démarrage du serveur sur le port 9090
	fmt.Println("Serveur démarré sur http://localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		// Affichage d'une erreur en cas de problème lors du démarrage du serveur
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
