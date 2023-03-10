package OurServices

import (
	"github.com/gin-gonic/gin"
	"github.com/snsilvam/partydise.git/models"
)

// getOurServices as JSON.
func GetAllOurServices(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, OurCurrentServices)
}

// postAddService addds an service from JSON received in the request body.
func PostAddService(c *gin.Context) {
	var newCurrentService models.OurServices

	//Call BindJson to bind the received JSON to new services.
	if err := c.BindJSON(&newCurrentService); err != nil {
		return
	}

	//Add the new service to the slice.
	/* OurCurrentServices = append(OurCurrentServices, newCurrentService)
	c.IndentedJSON(http.StatusCreated, newCurrentService) */
}
