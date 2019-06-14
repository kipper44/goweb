package models

import (
	"time"
)

type AdminAccount struct {
	Idx        int32  `gorm:"column:idx";"primary_key";"AUTO_INCREMENT"`
	Id         string `gorm:"column:id";"type:varchar(45);unique_index"`
	Password   string `gorm:"column:password";"type:varchar(256)"`
	Name       string `gorm:"column:name";"type:varchar(50)"`
	Permition  string `sql:"type:text"`
	Last_login time.Time
	Login      bool `sql:"type:boolean"`
}

func (AdminAccount) TableName() string {
	return "account"
}

// func (r *AdminAccount) MarshalJSON() ([]byte, error) {
// 	arr := []interface{}{r.idx, r.id, r.password, r.name, r.permition, r.last_login, r.login}
// 	return json.Marshal(arr)
// }

// func (r *AdminAccount) UnmarshalJSON(bs []byte) error {
// 	arr := []interface{}{}
// 	json.Unmarshal(bs, &arr)
// 	// TODO: add error handling here.
// 	// r.idx = arr[0].(int64)
// 	r.id = arr[1].(string)
// 	r.password = arr[2].(string)
// 	r.name = arr[3].(string)
// 	r.permition = arr[4].(string)
// 	//r.last_login = arr[5].(time.Time)
// 	//r.login = arr[6].(int8)
// 	return nil
// }
