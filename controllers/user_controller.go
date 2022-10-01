package controllers

import (
	"api/database"
	"api/models"
	"api/services"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func GetUser(c *gin.Context) {
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

	db := database.GetDatabase()

	var user models.User

	err = db.First(&user, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user" + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func ListUsers(c *gin.Context) {
	const Beare_schema = "Bearer "
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatus(401)
	}

	token := header[len(Beare_schema):]

	claim, _ := services.NewJWTServices().ExtractClaims(token)

	if claim["is_adm"] != true {
		c.JSON(401, gin.H{
			"error": "Access denied!",
		})
		return
	}

	db := database.GetDatabase()
	var users []models.User

	err := db.Find(&users).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list users: " + err.Error(),
		})
		return
	}
	c.JSON(200, users)
}
