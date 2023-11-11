package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string `json:"name"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Identity      string `json:"identity"`
	ClientId      string `json:"client_id"`
	ClientPort    string `json:"client_port"`
	LoginTime     uint64 `json:"login_time"`
	HeartbeatTime uint64 `json:"heartbeat_time"`
	LogOutTime    uint64 `json:"log_out_time"`
	IsLogOut      bool   `json:"is_log_out"`
	DeviceInfo    string `json:"device_info"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
