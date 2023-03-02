package service

import (
	"fmt"
	"gin_chat/models"
	"gin_chat/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} data	json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 创建用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} data	json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	user.PassWord = c.Query("password")
	repassword := c.Query("repassword")

	data, _ := models.FindUserByName(user.Name)
	if user.Name == data.Name {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "该姓名已被注册!",
		})
		return
	}
	if user.PassWord != repassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "确认密码不一致!",
		})
		return
	}
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	user.PassWord = utils.MakePassword(user.PassWord, user.Salt) //密码加密

	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "创建用户成功！",
	})
}

// FindUserByNameAndPassword
// @Summary 用户登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} data	json{"code","message"}
// @Router /user/findUserByNameAndPassword [get]
func FindUserByNameAndPassword(c *gin.Context) {
	data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	user, err := models.FindUserByName(name)
	if err != nil {
		fmt.Println("models.FindUserByName err:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "账户不正确",
		})
		return
	}
	if !utils.ValidPassword(password, user.Salt, user.PassWord) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码不正确",
		})
		return
	}
	data = models.FindUserByNameAndPassword(name, utils.MakePassword(password, user.Salt))
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "ID"
// @Success 200 {string} data	json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)

	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "删除用户成功！",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		fmt.Println("govalidator.ValidateStruct err:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "修改参数不匹配！",
		})
		return
	}
	models.UpdateUser(user)

	c.JSON(http.StatusOK, gin.H{
		"message": "修改用户成功！",
	})
}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upGrade.Upgrade err:", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println("ws.Close err:", err)
		}
	}(ws)
	MsgHandler(ws, c)

}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println("utils.Subscribe err:", err)
			return
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println("ws.WriteMessage err:", err)

		}
		return
	}
}
