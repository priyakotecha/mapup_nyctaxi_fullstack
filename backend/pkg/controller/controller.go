package controller

import (
	"log"
	"net/http"
	"nyctaxi_mapup/pkg/cache"
	"nyctaxi_mapup/pkg/middleware"
	"nyctaxi_mapup/pkg/queue"
	"nyctaxi_mapup/pkg/repository"

	"github.com/gin-gonic/gin"
)

func GetTaxiData(c *gin.Context) {
	taxiData, err := repository.GetTaxiData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, taxiData)
}

func Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Use the getter to retrieve credentials
	user, exists := cache.GetUserCredentials(loginRequest.Username)
	if !exists || user.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// Generate JWT token and cache it
	token, err := middleware.GenerateToken(loginRequest.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "role": user.Role})
}

func IngestData(c *gin.Context) {
	queue.InitRedis()
	err := queue.AddToQueue("process_parquet_file")
	if err != nil {
		log.Panic("error while adding job to queue: ", err)

	}

	err = queue.ProcessQueue()
	if err != nil {
		log.Panic("error while procesing queue: ", err)
	}

	c.JSON(http.StatusCreated, "ok")
}
