package week2

import (
	"fmt"
	"keepLearning/week2/dao"
	"keepLearning/week2/service"
	"testing"
)

// week2 error 学习作业
func TestDemo(t *testing.T) {
	dao.InitMysql()
	srv := service.NewWeek2Service()

	var uid = "u01"
	coin, err := srv.GetUserCoin(uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("get user coin success. uid:%s, coin:%d\n", uid, coin)
}