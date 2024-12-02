package models

import "gorm.io/gorm"

type Project struct {
	Id       uint   `json:"id"`
	Nama     string `json:"nama"`
	Link     string `json:"link"`
	Tipe     string `json:"tipe"`
	Gambar   string `json:"gambar"`
	WorkerId uint   `json:"worker_id"`
	Worker   Worker `json:"worker" gorm:"foreignKey:WorkerId"`
}

func (project *Project) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Project{}).Count(&total)

	return total
}

func (project *Project) Take(db *gorm.DB, limit int, offset int) interface{} {
	var projects []Project

	db.Preload("Worker").Offset(offset).Limit(limit).Find(&projects)

	return projects
}
