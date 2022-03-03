package seeders

import (
	"github.com/zumardirahman/golang-toko/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.UserFaker(db)},
		{Seeder: fakers.ProductFaker(db)},
	}
}

//perlu func utk proses seed ke database
func DBSeed(db *gorm.DB) error {
	//func dbseed ini akan melooping faker yang telah dibuat untuk diinsertkan kedalam databse
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}
