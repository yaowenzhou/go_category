package main

import (
	"go_category/consts"
	"go_category/domain/db"
	"log"
	"os"
)

func setLogFile() {
	f, err := os.OpenFile(consts.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}

func main() {
	setLogFile()
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	db.Automigrate()
}
