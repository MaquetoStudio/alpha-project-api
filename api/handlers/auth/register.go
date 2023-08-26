package auth

import (
	"net/http"

	"alpha-project-api/models"
	"alpha-project-api/pkg"
	"alpha-project-api/services"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	models.UserDTO
}

func RegisterHandler(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"errors": "invalid input"}) // You can translate the errors if needed
		return
	}

	ok, errors := pkg.ValidateStruct(input)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	user, err := services.NewUserService().CreateUser(input.UserDTO.ToEntity())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
