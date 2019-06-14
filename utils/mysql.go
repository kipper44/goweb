package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var cMysqlDb *WdMySql

func DB_Config() *WdMySql {
	if nil == cMysqlDb {
		cMysqlDb = new(WdMySql)
		return cMysqlDb
	}
	return cMysqlDb
}

type WdMySql struct {
	// admin *sql.DB
	// game  *sql.DB
	// log   *sql.DB
	admin, game, log string
}

func makeConnectString(name string) string {
	adminMap := cMysqlConfig.JsonData[name].(map[string]interface{})
	user := adminMap["user"]
	pass := adminMap["pass"]
	ip := adminMap["ip"]
	port := adminMap["port"]
	DB := adminMap["DB"]

	conString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True", user, pass, ip, port, DB)
	// fmt.Printf(conString)
	return conString
}

func (r *WdMySql) Init() {
	r.admin = makeConnectString("admin_tool")
	r.game = makeConnectString("game")
	r.log = makeConnectString("game_log")
}

// func (r *WdMySql) RunAdmin(strQuery string, callback func(rows *sql.Rows)) {
// 	fmt.Println(r.admin)
// 	db, err := sql.Open("mysql", r.admin)
// 	defer db.Close()
// 	if err != nil {
// 		fmt.Println("Connection Failed to Open")
// 		return
// 	}
// 	rows, err := db.Query(strQuery)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	callback(rows)
// 	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
// }

type mysql_DB struct {
	DB *sql.DB
}

func dbConnect(constring string) *mysql_DB {
	db, err := sql.Open("mysql", constring)
	// defer db.Close()
	if err != nil {
		fmt.Println("Connection Failed to Open")
		return nil
	}
	ret := new(mysql_DB)
	ret.DB = db
	return ret

}

func (r *WdMySql) GetAdmin() *mysql_DB {
	return dbConnect(r.admin)
}

func (r *WdMySql) GetGame() *mysql_DB {
	return dbConnect(r.game)
}
func (r *WdMySql) Getlog() *mysql_DB {
	return dbConnect(r.log)
}

func (r *mysql_DB) Select(strQuery string, callback func(rows *sql.Rows)) {
	rows, err := r.DB.Query(strQuery)
	defer func() {
		r.DB.Close()
		r = nil
	}()
	if err != nil {
		log.Fatal(err)
		return
	}
	callback(rows)
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
}
