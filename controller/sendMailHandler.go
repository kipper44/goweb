package controller

import (
	"fmt"
	"net/http"
	"newoo/WinningDerbyAdmin/models"
	"newoo/WinningDerbyAdmin/utils"

	"github.com/gin-gonic/gin"
)

func SendMailController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "sendmail", gin.H{"Name": "SendMail", "Admin": loginUser})
}

func writeMailToGameDb(UID int, contents map[string]interface{}) {
	DB := utils.GetGame()
	defer DB.Close()
	gold, dia, mileage := getItemCount(contents["itemSelect"].([]interface{}), contents["itemCount"].([]interface{}))
	fmt.Println(gold, dia, mileage)
	title := contents["title"].(string)
	mailText := contents["content"].(string)
	// contents['dueDate']
	newMail := models.MailBox{User_no: UID, Title: title, From: "", Coin: gold, Dia: dia, Mileage: mileage, Content: mailText, Deleted: nil}
	fmt.Print(newMail)
	DB.Create(&newMail)
}

func MailToLog(userNo int64, nickName string, gold int64, _dia int, _mileage int, adminUser string) {
	DB := utils.GetAdmin()
	defer DB.Close()
	var newLog = models.AdminMailToLog{UserNo: userNo, NickName: nickName, AdminName: adminUser, Coin: gold, Dia: _dia, Mileage: _mileage}
	DB.Create(&newLog)
}

func getItemCount(itemCodes []interface{}, itemCounts []interface{}) (int64, int, int) {
	var gold int64 = 0
	var dia, mileage int
	dia = 0
	mileage = 0
	for idx, val := range itemCodes {
		value := ToInt(val)
		if value == models.GOLD.Ordinal() {
			gold += ToInt64(itemCounts[idx])
		} else if value == models.DIA.Ordinal() {
			dia += ToInt(itemCounts[idx])
		} else if value == models.MILEAGE.Ordinal() {
			mileage += ToInt(itemCounts[idx])
		}
	}
	return gold, dia, mileage
}

func getAllUsersUserNo() []int {
	DB := utils.GetGame()
	defer DB.Close()
	var gameUsers []models.GameUserCustom
	var userNos []int
	DB.Select("no").Find(&gameUsers).Pluck("no", &userNos)
	return userNos
}

func sendToMailUsers(userList []interface{}, contents map[string]interface{}, adminUser string) {
	for _, userData := range userList {
		userJson := userData.(map[string]interface{})
		userNo := ToInt(userJson["UserNo"])
		nick := userJson["Nickname"].(string)
		writeMailToGameDb(userNo, contents)

		gold, dia, mileage := getItemCount(contents["itemSelect"].([]interface{}), contents["itemCount"].([]interface{}))
		MailToLog(int64(userNo), nick, gold, dia, mileage, adminUser)
	}
}

func sendToMailAllUser(contents map[string]interface{}, adminuser string) {
	ret := getAllUsersUserNo()
	for _, userNo := range ret {
		writeMailToGameDb(userNo, contents)
	}

	gold, dia, mileage := getItemCount(contents["itemSelect"].([]interface{}), contents["itemCount"].([]interface{}))
	MailToLog(0, "ALL", gold, dia, mileage, adminuser)

}

func SendMail(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	isAllUser := mapParam["allUser"]

	DB := utils.GetGame()
	defer DB.Close()
	var gameUsers []models.GameUserCustom

	// DB.Where("id = ?", searchKeyword).Find(&gameUsers)
	DB.Select("no").Find(&gameUsers)
	adminUser := "loginUser"
	if true == isAllUser {
		fmt.Println("All")
		contents := mapParam["data"].(map[string]interface{})
		sendToMailAllUser(contents, adminUser)
	} else {
		users := mapParam["users"].([]interface{})
		contents := mapParam["data"].(map[string]interface{})
		sendToMailUsers(users, contents, adminUser)
		// fmt.Println("USER")
	}

	// PayoutColumn := []string{"LogNo", "Stadium", "Date", "Round", "Type", "HorseNo1", "HorseNo2", "Ratio", "BettingMoney", "Reward", "ArrivalTime", "ReceiveRewardTime", "Deleted"}
	// column, _ := json.Marshal(PayoutColumn)
	// DB := utils.GetGame()
	// defer DB.Close()
	// var payOuts []models.PayoutCustom

	// DB.Where("UserNo = ?", userNo).Find(&payOuts)
	// // fmt.Println(assetLogs)
	// userData, _ := json.Marshal(payOuts)
	// // fmt.Println(err, string(userData))

	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data": "OK",
	})

}
