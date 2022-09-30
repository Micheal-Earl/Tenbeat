package handlers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/global"
)

func (h handler) Me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(global.Userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h handler) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
