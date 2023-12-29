package controllers

import (
	"bookend/inits"
	"bookend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func SendEmail(ctx *gin.Context) {
	url := "https://maker.ifttt.com/trigger/request_email/json/with/key/" + os.Getenv("IFTTT_WEBHOOK_KEY")

	// get entries
	var entries []models.Entry

	result := inits.DB.Find(&entries)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	type body struct {
		Title string
		Body  string
		Test  string
	}

	// choose a random entry
	randomIndex := rand.Intn(len(entries))
	pick := entries[randomIndex]

	out, err := json.Marshal(body{
		Title: pick.Title,
		Body:  pick.Body,
	})
	if err != nil {
		log.Fatalln("unable to marshal")
	}

	// create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(out))
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
