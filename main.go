package main

import (
	"graphql/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	// routing "/" ibarat di rest seperti group
	route.PUT("/", func(c *gin.Context) {
		type Query struct {
			Query string `json:"query"`
		}

		var query Query

		c.Bind(&query)
		result := models.ExecuteQuery(query.Query, models.Schema)

		c.JSON(http.StatusOK, result)
	})

	route.Run()
}
