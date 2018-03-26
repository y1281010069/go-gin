package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/y1281010069/go-gin/pkg/e"
	"github.com/y1281010069/go-gin/models"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/y1281010069/go-gin/pkg/logging"
)

type user struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary 登录
// @Description 用户登录接口
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {string} json "{"code":200,"data":{"user_info":[{"id":1,"username":"test"}]},"msg":"ok"}"
// @Router /api/v1/login [get]
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	// json测试
	// 初始化数据
	input := make(map[string]interface{})
	input["status"] = 200

	// json格式化
	inputJson, _ := ffjson.Marshal(input)

	// 输出
	logging.Info(200, inputJson)
	// [200 [123 34 115 116 97 116 117 115 34 58 50 48 48 125]]
	logging.Info(200, string(inputJson[:]))
	// {"status":200}

	// json解析
	inputDecode := make(map[string]interface{})
	ffjson.Unmarshal(inputJson[:], &inputDecode)
	// 输出
	logging.Info(200, inputDecode)
	// map[status:200]

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