package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/snsilvam/partydise.git/models"
)

// services data
var OurCurrentServices = []models.OurServices{
	{ID: "1", Title: "Recreación y Shows tematicoS", DescriptionOfServices: "Recreación, animación de eventos infantiles y para adultos, shows de personajes, baby showers."},
	{ID: "2", Title: "Comidas", DescriptionOfServices: "Ofrecemos de perros calientes ó sandwich, centro de mesas comestibles, gaseosas o jugos de caja , tortas."},
	{ID: "3", Title: "Decoraciones", DescriptionOfServices: "Decoraciones tematicas para cualquierocasión o persojane para darlemagia a tu dia."},
	{ID: "4", Title: "Ambientacion", DescriptionOfServices: "Cabinas de sonido, show de luces, show de burbujas chiquitecas para ambientar tu fiesta esepcial."},
}

// getOurServices as JSON.
func getAllOurServices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, OurCurrentServices)
}

func main() {
	router := gin.Default()
	router.GET("/OurCurrentServices", getAllOurServices)

	router.Run("localhost:8080")
}
