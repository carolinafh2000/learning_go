package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// albun Estructura de albun
type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// Lista de albuns
var albums = []album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "21", Artist: "Adele", Year: 2011},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2002},
	{ID: "4", Title: "Meteora", Artist: "Linkin Park", Year: 2003},
	{ID: "5", Title: "25", Artist: "Adele", Year: 2015},
}

func getAlbums(c *gin.Context) {
	// Obtener la dirección IP local de la máquina
	fmt.Println("IP local de la máquina:", GetLocalIp())

	// Responder con los datos de los álbumes
	c.IndentedJSON(http.StatusOK, GetLocalIp())
}

func GetLocalIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

// Agregar un albums
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Obtener po id un albuns
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Albun no encontrado"})
}

// Obtener po id un albuns
func isHealthy(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "ping")
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/", isHealthy)

	router.Run(":8080")
}
