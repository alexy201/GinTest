package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}
