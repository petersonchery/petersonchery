package service

/*
package main

import (
	"log"
	"os"
)

func main() {
	// Ouvrir un fichier pour écrire les logs
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Impossible d'ouvrir le fichier de log :", err)
	}
	defer logFile.Close()

	// Configurer le logger pour écrire dans le fichier
	log.SetOutput(logFile)

	// Écrire des logs
	log.Println("Ceci est un message de log.")
	log.Printf("Un autre message avec un paramètre : %d", 42)
}

package main

import (
	"os"
	"github.com/sirupsen/logrus"
)

func main() {
	// Ouvrir un fichier pour écrire les logs
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatal("Impossible d'ouvrir le fichier de log :", err)
	}
	defer logFile.Close()

	// Configurer le logger pour écrire dans le fichier
	logrus.SetOutput(logFile)

	// Écrire des logs avec logrus
	logrus.Info("Ceci est un message de log.")
	logrus.Errorf("Une erreur s'est produite : %s", "erreur critique")

	// Vous pouvez également définir d'autres niveaux de log comme Debug, Warning, etc.
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Ce message ne sera visible que si le niveau de log est défini sur Debug.")
}

*/
