package main

import (
	"go_category/configs"
	"go_category/domain/db"
	"log"
)

func setLogFile() {
	log.SetOutput(configs.LogFile)
}

func main() {
	setLogFile()
	defer configs.CloseLogFile()
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	db.Automigrate()
}
