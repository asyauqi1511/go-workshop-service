package entity

import "time"

type User struct {
	UserID     int64
	Username   string
	Role       int
	CreateTime time.Time
	UpdateTime time.Time
}

type UserDetail struct {
	DetailUserID int64
	UserID       int64
	Name         string
	Address      string
	SchoolName   string
	CreateTime   time.Time
	UpdateTime   time.Time
}

type UserRegisterParam struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	SchoolName string `json:"school_name"`
}

type UserDetailResult struct {
	UserID     int64  `json:"user_id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	SchoolName string `json:"school_name"`
}
