package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newoo/WinningDerbyAdmin/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func StDiaController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "statistics/dia", gin.H{"Name": "Statistcs Dia", "Admin": loginUser})
}

func StCoinController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "statistics/coin", gin.H{"Name": "Statistcs Coin", "Admin": loginUser})
}

func StDauController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "statistics/dau", gin.H{"Name": "Statistcs DAU", "Admin": loginUser})
}

func StCCUController(ctx *gin.Context) {
	loginUser := GetAdminName(ctx)
	ctx.HTML(http.StatusOK, "statistics/ccu", gin.H{"Name": "Statistcs CCU", "Admin": loginUser})
}

type DiaSearchType struct {
	Dia55, Dia110, Dia220, Dia440, Dia660, Dia1100, Dia2200, Revenue int
	DATE                                                             time.Time
}

func STDaiSearch(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)

	start := mapParam["start"].(string)
	end := mapParam["end"].(string)
	DB := utils.Getlog()
	defer DB.Close()
	query := `   SELECT SUM(Dia55) AS Dia55,SUM(Dia110) AS Dia110,SUM(Dia220) AS Dia220,SUM(Dia440) AS Dia440,
	SUM(Dia660) AS Dia660,SUM(Dia1100) AS Dia1100,SUM(Dia2200) AS Dia2200,Revenue,event_date AS DATE  FROM (
    SELECT 
        IF( change_amount =55,ca_count, 0) AS Dia55
        ,IF( change_amount =110,ca_count, 0) AS Dia110
        ,IF( change_amount =220,ca_count, 0) AS Dia220
        ,IF( change_amount =440,ca_count, 0) AS Dia440
        ,IF( change_amount =660,ca_count, 0) AS Dia660
        ,IF( change_amount =1100,ca_count, 0) AS Dia1100
        ,IF( change_amount =2200,ca_count, 0) AS Dia2200
        ,IF( change_amount = 0,ca_count, 0) AS Revenue
        ,event_date
        FROM 
        (
        SELECT change_amount,COUNT(change_amount) AS ca_count , DATE(created_at) AS event_date FROM assets
        WHERE event_id IN(7,8) AND asset_id = 0 
        AND change_amount IN(55,110,220,440,660,1100,2200)
        AND created_at >= '%s' AND created_at <= '%s'
        GROUP BY DATE(created_at),change_amount
        ) z
        ) s
		GROUP BY event_date `

	response := fmt.Sprintf(query, start, end)
	// fmt.Print(response)
	rows, _ := DB.Raw(response).Rows()
	cols, _ := rows.Columns()
	// fmt.Println(cols)
	defer rows.Close()
	var stData []DiaSearchType
	for rows.Next() {
		row := DiaSearchType{}
		rows.Scan(&row.Dia55, &row.Dia110, &row.Dia220, &row.Dia440, &row.Dia660, &row.Dia1100, &row.Dia2200, &row.Revenue, &row.DATE)
		stData = append(stData, row)
	}
	data, _ := json.Marshal(stData)
	column, _ := json.Marshal(cols)

	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   string(data),
		"column": string(column),
	})

}

type CoinStSearchType struct {
	Coin900K   int
	Coin1_5m   int
	Coin3_09m  int
	Coin6_24m  int
	Coin15_75m int
	Coin31_8m  int
	Coin64_2m  int

	DATE time.Time
}

