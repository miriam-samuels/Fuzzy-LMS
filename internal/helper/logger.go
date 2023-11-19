package helper

import (
	"log"
	"os"
)

func InitLogger() {
	// OPEN FILE (create if it doesn't exist and open file for writing)
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	defer file.Close()

	// Set the output of the logger to the file
	log.SetOutput(file)

}
