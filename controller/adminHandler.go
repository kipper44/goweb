package controller

import (
	"encoding/json"
	"math"
	"net/http"
	"newoo/WinningDerbyAdmin/models"
	"newoo/WinningDerbyAdmin/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AdminController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "admin/admin", gin.H{"Name": "Admin", "Admin": loginUser})
}

func AdminLogController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "admin/adminLog", gin.H{"Name": "AdminLog", "Admin": loginUser})
}

func adminCreateLog(_targetAdminUser string, _adminUser string) {
	DB := utils.GetAdmin()
	defer DB.Close()
	newLog := models.AdminAccountCreateLog{AdminName: _adminUser, LogType: models.Create.Ordinal(), TargetAdmin: _targetAdminUser}
	DB.Create(&newLog)
}

func adminModifyLog(_targetAdminUser string, _adminUser string) {
	DB := utils.GetAdmin()
	defer DB.Close()
	newLog := models.AdminAccountCreateLog{AdminName: _adminUser, LogType: models.Modify.Ordinal(), TargetAdmin: _targetAdminUser}
	DB.Create(&newLog)
}

func adminRemoveLog(_targetAdminUser string, _adminUser string) {
	DB := utils.GetAdmin()
	defer DB.Close()
	newLog := models.AdminAccountCreateLog{AdminName: _adminUser, LogType: models.Remove.Ordinal(), TargetAdmin: _targetAdminUser}
	DB.Create(&newLog)
}

func GetAccountList(ctx *gin.Context) {
	Column := []string{"Account", "Name", "LastLogin", "Modify"}
	column, _ := json.Marshal(Column)
	DB := utils.GetAdmin()
	defer DB.Close()
	var gmList []models.GmAccount
	DB.Find(&gmList)
	userData, _ := json.Marshal(gmList)
	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   string(userData),
		"column": string(column),
	})
}

func AdminCreateProcess(input map[string]interface{}, authority string, adminUser string) bool {

	// password = sha256_crypt.encrypt(input['Password'])
	passWd := input["Password"].(string)
	ID := input["ID"].(string)
	name := input["Name"].(string)
	DB := utils.GetAdmin()
	defer DB.Close()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passWd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(hashedPassword))
	newLog := models.GmAccount{Account: ID, Password: string(hashedPassword), Name: name, Permition: authority}
	DB.Create(&newLog)

	adminCreateLog(ID, adminUser)
	return true
}

func AdminUpdateProcess(input map[string]interface{}, authority string, adminUser string) bool {
	passWd := input["Password"].(string)
	ID := input["ID"].(string)
	name := input["Name"].(string)
	DB := utils.GetAdmin()
	defer DB.Close()
	if passWd != "" {
		DB.Model(&models.GmAccount{}).Where("id = ?", ID).Update(models.GmAccount{Name: name, Permition: authority, Password: passWd})
	} else {
		DB.Model(&models.GmAccount{}).Where("id = ?", ID).Update(models.GmAccount{Name: name, Permition: authority})
	}
	adminModifyLog(ID, adminUser)
	return true
}

func makeAuthString(auths []interface{}) string {
	var strReturn = "["
	for _, val := range auths {
		org := val.(map[string]interface{})
		data, _ := json.Marshal(org)
		strReturn += string(data) + ","
	}
	strReturn = strReturn[:len(strReturn)-1]
	strReturn += "]"
	return strReturn
}

func AccountApply(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	modified := mapParam["modify"]
	intput := mapParam["input"].(map[string]interface{})
	auth := mapParam["auth"].([]interface{})

	print(intput, auth)
	authString := makeAuthString(auth)
	adminUser := "loginUser"
	if false == modified {
		AdminCreateProcess(intput, authString, adminUser)
	} else {
		AdminUpdateProcess(intput, authString, adminUser)
	}

	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"path": "",
	})

}

