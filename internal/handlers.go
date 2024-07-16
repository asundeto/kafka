package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleMessages(c *gin.Context) {
	msg := c.Param("message")
	var newMessage Message
	newMessage.Message = msg
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&newMessage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	sendToKafka(newMessage.Message)
	c.JSON(http.StatusOK, newMessage)
}
