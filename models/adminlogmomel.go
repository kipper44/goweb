package models

import "time"

type AdminMailToLog struct {
	LogNo     int64      `gorm:"column:no"`
	UserNo    int64      `gorm:"column:UserNo"`
	NickName  string     `gorm:"column:NickName"`
	AdminName string     `gorm:"column:AdminName"`
	Coin      int64      `gorm:"column:coin"`
	Dia       int        `gorm:"column:dia"`
	Mileage   int        `gorm:"column:mileage"`
	Reg_date  *time.Time `gorm:"column:reg_date"`
}

func (AdminMailToLog) TableName() string {
	return "mailtolog"
}

type GmAccount struct {
	Idx       int        `gorm:"column:idx"`
	Account   string     `gorm:"column:id"`
	Password  string     `gorm:"column:password"`
	Name      string     `gorm:"column:name"`
	Permition string     `gorm:"column:permition";sql:"type:text"`
	LastLogin *time.Time `gorm:"column:last_login"`
	Login     bool       `gorm:"column:login"`
}

func (GmAccount) TableName() string {
	return "account"
}

type AdminAccountCreateLog struct {
	No          int64     `gorm:"column:no"`
	AdminName   string    `gorm:"column:AdminName"`
	LogType     int       `gorm:"column:logType"`
	TargetAdmin string    `gorm:"column:targetAdmin"`
	Time        time.Time `gorm:"column:reg_date"`
}

func (AdminAccountCreateLog) TableName() string {
	return "createaccountlog"
}

type AdminMgrLogType int

const (
	Create AdminMgrLogType = 0
	Modify AdminMgrLogType = 1
	Remove AdminMgrLogType = 2
)

func (val AdminMgrLogType) Ordinal() int {
	return int(val)
}

type AdminLogType int

const (
	AssetChange AdminLogType = 0
	MailTo      AdminLogType = 1
	Admin       AdminLogType = 2
)

func (val AdminLogType) Ordinal() int {
	return int(val)
}

func (val AdminLogType) String() string {
	names := [...]string{
		"AssetChange",
		"MailTo",
		"Admin"}
	if val < AssetChange || val > Admin {
		return "Unknown"
	}
	return names[val]
}

func (val AdminLogType) FromString(str string) int {
	names := [...]string{
		"AssetChange",
		"MailTo",
		"Admin"}
	for idx, name := range names {
		if name == str {
			return idx
		}
	}
	return -1
}

type AdminAssetLog struct {
	LogNo       int64      `gorm:"column:no"`
	UserNo      int64      `gorm:"column:UserNo"`
	NickName    string     `gorm:"column:NickName"`
	AssetType   int        `gorm:"column:assetType"`
	Admin       string     `gorm:"column:AdminName"`
	oldAsset    int64      `gorm:"column:oldAsset"`
	ChangeAsset int64      `gorm:"column:newAsset"`
	Time        *time.Time `gorm:"column:reg_date"`
}

func (AdminAssetLog) TableName() string {
	return "assetlog"
}