func AccountRemove(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	id := mapParam["ID"].(string)
	adminUser := "loginUser"
	DB := utils.GetAdmin()
	defer DB.Close()

	DB.Delete(&models.GmAccount{Account: id})

	adminRemoveLog(id, adminUser)

	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data": "OK",
	})

}

func SearchAssetLog(page int, pageSize int, maxPage int) ([]byte, int, []byte) {
	AssetColumn := []string{"LogNo", "UserNo", "NickName", "AssetType", "Admin", "oldAsset", "ChangeAsset", "Time"}
	column, _ := json.Marshal(AssetColumn)
	DB := utils.GetAdmin()
	defer DB.Close()

	if 0 == maxPage {
		var count int
		DB.Model(&models.AdminAssetLog{}).Count(&count)
		maxPage = int(math.Ceil(float64(count) / float64(pageSize)))
	}
	var logs []models.AdminAssetLog
	if page == 1 {
		DB.Limit(pageSize).Find(&logs)
	} else {
		// DB.Limit(page * pageSize).Offset((page - 1) * pageSize).Find(&logs)
		DB.Offset((page-1)*pageSize + 1).Limit(pageSize).Find(&logs)
	}
	data, _ := json.Marshal(logs)

	return data, maxPage, column
}

func SearchMailToLog(page int, pageSize int, maxPage int) ([]byte, int, []byte) {

	AssetColumn := []string{"LogNo", "UserNo", "NickName", "Admin", "Coin", "Dia", "Mileage", "Time"}
	column, _ := json.Marshal(AssetColumn)
	DB := utils.GetAdmin()
	defer DB.Close()

	if 0 == maxPage {
		var count int
		DB.Model(&models.AdminMailToLog{}).Count(&count)
		maxPage = int(math.Ceil(float64(count) / float64(pageSize)))
	}
	var logs []models.AdminMailToLog
	if page == 1 {
		DB.Limit(pageSize).Find(&logs)
	} else {
		// DB.Limit(page * pageSize).Offset((page - 1) * pageSize).Find(&logs)
		DB.Offset((page-1)*pageSize + 1).Limit(pageSize).Find(&logs)
	}
	data, _ := json.Marshal(logs)

	return data, maxPage, column
}

func SearchAdminModifyLog(page int, pageSize int, maxPage int) ([]byte, int, []byte) {

	AssetColumn := []string{"LogNo", "AdminName", "LogType", "TargetAdmin", "Time"}
	column, _ := json.Marshal(AssetColumn)
	DB := utils.GetAdmin()
	defer DB.Close()

	if 0 == maxPage {
		var count int
		DB.Model(&models.AdminAccountCreateLog{}).Count(&count)
		maxPage = int(math.Ceil(float64(count) / float64(pageSize)))
	}
	var logs []models.AdminAccountCreateLog
	if page == 1 {
		DB.Limit(pageSize).Find(&logs)
	} else {
		// DB.Limit(page * pageSize).Offset((page - 1) * pageSize).Find(&logs)
		DB.Offset((page-1)*pageSize + 1).Limit(pageSize).Find(&logs)
	}
	data, _ := json.Marshal(logs)

	return data, maxPage, column
}

func AdminLogSearch(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	var serchLogTypes models.AdminLogType
	serchLog := mapParam["method"].(string)
	serchLogType := serchLogTypes.FromString(serchLog)

	page := ToInt(mapParam["page"])
	pageSize := ToInt(mapParam["pageSize"])
	maxPage := ToInt(mapParam["maxPage"])
	var data, column []byte

	if models.AssetChange.Ordinal() == serchLogType {
		data, maxPage, column = SearchAssetLog(page, pageSize, maxPage)
	} else if models.MailTo.Ordinal() == serchLogType {
		data, maxPage, column = SearchMailToLog(page, pageSize, maxPage)
	} else {
		data, maxPage, column = SearchAdminModifyLog(page, pageSize, maxPage)
	}

	// adminUser := "loginUser"
	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":    string(data),
		"column":  string(column),
		"maxPage": maxPage,
	})

}
