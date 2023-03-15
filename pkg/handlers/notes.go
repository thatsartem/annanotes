package handlers

import (
	"annanotes/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddNoteRequestBody struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID string `json:"user_id"`
}

func (h handler) AddNote(c *gin.Context) {
	body := AddNoteRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	if err := h.DB.Where("id = ?", body.UserID).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var note models.Note

	note.Title = body.Title
	note.Text = body.Text
	note.UserID = body.UserID

	if result := h.DB.Create(&note); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &note)
}

func (h handler) GetNotes(c *gin.Context) {
	var notes []models.Note

	if result := h.DB.Find(&notes); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &notes)
}

func (h handler) GetNote(c *gin.Context) {
	id := c.Param("id")

	var note models.Note

	if result := h.DB.First(&note, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &note)
}

type UpdateNoteRequestBody struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (h handler) UpdateNote(c *gin.Context) {
	id := c.Param("id")
	body := UpdateNoteRequestBody{}

	// получаем тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var note models.Note

	if result := h.DB.First(&note, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	note.Title = body.Title
	note.Text = body.Text

	h.DB.Save(&note)

	c.JSON(http.StatusOK, &note)
}

func (h handler) DeleteNote(c *gin.Context) {
	id := c.Param("id")

	var note models.Note

	if result := h.DB.First(&note, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	h.DB.Delete(&note)
	c.Status(http.StatusOK)
}
