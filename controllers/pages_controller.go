package controllers

import (
	"api/database"
	"api/models"
	"api/services"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func PublicPages(c *gin.Context) {
	database.StartMongoDB()
	const Beare_schema = "Bearer "

	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatus(401)
	}
	token := header[len(Beare_schema):]

	claim, _ := services.NewJWTServices().ExtractClaims(token)

	id := fmt.Sprint(claim["sum"])

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	myString := string(jsonData[:])

	var diario models.Diario

	id_user, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	diario.Id_user = id_user
	diario.Content = myString
	diario.CreatedAt = time.Now()

	db_mongo := database.GetDatabaseMongo()

	result, err := db_mongo.InsertOne(context.Background(), diario)
	if err != nil {
		panic(err)
	}

	// c.JSON(200, result.InsertedID)
	c.JSON(200, result)
}

func ListPages(c *gin.Context) {
	database.StartMongoDB()
	const Beare_schema = "Bearer "

	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatus(401)
	}

	token := header[len(Beare_schema):]

	claim, _ := services.NewJWTServices().ExtractClaims(token)

	id := fmt.Sprint(claim["sum"])
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db_mongo := database.GetDatabaseMongo()

	listpages := []models.Diario{}
	pages, err := db_mongo.Find(context.Background(), bson.D{{"id_user", newid}})
	defer pages.Close(context.Background())

	for pages.Next(context.Background()) {
		pagesResult := models.Diario{}

		err := pages.Decode(&pagesResult)
		if err != nil {
			log.Fatal(err)
		}
		listpages = append(listpages, pagesResult)
	}
	c.JSON(200, listpages)
}
