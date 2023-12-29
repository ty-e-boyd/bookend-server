package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"net/http"
	"os"
)

func SendAText(ctx *gin.Context) {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody("Hello from Golang!")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		fmt.Println("SMS sent successfully!")
		ctx.JSON(http.StatusOK, gin.H{"message": "message sent successfully"})
	}
}
