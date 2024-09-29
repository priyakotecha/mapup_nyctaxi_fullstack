package controller

import (
	"net/http"
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
