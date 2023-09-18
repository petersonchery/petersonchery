package utils

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Nom = "Ok pour utils"

func VerifError(e error) {
	if e != nil {
		fmt.Println("il y a erreur")
		return
	}
}

func LogWriter() error {
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatal("Impossible d'ouvrir le fichier de log :", err)
		return fmt.Errorf("y a erreur dans logwriter")
	}
	logrus.SetOutput(logFile)
	logrus.Errorf("Une erreur s'est produite : %s", "erreur critique")

	return nil
}
