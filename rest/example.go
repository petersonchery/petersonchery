package rest

import (
	"FirstProject/controller"
	"FirstProject/model"
	"FirstProject/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var Nom = "ok pour route"

//donneUser

func Welcome() {
	fmt.Println("Bienvenue sur mon API")
	fmt.Println("Voici le menu")
}

func Menu() {
	var choix int
	var clavier = bufio.NewScanner(os.Stdin)
	fmt.Println("1.- Ajouter un L")
	fmt.Println("2.- retirer un L")
	fmt.Println("3.- lister tout Les L")
	fmt.Println("Faire votre choix")

	clavier.Scan()
	choix, err := strconv.Atoi(clavier.Text())
	utils.VerifError(err)

	switch choix {
	case 1:
		fmt.Println("Ajouter un livre")
		GestionDeRoute2()
	case 2:
		fmt.Println("Retirer un L")
	case 3:
		fmt.Println("Lister tout les")
	default:
		fmt.Println("It's some other.")
	}

}

func GestionDeRoute2() {
	// Hello world, the web server
	var l model.Livre
	l = controller.Ajouter_livre()
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		//io.WriteString(w, "Hello, world!\n")
		//w.Write([]byte("Hello, world!\n"))
		fmt.Println("new request 1")
		//fmt.Fprintf(w, "Hello, %s!\n", "inconnu")
		fmt.Fprintf(w, "hello Voici les donnees qui sont enregistres; Titre du livre: %s Identifie au numero  %d ", l.Titre_livre, l.Id_livre)
		fmt.Println(controller.Ajouter_livre())

	}

	helloHandler2 := func(w http.ResponseWriter, req *http.Request) {
		// l := req.ContentLength
		// b := make([]byte, l)
		// req.Body.Read(b)
		var person model.Tenant
		b, err := io.ReadAll(req.Body)
		utils.VerifError(err)
		json.Unmarshal(b, &person)
		person.Id_Tenant *= 2
		// json.NewDecoder(req.Body).Decode(&person)
		fmt.Fprintf(w, "hello %s, you are %d years old!!", person.Name_tenant, person.Id_Tenant)
		b, err = json.Marshal(person)
		utils.VerifError(err)

		//json.NewEncoder(w).Encode(person)
		//fmt.Fprintf(w, "%v", b)
		//fmt.Print(w, "hello %s, you are %d years old!!", person.Name, person.Age)
		w.Write(b)
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello/", helloHandler2)
	log.Fatal(http.ListenAndServe(":8080", nil))
	/////
	////
}

/*
// cote serveur
func GestionDeRoute1() {
	fmt.Println("On est dans la route")
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("/pixabay.com/images/search/nature/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/
