package controller

import (
	"net/http"

	"golang-gin-gonic-studing/model"

	"github.com/gin-gonic/gin"
)

// GetUsers obtiene la lista de usuarios
// @Summary Obtiene la lista de usuarios
// @Description Retorna una lista de usuarios registrados
// @Tags Usuarios
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, model.Users)
}

// CreateUser crea un nuevo usuario
// @Summary Crea un usuario
// @Description Recibe un JSON con los datos del usuario y lo agrega a la lista
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param user body model.User true "Datos del usuario"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = len(model.Users) + 1
	model.Users = append(model.Users, user)
	c.JSON(http.StatusCreated, user)
}
