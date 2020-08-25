package model

import (
    "github.com/jinzhu/gorm"
    "echo-framework/lib/db"
    "echo-framework/lib/helper"
)

//设备日志表
type DeviceLog struct {
    Id            uint64 `json:"id"`
    DeviceId      uint64 `gorm:"size:64;not null;default:0;" json:"device_id"` //设备 id
    ServerAddress string `gorm:"size:32;not null;default:'';" json:"server_address"`
    ClientId      string `gorm:"size:32;not null;default:'';" json:"client_id"`

    LogUrl   string `gorm:"size:128;not null;default:'';unique_index" json:"log_url"` //Log路径

    CreatedAt helper.JSONTime `gorm:"not null;default:current_timestamp" json:"created_at"`
    UpdatedAt helper.JSONTime `gorm:"not null;default:current_timestamp" json:"updated_at"`
}

func (m DeviceLog) Model() *gorm.DB {
    return db.Mysql((&m).Connection()).Model(&m)
}

func (*DeviceLog) Connection() string {
    return "default"
}

func (*DeviceLog) TableName() string {
    return "device_log"
}