func stattistics_coin(startdate string, enddate string) ([]byte, []byte) {
	DB := utils.Getlog()
	defer DB.Close()
	strquery := `
    SELECT SUM(A) AS Coin900K,SUM(B) AS 'Coin1.5m',SUM(C) AS 'Coin3.09m',SUM(D) AS 'Coin6.24m',
	SUM(E) AS 'Coin15.75m',SUM(F) AS 'Coin31.8m',SUM(G) AS 'Coin64.2m',event_date AS DATE  FROM (
    SELECT 
    IF( change_amount =900000,ca_count, 0) AS A
    ,IF( change_amount =1500000,ca_count, 0) AS B
    ,IF( change_amount =3090000,ca_count, 0) AS C
    ,IF( change_amount =6240000,ca_count, 0) AS D
    ,IF( change_amount =15750000,ca_count, 0) AS E
    ,IF( change_amount =31800000,ca_count, 0) AS F
    ,IF( change_amount =64200000,ca_count, 0) AS G
    ,event_date
    FROM 
    (
    SELECT change_amount,COUNT(change_amount) AS ca_count , DATE(created_at) AS event_date FROM assets
    WHERE event_id IN(7,8) AND asset_id = 1 
    AND change_amount IN (900000,1500000,3090000,6240000,15750000,31800000,64200000)
    AND created_at >= '{%s}' AND created_at <= '{%s}'
    GROUP BY DATE(created_at),change_amount
    ) z
    ) s
    GROUP BY event_date`

	response := fmt.Sprintf(strquery, startdate, enddate)
	rows, _ := DB.Raw(response).Rows()
	cols, _ := rows.Columns()
	fmt.Println(cols)
	defer rows.Close()
	var stData []CoinStSearchType
	for rows.Next() {
		row := CoinStSearchType{}
		rows.Scan(&row.Coin900K, &row.Coin1_5m, &row.Coin3_09m, &row.Coin6_24m, &row.Coin15_75m, &row.Coin31_8m, &row.Coin64_2m, &row.DATE)
		stData = append(stData, row)
	}
	data, _ := json.Marshal(stData)
	column, _ := json.Marshal(cols)
	return data, column
}

type TotalBettinSearchType struct {
	Seoul      int
	Kentucky   int
	HappyValue int
	Tokyo      int
	Dainam     int
	DATE       time.Time
}

func statics_total_betting(startdate string, enddate string) ([]byte, []byte) {
	DB := utils.GetGame()
	defer DB.Close()
	strquery := `
    SELECT SUM(Seoul) AS Seoul,SUM(Kentucky) AS Kentucky,SUM(HappyValue) AS HappyValue, SUM(Tokyo) AS Tokyo,SUM(Dainam) AS Dainam,g_date AS DATE FROM 
    (
    SELECT 
    IF( stadium_id =1,Betting_Total, 0) AS Seoul
    ,IF( stadium_id =2,Betting_Total, 0) AS Kentucky
    ,IF( stadium_id =3,Betting_Total, 0) AS HappyValue
    ,IF( stadium_id =4,Betting_Total, 0) AS Tokyo
    ,IF( stadium_id =5,Betting_Total, 0) AS Dainam
    ,g_date
    FROM 
    (
    SELECT stadium_id , SUM(bet) AS Betting_Total ,g_date  FROM 
    (
        SELECT stadium_id,bet,DATE(bet_date) AS g_date FROM exacta_accounts
        WHERE bet_date >= '%s' AND bet_date <= '%s'
        UNION ALL 
        SELECT stadium_id,bet,DATE(bet_date) AS g_date FROM quinella_accounts
        WHERE bet_date >= '%s' AND bet_date <= '%s'
        UNION ALL 
        SELECT stadium_id,bet,DATE(bet_date) AS g_date FROM show_accounts
        WHERE bet_date >= '%s' AND bet_date <= '%s'
        UNION ALL 
        SELECT stadium_id,bet,DATE(bet_date) AS g_date FROM win_accounts
        WHERE bet_date >= '%s' AND bet_date <= '%s'
    )a
    GROUP BY DATE(a.g_date),a.stadium_id
    ORDER BY stadium_id,a.g_date
    ) z
    ) s
    GROUP BY g_date    
    `
	response := fmt.Sprintf(strquery, startdate, enddate, startdate, enddate, startdate, enddate, startdate, enddate)
	// fmt.Println(response)
	rows, _ := DB.Raw(response).Rows()
	cols, _ := rows.Columns()
	fmt.Println(cols)
	defer rows.Close()
	var stData []TotalBettinSearchType
	for rows.Next() {
		row := TotalBettinSearchType{}
		rows.Scan(&row.Seoul, &row.Kentucky, &row.HappyValue, &row.Tokyo, &row.Dainam, &row.DATE)
		stData = append(stData, row)
	}
	data, _ := json.Marshal(stData)
	column, _ := json.Marshal(cols)
	return data, column

}

