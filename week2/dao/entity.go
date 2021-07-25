package dao



// 玩家数据
type UserModel struct {
	Uid    string `xorm:"not null pk"`
	Name   string //
	Coin   int64  //
}