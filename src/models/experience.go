package models

import "gorm.io/gorm"

type Experience struct {
	Id         uint   `json:"id"`
	Posisi     string `json:"posisi"`
	Perusahaan string `json:"perusahaan"`
	Tahun      string `json:"tahun"`
	Deskripsi  string `json:"deskripsi"`
	WorkerId   uint   `json:"worker_id"`
	Worker     Worker `json:"worker" gorm:"foreignKey:WorkerId"`
}

func (experience *Experience) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Experience{}).Count(&total)

	return total
}

func (experience *Experience) Take(db *gorm.DB, limit int, offset int) interface{} {
	var experiences []Experience

	db.Preload("Worker").Offset(offset).Limit(limit).Find(&experiences)

	return experiences
}
