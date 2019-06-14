package models

import "time"

type EventNo int

const (
	BET_ON_WIN            EventNo = 0
	BET_ON_SHOW           EventNo = 1
	BET_ON_QUINELLA       EventNo = 2
	BET_ON_EXACTA         EventNo = 3
	RECEIVING_MAIL        EventNo = 4
	RECEIVING_PAYOUT      EventNo = 5
	PERIODICAL_RECHARGE   EventNo = 6
	PURCHASE_IN_CASH_SHOP EventNo = 7
	PURCHASE_IN_GAME_SHOP EventNo = 8
)

func (val EventNo) Ordinal() int {
	return int(val)
}
func (val EventNo) String() string {
	names := [...]string{
		"BET_ON_WIN",
		"BET_ON_SHOW",
		"BET_ON_QUINELLA",
		"BET_ON_EXACTA",
		"RECEIVING_MAIL",
		"RECEIVING_PAYOUT",
		"PERIODICAL_RECHARGE",
		"PURCHASE_IN_CASH_SHOP",
		"PURCHASE_IN_GAME_SHOP"}
	if val < BET_ON_WIN || val > PURCHASE_IN_GAME_SHOP {
		return "Unknown"
	}
	return names[val]
}

type AssetNo int

const (
	DIA     AssetNo = 0
	GOLD    AssetNo = 1
	MILEAGE AssetNo = 2
)

func (val AssetNo) Ordinal() int {
	return int(val)
}

func (val AssetNo) String() string {
	names := [...]string{
		"DIA",
		"GOLD",
		"MILEAGE"}
	if val < DIA || val > MILEAGE {
		return "Unknown"
	}
	return names[val]
}

type AssetLog struct {
	User_no       int `gorm:"primary_key";"AUTO_INCREMENT"`
	event_id      int
	asset_id      int
	change_amount int64
	result_amount int64
	created_at    time.Time
}

func (AssetLog) TableName() string {
	return "assets"
}

// LogColumn = ['UserNo','Event','Asset','change_amount','Result_amount','Time']

type AssetLogCustom struct {
	UserNo        int       `gorm:"column:user_no";"primary_key";"AUTO_INCREMENT"`
	Event         int       `gorm:"column:event_id"`
	Asset         int       `gorm:"column:asset_id"`
	change_amount int64     `gorm:"column:change_amount"`
	Result_amount int64     `gorm:"column:result_amount"`
	Time          time.Time `gorm:"column:created_at"`
}

func (AssetLogCustom) TableName() string {
	return "assets"
}
