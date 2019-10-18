package model

import (
	"time"
)

//Users Users
type Users struct {
	Id            int64     `xorm:"pk BIGINT(20)"`
	PrefList      string    `xorm:"not null TEXT"`
	LatestLogTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Name          string    `xorm:"not null VARCHAR(255)"`
}

//TableName users
func (Users) TableName() string {
	return "users"
}
