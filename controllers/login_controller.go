package controllers

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"web-gin/databases"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录"})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)

	user, err := databases.QueryAll(username)
	if err != nil && !user.Status {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "账户错误"})
		return
	}
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))

	if user.Password != password {
		fmt.Println(password)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "密码错误"})
		return
	}
	if user.Id > 0 {
		session := sessions.Default(c)
		session.Set("loginuser", username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}

}

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginuser := session.Get("loginuser")
	fmt.Println(loginuser)
	if loginuser != nil {
		return true
	} else {
		return false
	}
}

func HomeGet(c *gin.Context) {
	islogin := GetSession(c)
	c.HTML(http.StatusOK, "home.html", gin.H{"islogin": islogin})
}

func ExitGet(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("loginuser")
	session.Save()
	fmt.Println("delete session...", session.Get("loginuser"))
	c.Redirect(http.StatusMovedPermanently, "/")
}
