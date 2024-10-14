package models

type Caption struct {
	Id     uint   `gorm:"primaryKey;autoIncrement"`
	Email  string `gorm:"size:50;not null"`
	Baris1 string `gorm:"type:text"`
	Baris2 string `gorm:"type:text"`
	Baris3 string `gorm:"type:text"`
	Baris4 string `gorm:"type:text"`
	Status bool   `gorm:"default:false"`
}
