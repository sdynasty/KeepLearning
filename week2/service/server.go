package service

import (
	"keepLearning/week2/dao"
)

type week2Service struct {

}

func NewWeek2Service() * week2Service {
	return new(week2Service)
}

// 获取玩家货币
func (s *week2Service) GetUserCoin(uid string) (int64, error) {
	c, err := dao.GetUserCoinById(uid)
	return c, err
}
