package main

/*
import (
	"fmt"

	"github.com/snsilvam/partydise.git/models"
)

// services data
var OurCurrentServices = []models.OurServices{
	{ID: "1", Title: "Recreación y Shows tematicoS", DescriptionOfServices: "Recreación, animación de eventos infantiles y para adultos, shows de personajes, baby showers."},
	{ID: "2", Title: "Comidas", DescriptionOfServices: "Ofrecemos de perros calientes ó sandwich, centro de mesas comestibles, gaseosas o jugos de caja , tortas."},
	{ID: "3", Title: "Decoraciones", DescriptionOfServices: "Decoraciones tematicas para cualquierocasión o persojane para darlemagia a tu dia."},
	{ID: "4", Title: "Ambientacion", DescriptionOfServices: "Cabinas de sonido, show de luces, show de burbujas chiquitecas para ambientar tu fiesta esepcial."},
}

func main() {
	fmt.Println("Hello, now this project is uplouding. 🛩️")
}
*/

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "partydise-378916")
	if err != nil {
		// TODO: Handle error.
		fmt.Println("Error: ", err)
	}
	OurServices := client.Collection("OurServices")
	ny := OurServices.Doc("mLAl4C8lXSmjMpmxQLnK")
	docsnap, err := ny.Get(ctx)
	if err != nil {
		// TODO: Handle error.
		fmt.Println("Error: ", err)
	}
	dataMap := docsnap.Data()
	fmt.Println(dataMap)
}
