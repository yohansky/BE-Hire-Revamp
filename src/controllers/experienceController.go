package controllers

import (
	"be-hire-revamp/src/config"
	"be-hire-revamp/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllExperiences(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Experience{}, page))
}

func CreateExperience(c *fiber.Ctx) error {
	var experience models.Experience

	if err := c.BodyParser(&experience); err != nil {
		return err
	}

	config.DB.Create(&experience)

	return c.JSON(experience)
}

func GetExperience(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var experience models.Experience

	experience.Id = uint(id)

	config.DB.Preload("Worker").Find(&experience)

	return c.JSON(experience)
}

func GetWorkerByWorkerIDExperience(c *fiber.Ctx) error {
	// Ambil ID dari parameter URL
	id := c.Params("id")

	// Lakukan pencarian Worker berdasarkan ID di database
	var experience []models.Experience
	if err := config.DB.Where("worker_id = ?", id).First(&experience).Error; err != nil {
		// Jika tidak ditemukan, kirim respons error 404
		return c.JSON(fiber.Map{"error": "Worker not found"})
	}

	// Kirim respons JSON dengan data Worker yang ditemukan
	return c.JSON(experience)
}

func GetExperiencesByWorkerID(c *fiber.Ctx) error {

	id := c.Params("id")

	var experiences []models.Experience
	if err := config.DB.Where("worker_id = ?", id).Find(&experiences).Error; err != nil {

		return c.JSON(fiber.Map{"error": "experiences not found"})
	}

	return c.JSON(experiences)
}

func UpdateExperience(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var experience models.Experience

	experience.Id = uint(id)

	if err := c.BodyParser(&experience); err != nil {
		return err
	}

	config.DB.Model(&experience).Updates(experience)

	return c.JSON(experience)
}

func DeleteExperience(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var experience models.Experience

	experience.Id = uint(id)

	config.DB.Delete(&experience)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
