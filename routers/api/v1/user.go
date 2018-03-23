package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/y1281010069/go-gin/pkg/e"
	"github.com/y1281010069/go-gin/models"
)

type user struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if username != "" {
		maps["username"] = username
	}
	if password != "" {
		maps["password"] = password
	}

	code := e.SUCCESS
	data["user_info"] = models.Login(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}