package controllers

import (
	"be-hire-revamp/src/config"
	"be-hire-revamp/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllRecruiters(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Recruiter{}, page))
}

func CreateRecruiter(c *fiber.Ctx) error {
	var recruiter models.Recruiter

	if err := c.BodyParser(&recruiter); err != nil {
		return err
	}

	config.DB.Create(&recruiter)

	return c.JSON(recruiter)
}

func GetRecruiter(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var recruiter models.Recruiter

	recruiter.Id = uint(id)

	config.DB.Preload("User").Find(&recruiter)

	return c.JSON(recruiter)
}

func UpdateRecruiter(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var recruiter models.Recruiter

	recruiter.Id = uint(id)

	if err := c.BodyParser(&recruiter); err != nil {
		return err
	}

	config.DB.Model(&recruiter).Updates(recruiter)

	return c.JSON(recruiter)
}

func DeleteRecruiter(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var recruiter models.Recruiter

	recruiter.Id = uint(id)

	config.DB.Delete(&recruiter)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
