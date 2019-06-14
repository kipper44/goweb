package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newoo/WinningDerbyAdmin/models"
	"newoo/WinningDerbyAdmin/utils"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetAdminName(r *gin.Context) string {
	session := sessions.Default(r)
	user := session.Get("user")
	if user == nil {
		return "UnKnown"
	}

	return user.(string)

}

func UserController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "user/user", gin.H{"Name": "User", "Admin": loginUser})
}

func PurchaseController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "user/purchase", gin.H{"Name": "Puchase", "Admin": loginUser})
}

func CoinController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "user/coin", gin.H{"Name": "Coin", "Admin": loginUser})
}

func DiaController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "user/dia", gin.H{"Name": "Dia", "Admin": loginUser})
}

func MileageController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "user/mileage", gin.H{"Name": "Mileage", "Admin": loginUser})
}
func MailboxController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "user/mailbox", gin.H{"Name": "MailBox", "Admin": loginUser})
}

func PayoutController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "user/payout", gin.H{"Name": "PayOut", "Admin": loginUser})
}

func searchNickName(searchText string) (string, string) {
	UsersColumn := []string{"UserNo", "Nickname", "UID", "LastLogin", "Coin", "Dia", "Mileage"}
	column, err := json.Marshal(UsersColumn)

	DB := utils.GetGame()
	defer DB.Close()
	var gameUsers []models.GameUserCustom

	// DB.Where("id = ?", searchKeyword).Find(&gameUsers)
	DB.Where("nickname LIKE ?", "%"+searchText+"%").Find(&gameUsers)
	fmt.Println(gameUsers)
	userData, err := json.Marshal(gameUsers)
	fmt.Println(err)

	return string(userData), string(column)
}

func searchUID(searchText string) (string, string) {
	UsersColumn := []string{"UserNo", "Nickname", "UID", "LastLogin", "Coin", "Dia", "Mileage"}
	column, err := json.Marshal(UsersColumn)
	DB := utils.GetGame()
	defer DB.Close()
	var gameUsers []models.GameUserCustom

	// DB.Where("id = ?", searchKeyword).Find(&gameUsers)
	DB.Where("id = ?", searchText).Find(&gameUsers)
	fmt.Println(gameUsers)
	userData, err := json.Marshal(gameUsers)
	fmt.Println(err)

	return string(userData), string(column)
}

func UserSearch(ctx *gin.Context) {

	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	fmt.Println(mapParam)

	method := mapParam["method"].(string)
	searchKeyword := mapParam["searchText"].(string)
	fmt.Println(method, searchKeyword)

	var returnData, returnColumn string

	switch method {
	case "Nickname":
		returnData, returnColumn = searchNickName(searchKeyword)
	case "UID":
		returnData, returnColumn = searchUID(searchKeyword)
	default:
		fmt.Println("search Keyword error")
	}

	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   returnData,
		"column": returnColumn,
	})
}
func ToInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return value.(int)
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		return int(value.(float64))
	case string:
		data := value.(string)
		ret, err := strconv.Atoi(data)
		fmt.Println(err)
		return ret
	default:
		fmt.Println(v)
		return 0
	}
	return 0
}

func ToInt64(value interface{}) int64 {
	switch v := value.(type) {
	case int:
		return int64(value.(int))
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		return int64(value.(float64))
	case string:
		data := value.(string)
		ret, err := strconv.ParseInt(data, 10, 64)
		fmt.Println(err)
		return ret
	default:
		fmt.Println(v)
		return 0
	}
	return 0
}

func UserDataSave(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)

	saveData := mapParam["data"].([]interface{})
	orgData := mapParam["pre"].([]interface{})

	for idx, org := range orgData {
		fmt.Println(idx)
		orgdata := org.(map[string]interface{})
		userNo := ToInt64(orgdata["UserNo"])
		var gold int64
		var dia, mileage int
		search := false
		for i, v := range saveData {
			fmt.Println(i)
			data := v.(map[string]interface{})
			userNo2 := ToInt64(data["UserNo"])
			if userNo == userNo2 {
				gold = ToInt64(data["Coin"])
				dia = ToInt(data["Dia"])
				mileage = ToInt(data["Mileage"])
				search = true
				break
			}
		}
		if search == true {
			fmt.Println(gold, dia, mileage)
			DB := utils.GetGame()
			defer DB.Close()
			a := DB.Model(&models.GameUserCustom{}).Where("no = ?", userNo).Update(models.GameUserCustom{Coin: gold, Dia: dia, Mileage: mileage})
			fmt.Println(a)
		}
	}
	// fmt.Println(reflect.TypeOf(saveData))
	// fmt.Println(orgData)

	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data": "OK",
	})

}

func PurchaseDetail(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	// fmt.Println(mapParam)
	userNo := mapParam["user_no"]
	fmt.Println(userNo)

	UsersColumn := []string{"UserNo", "Event", "Asset", "change_amount", "Result_amount", "Time"}
	column, err := json.Marshal(UsersColumn)
	fmt.Println(err, column)
	DB := utils.Getlog()
	defer DB.Close()
	var assetLogs []models.AssetLogCustom

	DB.Where("user_no = ? AND  event_id =?", userNo, models.PURCHASE_IN_CASH_SHOP).Find(&assetLogs)
	fmt.Println(assetLogs)
	userData, err := json.Marshal(assetLogs)
	fmt.Println(err, string(userData))
	var targetData []interface{}
	json.Unmarshal(userData, &targetData)
	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   IFMarshal(targetData),
		"column": string(column),
	})

}

