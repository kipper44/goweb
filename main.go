package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"newoo/WinningDerbyAdmin/utils"
	"newoo/WinningDerbyAdmin/winService"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func mysqlModule() {
	fmt.Println(utils.GetAdmin())
	utils.DB_Config().Init()

	var id, name string

	utils.DB_Config().GetAdmin().Select("SELECT id,name FROM account", func(rows *sql.Rows) {

		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, name)
		}
	})
}

func WebInit() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	utils.MysqlConfig().Init(dir + "/config/mysql_config.json")
	utils.ServerConfig().Init(dir + "/config/server_config.json")
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())

	stored := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("WinningDerbyAdminSession", stored))
	// elog.Error(1, "web init start2")
	// router.HTMLRender = ginview.Default()
	router.HTMLRender = ginview.New(
		goview.Config{
			Root:      dir + "/views",
			Extension: ".html",
			Master:    "layouts/master",
			Partials:  []string{},
			Funcs: template.FuncMap{
				"sub": func(a, b int) int {
					return a - b
				},
				"copy": func() string {
					return time.Now().Format("2006")
				},
			},
			DisableCache: true,
		})

	// router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.Static("/static", dir+"/static")
	varRouter.Init(router)
	varRouter.SetRoute()
	// router.Use(static.Serve("/static", static.LocalFile(dir+"/static", false)))
	// elog.Error(1, "web init start3")
	ip := utils.ServerConfig().JsonData["ip"].(string)
	port := utils.ServerConfig().JsonData["port"].(float64)
	iPort := int(port)
	// fmt.Println(reflect.TypeOf(port))
	// fmt.Println(utils.ServerConfig().JsonData)
	command := ip + ":"
	command += strconv.Itoa(iPort)
	// winService.Elog.Error(1, command)
	router.Run(command)
}

func main() {

	const svcName = "WinningDerbyAdmin"
	const svcDesc = "WinningDerbyAdmin Service"
	winService.MainLoop = WebInit
	winService.CommandParser(svcName, svcDesc)
}

// https://github.com/hhkbp2/go-logging
