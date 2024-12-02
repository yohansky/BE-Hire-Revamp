package controllers

import (
	"be-hire-revamp/src/config"
	"be-hire-revamp/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllSkills(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Skill{}, page))
}

func CreateSkill(c *fiber.Ctx) error {
	var skill models.Skill

	if err := c.BodyParser(&skill); err != nil {
		return err
	}

	config.DB.Create(&skill)

	return c.JSON(skill)
}

func GetSkill(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var skill models.Skill

	skill.Id = uint(id)

	config.DB.Preload("User").Find(&skill)

	return c.JSON(skill)
}

func UpdateSkill(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var skill models.Skill

	skill.Id = uint(id)

	if err := c.BodyParser(&skill); err != nil {
		return err
	}

	config.DB.Model(&skill).Updates(skill)

	return c.JSON(skill)
}

func DeleteSkill(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var skill models.Skill

	skill.Id = uint(id)

	config.DB.Delete(&skill)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
