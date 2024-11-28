package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Perusahaan   string `json:"perusahaan"`
	Jabatan      string `json:"jabatan"`
	Nomortelepon string `json:"nomor_telepon"`
	Password     []byte `json:"-"`
	RoleId       uint   `json:"role_id"`
	Role         Role   `gorm:"foreignKey:RoleId"`
}

// nomer

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
