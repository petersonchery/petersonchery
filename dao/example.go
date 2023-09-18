/*
Les relations entre les entités "bibliothèque", "livres" et "emprunteurs"
dans un système de bibliothèque sont généralement définies par le modèle de données
que vous choisissez d'implémenter. Voici quelques relations typiques qui pourraient
exister entre ces entités :

Relation entre Bibliothèque et Livres :

Une bibliothèque peut contenir de nombreux livres, mais un livre peut appartenir à

	une seule bibliothèque (relation un-à-plusieurs).

Cette relation peut être modélisée en ajoutant une clé étrangère "bibliothèque_id"

	la table "livres" qui fait référence à la table "bibliothèque".

Relation entre Emprunteurs et Livres :

Un emprunteur peut emprunter plusieurs livres, et un livre peut être emprunté par

	emprunteurs (relation plusieurs-à-plusieurs).

Cette relation nécessite l'utilisation d'une table de jointure (par exemple, "emprunts")
qui enregistre les emprunts spécifiques avec des clés étrangères vers les tables "emprunteurs" et

	"livres".

Relation entre Emprunteurs et Bibliothèque :

Un emprunteur peut être inscrit dans une bibliothèque, mais une bibliothèque peut avoir
plusieurs emprunteurs (relation un-à-plusieurs).
Cette relation peut être modélisée en ajoutant une clé étrangère "bibliothèque_id" dans
la table "emprunteurs" qui fait référence à la table "bibliothèque".
*/
package dao

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var dB *sql.DB

func Connect() {
	// Chaîne de connexion à la base de données PostgreSQL
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable host=172.21.64.1 port=5432"
	//postgres://user:pass@localhost:8081/database?sslmode=disable pgsql=192.168.10.146&username=postgres&db=mydb&ns=public
	// Connexion à la base de données PostgreSQL
	db, err := sql.Open("postgres", connStr)
	fmt.Println("on est la aussi")
	if err != nil {
		fmt.Println("ca n'a pas marche")
		log.Fatal(err)
		return
	}
	defer db.Close()

	// Vérification de la connexion
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connexion à la base de données PostgreSQL réussie")

	// Création de la table

}

func insert(id int, nom string) {
	//insertion
	insertSQL := "INSERT INTO bibliotheque (id_bibliotheque, nom_bibliotheque) VALUES ($1, $2)"
	_, err := dB.Exec(insertSQL, "33", "Mon Livre")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Donnees entrees avec succès")

	/*
		fmt.Println("Valeurs insérées avec succès")

		// Exemple d'exécution de requête
		rows, err := db.Query("SELECT * FROM biblio")
		if err != nil {
			fmt.Println("pas marche")
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("ID: %d, Name: %s\n", id, name)
		}
	*/
}
