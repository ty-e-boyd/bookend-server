package main

import (
	"bookend/inits"
	"bookend/models"
	"log"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	err := inits.DB.AutoMigrate(&models.Entry{})
	if err != nil {
		log.Fatalf("Unable to complete migration, %v", err)
		return
	}
}
