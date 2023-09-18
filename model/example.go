// Projet d'evaluation
// Creation d'une API pour la gestion d'un bibliotheque

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

package model

import (
	"time"
)

// interface(CRUD)
type TenantDAO interface {
	CreateTenant(tenant *Tenant) error
	GetTenantById(id int) (*Tenant, error)
	UpdateTenant(id int) (*Tenant, error)
	DeleteTenant(id int)
}
type LivreDAO interface {
	CreateLivre(livre *Livre) error
	GetLivreById(id int) (*Livre, error)
	UpdateLivre(id int) (*Livre, error)
	DeleteLivre(id int)
}
type ClientDAO interface {
	CreateClient(client *Client) error
	GetClientById(id int) (*Client, error)
	UpdateClient(id int) (*Client, error)
	DeleteClient(id int)
}

type EmpruntDAO interface {
	CreateEmprunt(emprunt *Emprunt) error
	GetEmpruntById(id int) (*Emprunt, error)
	UpdateClient(id int) (*Emprunt, error)
	DeleteClient(id int)
}

// type CanGet interface {
// 	Get(int) error
// 	GetAll() ([]CanGet, error)
// }

// structure
/*Username string `json:"username"`
    Email    string `json:"email"`
    Age      int    `json:"age,omitempty"`
}
*/
type Livre struct {
	Id_livre        int    `json:"id_livre"`
	Id_bibliotheque int    `json:"id_bibliotheque"`
	Titre_livre     string `json:"titre_livre"`
	Auteur          string `json:"auteur"`
	Desc_livre      string `json:"description_livre"`
}
type Revue struct {
	Id_revue     int    `json:"id_revue"`
	Titre_revue  string `json:"titre_revue"`
	Auteur_revue string `json:"auteur_revue"`
	Des_revue    string `json:"description_revue"`
}
type Magazine struct {
	Id_magazine     int    `json:"id_magazine"`
	Titre_magazine  string `json:"titre_magazine"`
	Auteur_magazine string `json:"auteur_magazine"`
	Des_magazine    string `json:"description_magazine"`
}

type Tenant struct {
	Id_Tenant   int    `json:"id_tenant"`
	Name_tenant string `json:"nom_tenant"`
}

//var Liv = []Livre{}

type Emprunt struct {
	Id_emprunt       int       `json:"id_emprunt"`
	Id_livre_emprunt int       `json:"id_livre_emprunt"`
	Id_tenant        int       `json:"id_tenant"`
	Date_emprunt     time.Time `json:"date_emprunt"`
	Date_retour      time.Time `json:"date_retour"`
	//"date_field": m.DateField.Format("2006-01-02T15:04:05"), // Formater le champ time.Time
}

var emp Emprunt

//var Emprunter = []Emprunt{}

type Client struct {
	Id_client       int    `json:"id_client"`
	Nom_client      string `json:"nom_client"`
	Email_client    string `json:"email_client"`
	Password_client string `json:"password_client"`
}

//var cl Client

type Bibliotheque struct {
	Id_bibliotheque  int    `json:"id_bibliotheque"`
	Nom_bibliotheque string `json:"nom_bibliotheque"`
	Id_tenant        int    `json:"id_tenant"`
	Is_deleted       bool   `json:"is_deleted,omitempty"`
}

type Users struct {
	Id_user    int    `json:"id_user"`
	Id_tenant  int    `json:"id_tenant"`
	Nom_user   string `json:"nom_user"`
	Id_session int    `json:"id_session"`
	Id_role    int    `json:"id_role"`
	//	Id_role   int //on va avoir 3 roles: admin, bibliothecaire, simpleUser
	//	Permission []string // les permissions sont les actions que les roles peuvent effectuer
	Email_user    string
	Password_user string
}
type Session struct {
	Id_session  int       `json:"id_session"`
	Nom_session string    `json:"nom_session"`
	Id_user     int       `json:"id_user"`
	Duree       time.Time `json:"duree"`
	Token       string    `json:"token"`
}
type Role struct {
	Id_role  int    `json:"id_role"`
	Nom_role string `json:"nom_role"`
}

type Permission struct {
	Id_permission  int    `json:"id_permission"`
	Nom_permission string `json:"nom_permission"`
}

// func (b *Bibliothèque) CanGet(id int) error {

// }

// func get(c CanGet){
// 	c.Get()
// }

//var User = []Client{}
