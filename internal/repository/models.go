package repository

import "gorm.io/gorm"

type Author struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(200);not null"`
}

func (Author) TableName() string {
	return "authors"
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Author{},
	)
}