func STCoinSearch(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)

	mode := ToInt(mapParam["mode"])
	start := mapParam["start"].(string)
	end := mapParam["end"].(string)
	var data, column []byte
	if 1 == mode {
		data, column = stattistics_coin(start, end)
	} else {
		data, column = statics_total_betting(start, end)
	}
	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   string(data),
		"column": string(column),
	})

}

func stattistics_newUser(returnData []DauSearchType, startdate string, enddate string) {
	DB := utils.GetGame()
	defer DB.Close()
	strquery := `
        SELECT COUNT(id),DATE(new_date) AS searchDate FROM users
        WHERE new_date >= '%s' AND new_date <= '%s'
        GROUP BY DATE(new_date)
    `
	response := fmt.Sprintf(strquery, startdate, enddate)
	// fmt.Println(response)
	rows, _ := DB.Raw(response).Rows()
	cols, _ := rows.Columns()
	fmt.Println(cols)
	defer rows.Close()
	for rows.Next() {
		var count int
		var cdate time.Time
		find := false
		rows.Scan(&count, &cdate)
		for i, val := range returnData {
			if cdate == val.Date {
				returnData[i].NewUser = count
				find = true
				break
			}
		}
		if false == find {
			add := DauSearchType{NewUser: count, Date: cdate}
			returnData = append(returnData, add)
		}
	}

}

type DauSearchType struct {
	DAU, NewUser, Revenue int
	Date                  time.Time
}

func STDAUSearch(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	DB := utils.Getlog()
	defer DB.Close()

	startdate := mapParam["start"].(string)
	enddate := mapParam["end"].(string)

	strquery := `
	SELECT COUNT(DISTINCT(a.user_no)) as DAU ,IFNULL(SUM(b.change_amount),0) as Revenue, DATE(a.eDate) AS Date FROM (
		SELECT  user_no , DATE(created_at) AS eDate FROM logins
		WHERE OPEN=b'1' AND created_at >= '%s' AND created_at <= '%s'
	) a 
	LEFT OUTER JOIN
	(
		SELECT change_amount,DATE(created_at) AS eDate FROM assets 
		WHERE asset_id = 0 AND change_amount > 0 AND created_at >= '%s' AND created_at <= '%s'
		) b
	ON a.eDate = b.eDate
	GROUP BY DATE(a.eDate)
	`
	response := fmt.Sprintf(strquery, startdate, enddate, startdate, enddate)
	// fmt.Println(response)
	rows, _ := DB.Raw(response).Rows()
	cols, _ := rows.Columns()
	cols = append(cols, "NewUser")
	fmt.Println(cols)
	defer rows.Close()
	var stData []DauSearchType
	revnuValue := ToInt(utils.ServerConfig().JsonData["revenue"])
	for rows.Next() {
		row := DauSearchType{}
		rows.Scan(&row.DAU, &row.Revenue, &row.Date)
		row.Revenue = row.Revenue * revnuValue
		stData = append(stData, row)
	}
	stattistics_newUser(stData, startdate, enddate)
	// fmt.Print(stData)
	data, _ := json.Marshal(stData)
	column, _ := json.Marshal(cols)
	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   string(data),
		"column": string(column),
	})

}

type CCULogType struct {
	Time time.Time
	CCU  int
}

func STCCUSearch(ctx *gin.Context) {
	var mapParam map[string]interface{}
	ctx.BindJSON(&mapParam)
	DB := utils.Getlog()
	defer DB.Close()

	startdate := mapParam["searchDate"].(string)
	// enddate := mapParam["end"].(string)

	strquery := `
	SELECT connection as CCU,created_at as Time FROM log.connections
	WHERE created_at >='%s 00:00:00' AND created_at <='%s 23:59:59'
	`
	response := fmt.Sprintf(strquery, startdate, startdate)
	fmt.Println(response)
	rows, _ := DB.Raw(response).Rows()
	cols, _ := rows.Columns()
	fmt.Println(cols)
	defer rows.Close()
	var stData []CCULogType
	for rows.Next() {
		row := CCULogType{}
		rows.Scan(&row.CCU, &row.Time)
		stData = append(stData, row)
	}
	data, _ := json.Marshal(stData)
	column, _ := json.Marshal(cols)
	ctx.JSON(http.StatusOK, gin.H{
		// "code":   http.StatusOK,
		"data":   string(data),
		"column": string(column),
	})

}
