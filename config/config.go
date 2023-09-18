package config

// config/database.go
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func InitDB(connectionString string) {
	var err error

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connexion à la base de données PostgreSQL réussie")

}
