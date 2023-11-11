package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Identity      string    `json:"identity"`
	ClientId      string    `json:"client_id"`
	ClientPort    string    `json:"client_port"`
	LoginTime     time.Time `json:"login_time"`
	HeartbeatTime time.Time `json:"heartbeat_time"`
	LoginOutTime  time.Time `json:"login_out_time"`
	IsLogOut      bool      `json:"is_log_out"`
	DeviceInfo    string    `json:"device_info"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
