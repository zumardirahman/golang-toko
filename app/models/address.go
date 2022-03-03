package models

import (
	"time"
)

type Address struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User      User   //belongs to
	UserID    string `gorm:"size:36;index"` //belongs to
	Name      string `gorm:"size:100"`
	IsPrimary bool   //menentukan alamat utama
	CityID    string //perlu menginputkan id kota dan provinsi
	Address1  string `gorm:"size:255"`
	Address2  string `gorm:"size:255"`
	Phone     string `gorm:"size:100"`
	Email     string `gorm:"size:100"`
	PostCode  string `gorm:"size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
