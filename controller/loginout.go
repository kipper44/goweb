package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"newoo/WinningDerbyAdmin/utils"
	"strings"

	"newoo/WinningDerbyAdmin/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// const seesionTime = 120
const seesionTime = 60 * 10 // 10 min

type MenuChecker map[string]string

var cMenuChecker MenuChecker

func MenuCheckerNew() MenuChecker {
	if nil == cMenuChecker {
		cMenuChecker = make(map[string]string)
	}
	return cMenuChecker
}

func (r MenuChecker) AuthAdd(path string, memuName string) {
	prePath := "/auth"
	r.Add(prePath, path, memuName)
}
func (r MenuChecker) Add(prePath string, path string, memuName string) {
	r[prePath+path] = memuName
}

func (r MenuChecker) GetMenu(path string) string {

	for key, str := range cMenuChecker {
		fmt.Println(key, path)
		if key == path {
			return str
		}
	}
	return ""
}

func (r MenuChecker) check(id string, path string) bool {
	// MenuString := r.GetMenu(path)
	MenuString, _ := cMenuChecker[path]
	DB := utils.GetAdmin()
	defer DB.Close()
	admin := models.AdminAccount{}
	DB.Where("id=?", id).Find(&admin)
	// fmt.Println(admin)
	var menu []interface{}
	json.Unmarshal([]byte(admin.Permition), &menu)

	for _, val := range menu {
		data := val.(map[string]interface{})
		str := data["Menu"].(string)
		if strings.Compare(str, MenuString) == 0 {
			Permit := data["Permit"].(bool)
			return Permit
		}
	}

	return false
}

func RequireLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get("user")
	fmt.Println("RequireLogin =>", ctx.Request.URL.Path)
	if user == nil {
		// You'd normally redirect to login page
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		OnError(ctx, "Invalid session token")
	} else {
		chk := cMenuChecker.check(user.(string), ctx.Request.URL.Path)
		if chk == false {
			OnError(ctx, "Check Menu Auth ")
		} else {
			// Continue down the chain to handler etc
			sessions.Default(ctx).Options(sessions.Options{MaxAge: seesionTime})
			ctx.Next()
		}
	}
}

func LoginController(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get("user")
	if user != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/auth/user")
		return
	}

	ctx.HTML(http.StatusOK, "Login.html", gin.H{})
}

func OnError(ctx *gin.Context, Message string) {
	ctx.HTML(http.StatusOK, "error.html", gin.H{"Message": Message})
}

func PostLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	fmt.Println(mapParam)
	id := mapParam["id"].(string)
	pass := mapParam["pass"].(string)

	admin := models.AdminAccount{}
	DB := utils.GetAdmin()
	defer DB.Close()
	DB.Where("id=?", id).Find(&admin)
	fmt.Println(admin)

	if pass == admin.Password {
		session.Set("user", admin.Id) //In real world usage you'd set this to the users ID
		// session.Set("Expires", time.Now().Add(120*time.Second))
		session.Options(sessions.Options{
			// MaxAge: 3600 * 12, // 12hrs
			MaxAge: seesionTime, // 2min
		})
		err := session.Save()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
			panic("login error")
		} else {
			// ctx.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
			// ctx.Redirect(http.StatusMovedPermanently, "/auth/user")
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		// ctx.JSON(http.StatusOK, gin.H{"status": "error"})
		// panic("login error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	// ctx.Redirect(http.StatusMovedPermanently, "/user")
}

func LogOut(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get("user")
	if user == nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		ctx.Redirect(http.StatusMovedPermanently, "/")
	} else {
		log.Println(user)
		session.Delete("user")
		session.Save()
		// ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
		ctx.Redirect(http.StatusMovedPermanently, "/")
	}
}

func ToExcelController(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// user := session.Get("user")
	// if user == nil {
	// 	// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	// 	ctx.Redirect(http.StatusMovedPermanently, "/")
	// } else {
	// 	log.Println(user)
	// 	session.Delete("user")
	// 	session.Save()
	// 	// ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	// 	ctx.Redirect(http.StatusMovedPermanently, "/")
	// }
	// var mapParam map[string]interface{}
	// ctx.BindJSON(&mapParam)
	// fmt.Println(mapParam)

	buf := new(bytes.Buffer)
	buf.ReadFrom(ctx.Request.Body)
	str := buf.String()
	fmt.Println(str)
}
