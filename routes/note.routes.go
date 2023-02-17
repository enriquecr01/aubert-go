package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/enriquecr01/aubert-go/db"
	"github.com/enriquecr01/aubert-go/models"
)

func GetAllNotes(c *gin.Context) {
	var notes []models.Note

	userId := c.Param("userid")
	where := "user_id = " + userId
	db.DB.Select("id", "title", "note", "color").Omit("CreatedAt", "UpdatedAt", "DeletedAt").Where(where).Find(&notes)

	c.JSON(200, gin.H{
		"notes": &notes,
	})
}

func SearchNotes(c *gin.Context) {
	var notes []models.Note

	userId := c.Param("userid")
	term := c.Param("term")
	status := 0

	where := "user_id = " + userId + " AND title LIKE '%" + term + "%' " + " OR note LIKE '%" + term + "%' "
	db.DB.Where(where).Find(&notes)

	c.JSON(200, gin.H{
		"status": status,
		"notes":  &notes,
	})
}

func AddNote(c *gin.Context) {

	title := c.PostForm("title")
	note := c.PostForm("note")
	color := c.PostForm("color")
	userId, err := strconv.Atoi(c.PostForm("userId"))
	message := "Registered Succesfully"
	status := 0

	if err != nil {
		message = "Error during parsing"
		status = 1
	} else {
		newNote := models.Note{Title: title, Note: note, Color: color, UserId: userId}
		result := db.DB.Omit("updated_at", "deleted_at").Create(&newNote)

		if result.Error != nil {
			status = 1
			message = "Error during insert"
		}
	}

	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
	})
}

func ModifyNote(c *gin.Context) {

	title := c.PostForm("title")
	note := c.PostForm("note")
	color := c.PostForm("color")
	noteId, err := strconv.Atoi(c.PostForm("noteId"))
	message := "Updated Succesfully"
	status := 0

	if err != nil {
		message = "Error during parsing"
		status = 1
	} else {
		updatedSecret := db.DB.Where("id = ?", noteId).Updates(models.Note{Title: title, Note: note, Color: color})

		fmt.Printf("%+v\n", updatedSecret)
	}

	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
	})
}

func DeleteNote(c *gin.Context) {

	noteId := c.Param("noteId")
	message := "Deleted Succesfully"
	status := 0

	deletedNote := db.DB.Delete(&models.Note{}, noteId)

	if deletedNote.Error != nil {
		message = "Error deleting"
		status = 1
	}

	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
	})
}
