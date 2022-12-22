package main

import (
	"net/http"

	sources "kindlescrapper/structures"

	"github.com/gin-gonic/gin"
)

var (
	albums = []sources.Content{
		{ID: "1", Title: "Blue Train", URL: "John Coltrane", Origin: 56.99},
		{ID: "2", Title: "Jeru", URL: "Gerry Mulligan", Origin: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", URL: "Sarah Vaughan", Origin: 39.99},
	}
)

// homePage
func homePage(c *gin.Context) {
	c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "index.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
          "title": "Home Page",
      },
  )
}

func getContents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func saveContent(c *gin.Context) {
	var newAlbum sources.Content

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
