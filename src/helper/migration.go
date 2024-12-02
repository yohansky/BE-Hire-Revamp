package helper

import (
	"be-hire-revamp/src/config"
	"be-hire-revamp/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Recruiter{},
		&models.Skill{},
		&models.Worker{},
		&models.Project{},
		&models.Experience{},
	)
}
