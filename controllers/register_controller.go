package controllers

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"
	"web-gin/databases"

	"github.com/gin-gonic/gin"
)

func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册"})
}

func RegisterPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println(username, password, repassword)

	// 验证用户名和密码
	if username == "" || password == "" || repassword == "" {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "用户名或密码不能为空"})
	}
	if password != repassword {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "两次密码不一致"})
	}

	//判断是否已经被注册
	_, err := databases.QueryUsername(username)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "用户名已存在"})
		return
	} else {
		password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
		user := databases.User{
			Username:   username,
			Password:   password,
			Status:     true,
			CreateTime: time.Now(),
		}
		_, err := databases.Insert(user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "注册失败"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "注册成功"})
		}
	}

}
