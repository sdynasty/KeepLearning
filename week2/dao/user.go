package dao

import (
	"database/sql"
	xerr "github.com/pkg/errors"
)

//// 获取玩家信息
//func GetUserInfoById(uid string) (user *UserModel, err error) {
//	has, err := getDBEngine().Table("user").Where("uid=?", uid).Get(user)
//	if err != nil {
//		err = xerr.Wrapf(err, "GetUserInfoById uid=%s", uid)
//	} else if !has {
//		err = fmt.Errorf("can`t get userInfo by uid:%s", uid)
//	}
//	return
//}

// 获取玩家金币
func GetUserCoinById(uid string) (int64, error) {
	var c int64
	err := getDB().QueryRow("select coin from users where uid=?", uid).Scan(&c)
	if err != nil {
		// 找不到就返回0，一般情况下也不影响逻辑
		if err == sql.ErrNoRows {
			return 0, nil
		}
		err = xerr.Wrapf(err, "GetUserCoinById uid=%s", uid)
	}
	return c, err
}