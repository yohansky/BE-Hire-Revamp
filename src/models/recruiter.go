package models

import "gorm.io/gorm"

type Recruiter struct {
	Id        uint   `json:"id"`
	Bidang    string `json:"bidang"`
	Kota      string `json:"kota"`
	Deskripsi string `json:"deskripsi"`
	Instagram string `json:"instagram"`
	Linkedin  string `json:"linkedin"`
	UserId    uint   `json:"user_id"`
	User      User   `json:"user" gorm:"foreignKey:UserId"`
}

func (recruiter *Recruiter) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Recruiter{}).Count(&total)

	return total
}

func (recruiter *Recruiter) Take(db *gorm.DB, limit int, offset int) interface{} {
	var recruiters []Recruiter

	db.Preload("User").Offset(offset).Limit(limit).Find(&recruiters)

	return recruiters
}
