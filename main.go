package main

import (
	albums "go-api/api"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.GET("/albums", albums.GetAlbums)
	router.GET("/album/:id", albums.GetAlbumById)
	router.POST("/add", albums.AddAlbums)
	router.POST("/delete", albums.DeleteAlbum)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Run server successfully"})    
	})

	router.Run("localhost:5500")
}