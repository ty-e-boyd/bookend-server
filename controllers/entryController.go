package controllers

import (
	"bookend/inits"
	"bookend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateEntry(ctx *gin.Context) {
	var body struct {
		Title      string
		Body       string
		BookName   string
		Author     string
		PageNumber string
	}

	err := ctx.BindJSON(&body)
	if err != nil {
		log.Panicln("error binding request body to internal struct")
	}

	entry := models.Entry{
		Title:      body.Title,
		Body:       body.Body,
		BookName:   body.BookName,
		Author:     body.Author,
		PageNumber: body.PageNumber,
	}

	result := inits.DB.Create(&entry)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": entry})
}

func GetEntries(ctx *gin.Context) {
	var entries []models.Entry

	result := inits.DB.Find(&entries)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": entries})
}

func GetEntry(ctx *gin.Context) {
	var entry models.Entry

	result := inits.DB.First(&entry, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": entry})
}
