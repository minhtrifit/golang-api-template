package albums

import (
	"fmt"
	"go-api/datatype"
	"go-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []datatype.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(context *gin.Context) {
	// Return to client
    context.JSON(http.StatusOK, albums)
}

func GetAlbumById(context *gin.Context) {
	id := context.Param("id")

	fmt.Println("Album id: ", id);

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            context.IndentedJSON(http.StatusOK, a)
            return
        }
    }
	
    context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "album not found"})
}

func AddAlbums(context *gin.Context) {
	var newAlbum datatype.Album;

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := context.BindJSON(&newAlbum); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong album format"})    
    } else {
		// Add the new album to the slice.
		albums = append(albums, newAlbum)

		// Return to client
		context.IndentedJSON(http.StatusCreated, newAlbum)
	}

	// {
	// 	"id": "4",
	// 	"title": "The Modern Sound of Betty Carter",
	// 	"artist": "Betty Carter",
	// 	"price": 49.99
	// }
}

func DeleteAlbum(context *gin.Context) {
	var album datatype.AlbumTarget;

	if err := context.BindJSON(&album); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong album format"});
	} else {
		fmt.Println("CHECK:", album.ID);
		newAlbums := []datatype.Album{};

		for _, a := range albums {
			if a.ID == album.ID {
				newAlbums = helper.DeleteSlice(albums, 1);
				albums = newAlbums;
				context.JSON(http.StatusOK, gin.H{"msg": "Delete successfully"});
				return
			}
		}

		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})	
	}
}