func CoinDetail(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	// fmt.Println(mapParam)
	userNo := mapParam["user_no"]
	fmt.Println(userNo)

	UsersColumn := []string{"UserNo", "Event", "Asset", "change_amount", "Result_amount", "Time"}
	column, _ := json.Marshal(UsersColumn)
	// fmt.Println(err, column)
	DB := utils.Getlog()
	defer DB.Close()
	var assetLogs []models.AssetLogCustom

	DB.Where("user_no = ? AND  asset_id =?", userNo, models.GOLD).Find(&assetLogs)
	// fmt.Println(assetLogs)
	userData, _ := json.Marshal(assetLogs)
	// fmt.Println(err, string(userData))

	var targetData []interface{}
	json.Unmarshal(userData, &targetData)
	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   IFMarshal(targetData),
		"column": string(column),
	})

}

func IFMarshal(value []interface{}) string {
	var tmpString string = "["
	if len(value) > 0 {
		for _, v := range value {
			org := v.(map[string]interface{})
			var eventID models.EventNo
			eventID = models.EventNo(ToInt(org["Event"]))
			strEventID := strconv.Itoa(eventID.Ordinal())
			org["Event"] = eventID.String() + "(" + strEventID + ")"

			var assetID models.AssetNo
			assetID = models.AssetNo(ToInt(org["Asset"]))
			org["Asset"] = assetID.String() + "(" + strconv.Itoa(assetID.Ordinal()) + ")"

			data, _ := json.Marshal(org)
			tmpString += string(data) + ","
		}
		tmpString = tmpString[:len(tmpString)-1]
	}
	tmpString += "]"
	return tmpString
}

func DiaDetail(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	// fmt.Println(mapParam)
	userNo := mapParam["user_no"]
	fmt.Println(userNo)

	UsersColumn := []string{"UserNo", "Event", "Asset", "change_amount", "Result_amount", "Time"}
	column, _ := json.Marshal(UsersColumn)
	// fmt.Println(err, column)
	DB := utils.Getlog()
	defer DB.Close()
	var assetLogs []models.AssetLogCustom

	DB.Where("user_no = ? AND  asset_id =?", userNo, models.DIA).Find(&assetLogs)
	// fmt.Println(assetLogs)
	userData, _ := json.Marshal(assetLogs)
	// fmt.Println(err, string(userData))

	var targetData []interface{}
	json.Unmarshal(userData, &targetData)
	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   IFMarshal(targetData),
		"column": string(column),
	})

}

func MileageDetail(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	// fmt.Println(mapParam)
	userNo := mapParam["user_no"]
	fmt.Println(userNo)

	UsersColumn := []string{"UserNo", "Event", "Asset", "change_amount", "Result_amount", "Time"}
	column, _ := json.Marshal(UsersColumn)
	// fmt.Println(err, column)
	DB := utils.Getlog()
	defer DB.Close()
	var assetLogs []models.AssetLogCustom

	DB.Where("user_no = ? AND  asset_id =?", userNo, models.MILEAGE).Find(&assetLogs)
	// fmt.Println(assetLogs)
	userData, _ := json.Marshal(assetLogs)
	// fmt.Println(err, string(userData))

	var targetData []interface{}
	json.Unmarshal(userData, &targetData)
	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   IFMarshal(targetData),
		"column": string(column),
	})

}

func MailboxDetail(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	userNo := mapParam["user_no"]
	MailsColumn := []string{"LogNo", "Title", "Content", "From", "Coin", "Dia", "Mileage", "ArrivalTime", "ReceiveRewardTime", "Deleted"}
	column, _ := json.Marshal(MailsColumn)
	DB := utils.GetGame()
	defer DB.Close()
	var mailBoxs []models.MailBox

	DB.Where("user_no = ?", userNo).Find(&mailBoxs)
	// fmt.Println(assetLogs)
	userData, _ := json.Marshal(mailBoxs)
	// fmt.Println(err, string(userData))

	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   string(userData),
		"column": string(column),
	})

}

func PayoutDetail(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	userNo := mapParam["user_no"]
	PayoutColumn := []string{"LogNo", "Stadium", "Date", "Round", "Type", "HorseNo1", "HorseNo2", "Ratio", "BettingMoney", "Reward", "ArrivalTime", "ReceiveRewardTime", "Deleted"}
	column, _ := json.Marshal(PayoutColumn)
	DB := utils.GetGame()
	defer DB.Close()
	var payOuts []models.PayoutCustom

	DB.Where("UserNo = ?", userNo).Find(&payOuts)
	// fmt.Println(assetLogs)
	userData, _ := json.Marshal(payOuts)
	// fmt.Println(err, string(userData))

	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   string(userData),
		"column": string(column),
	})

}
