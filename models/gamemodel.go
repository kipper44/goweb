package models

import "time"

type GameUser struct {
	No                         int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Id                         string `gorm:"column:id";"type:varchar(128);unique_index"`
	Nickname                   string `gorm:"type:varchar(14);unique_index"`
	Terms_agreement            int
	Recent_connection          time.Time
	Recent_disconnection       time.Time
	Recent_recharging          time.Time
	Gold                       int64
	Dia                        int
	Silver                     int
	Ruby                       int
	Mileage                    int
	Accrued_prize_gold         int
	Total_spent_gold           int
	Win_prediction_number      int
	Win_betting_number         int
	Show_prediction_number     int
	Show_betting_number        int
	Quinella_prediction_number int
	Quinella_betting_number    int
	Exacta_betting_number      int
	Exacta_prediction_number   int
	Platform                   int
	Abandoned                  int
}

func (GameUser) TableName() string {
	return "users"
}

type GameUserCustom struct {
	UserNo                     int    `gorm:"column:no";"primary_key";"AUTO_INCREMENT"`
	UID                        string `gorm:"column:id";"type:varchar(128);unique_index"`
	Nickname                   string `gorm:"column:nickname";"type:varchar(14);unique_index"`
	Terms_agreement            int
	LastLogin                  time.Time `gorm:"column:recent_connection"`
	Recent_disconnection       time.Time
	Recent_recharging          time.Time
	Coin                       int64 `gorm:"column:gold"`
	Dia                        int   `gorm:"column:dia"`
	Silver                     int
	Ruby                       int
	Mileage                    int `gorm:"column:mileage"`
	Accrued_prize_gold         int
	Total_spent_gold           int
	Win_prediction_number      int
	Win_betting_number         int
	Show_prediction_number     int
	Show_betting_number        int
	Quinella_prediction_number int
	Quinella_betting_number    int
	Exacta_betting_number      int
	Exacta_prediction_number   int
	Platform                   int
	Abandoned                  int
}

func (GameUserCustom) TableName() string {
	return "users"
}

// MailsColumn := []string{"LogNo", "Title", "Content", "From", "Coin", "Dia", "Mileage", "ArrivalTime", "ReceiveRewardTime", "Deleted"}
type MailBox struct {
	LogNo             int64      `gorm:"column:no"`
	User_no           int        `gorm:"column:user_no"`
	Title             string     `gorm:"column:title"`
	From              string     `gorm:"column:writer"`
	Coin              int64      `gorm:"column:gold"`
	Dia               int        `gorm:"column:dia"`
	Mileage           int        `gorm:"column:mileage"`
	Content           string     `gorm:"column:contents"`
	ArrivalTime       time.Time  `gorm:"column:created_at"`
	ReceiveRewardTime *time.Time `gorm:"column:received_at"`
	Deleted           *time.Time `gorm:"column:deleted_at"`
}

func (MailBox) TableName() string {
	return "mails"
}

// PayoutColumn := []string{"LogNo", "Stadium", "Date", "Round", "Type", "HorseNo1", "HorseNo2", "Ratio", "BettingMoney", "Reward", "ArrivalTime", "ReceiveRewardTime", "Deleted"}

type Payout struct {
	LogNo             int64     `gorm:"column:no"`
	userNo            int       `gorm:"column:userNo"`
	Type              int       `gorm:"column:betting_type"`
	Round             int       `gorm:"column:race_round"`
	Reward            int64     `gorm:"column:gold"`
	BettingMoney      int64     `gorm:"column:bet_gold"`
	Ratio             float32   `gorm:"column:payout_ratio"`
	Date              time.Time `gorm:"column:race_date"`
	ReceiveRewardTime time.Time `gorm:"column:receive_date"`
	Deleted           time.Time `gorm:"column:delete_date"`
	Stadium           int       `gorm:"column:stadium"`
	HorseNo1          int       `gorm:"column:horseNo1"`
	HorseNo2          int       `gorm:"column:horseNo2"`
	raceNo            int       `gorm:"column:raceNo"`
}

func (Payout) TableName() string {
	return "payout"
}

type PayoutCustom struct {
	LogNo             int64     `gorm:"column:no"`
	userNo            int       `gorm:"column:userNo"`
	Type              int       `gorm:"column:betting_type"`
	Round             int       `gorm:"column:race_round"`
	Reward            int64     `gorm:"column:gold"`
	BettingMoney      int64     `gorm:"column:bet_gold"`
	Ratio             float32   `gorm:"column:payout_ratio"`
	Date              time.Time `gorm:"column:race_date"`
	ArrivalTime       time.Time `gorm:"column:race_date"`
	ReceiveRewardTime time.Time `gorm:"column:receive_date"`
	Deleted           time.Time `gorm:"column:delete_date"`
	Stadium           int       `gorm:"column:stadium"`
	HorseNo1          int       `gorm:"column:horseNo1"`
	HorseNo2          int       `gorm:"column:horseNo2"`
	raceNo            int       `gorm:"column:raceNo"`
}

func (PayoutCustom) TableName() string {
	return "payout"
}
