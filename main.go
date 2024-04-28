package main

import (
	"go-plus-gin/database"
	"go-plus-gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	Title  string  `binding:"required" json:"title"`
	Artist string  `binding:"required" json:"artist"`
	Price  float64 `binding:"required" json:"price"`
}

// var albums = []Album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	database.Connect()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	var albums []model.Album
	database.DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	var album model.Album
	result := database.DB.First(&album, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func postAlbum(c *gin.Context) {
	var newAlbum Album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	album := model.Album{Title: newAlbum.Title, Artist: newAlbum.Artist, Price: newAlbum.Price}
	result := database.DB.Create(&album)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}

	c.IndentedJSON(http.StatusCreated, album)
}
