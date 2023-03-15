package handlers

import (
	"annanotes/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddUserRequestBody struct {
	Username string `json:"username"`
}

func (h handler) AddUser(c *gin.Context) {
	body := AddUserRequestBody{}

	// получаем тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.Username = body.Username

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func (h handler) GetUsers(c *gin.Context) {
	var users []models.User
	if result := h.DB.Preload("Notes").Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h handler) GetUserNotesById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if result := h.DB.Preload("Notes").First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	notes := user.Notes

	c.JSON(http.StatusOK, &notes)
}

func (h handler) GetUserNotes(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if result := h.DB.Preload("Notes").First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	notes := user.Notes

	c.JSON(http.StatusOK, &notes)
}
