package controllers

import (
	"be-hire-revamp/src/config"
	"be-hire-revamp/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllWorkers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Worker{}, page))
}

func CreateWorker(c *fiber.Ctx) error {
	var worker models.Worker

	if err := c.BodyParser(&worker); err != nil {
		return err
	}

	config.DB.Create(&worker)

	return c.JSON(worker)
}

func GetWorker(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var worker models.Worker

	worker.Id = uint(id)

	config.DB.Preload("User").Preload("Skill").Find(&worker)

	return c.JSON(worker)
}

func GetWorkerByUserID(c *fiber.Ctx) error {

	id := c.Params("id")

	var worker []models.Worker
	if err := config.DB.Where("user_id = ?", id).First(&worker).Error; err != nil {

		return c.JSON(fiber.Map{"error": "Worker not found"})
	}

	return c.JSON(worker)
}

func UpdateWorker(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var worker models.Worker

	worker.Id = uint(id)

	if err := c.BodyParser(&worker); err != nil {
		return err
	}

	config.DB.Model(&worker).Updates(worker)

	return c.JSON(worker)
}

func DeleteWorker(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var worker models.Worker

	worker.Id = uint(id)

	config.DB.Delete(&worker)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
