package entity

import "time"

type Cookie struct {
	UserID       uint
	SessionToken string
	MaxAge       time.Duration
}

type User struct {
	ID       uint   `json:"Id"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Status   string `json:"Status"`
}

type DBDeleteCookie struct {
	SessionToken string
}

func ToDBDeleteCookie(cookie *Cookie) *DBDeleteCookie {
	return &DBDeleteCookie{
		SessionToken: cookie.SessionToken,
	}
}
