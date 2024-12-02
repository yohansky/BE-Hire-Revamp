package controllers

import (
	"be-hire-revamp/src/config"
	"be-hire-revamp/src/helper"
	"be-hire-revamp/src/models"
	"be-hire-revamp/src/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllProjects(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Project{}, page))
}

func CreateProject(c *fiber.Ctx) error {
	file, err := c.FormFile("Gambar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Gagal mengunggah file: " + err.Error())
	}

	maxFileSize := int64(2 << 20)
	if err := helper.SizeUploadValidation(file.Size, maxFileSize); err != nil {
		return err
	}

	fileHeader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal membuka file: " + err.Error())
	}
	defer fileHeader.Close()

	buffer := make([]byte, 512)
	if _, err := fileHeader.Read(buffer); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal membaca file: " + err.Error())
	}

	validFileTypes := []string{"image/png", "image/jpeg", "image/jpg", "application/pdf"}
	if err := helper.TypeUploadValidation(buffer, validFileTypes); err != nil {
		return err
	}

	uploadResult, err := services.UploadCloudinary(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	values := form.Value

	workerID, err := strconv.ParseUint(values["WorkerId"][0], 10, 64)

	project := models.Project{
		Nama:     values["Nama"][0],
		Link:     values["Link"][0],
		Tipe:     values["Tipe"][0],
		Gambar:   uploadResult.SecureURL,
		WorkerId: uint(workerID),
	}

	config.DB.Create(&project)

	return c.JSON(fiber.Map{
		"Message": "Project created",
		"data":    project,
	})
}

func GetProject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var project models.Project

	project.Id = uint(id)

	config.DB.Preload("Worker").Find(&project)

	return c.JSON(project)
}

func GetWorkerByWorkerIDProject(c *fiber.Ctx) error {
	id := c.Params("id")

	var project []models.Project
	if err := config.DB.Where("worker_id = ?", id).First(&project).Error; err != nil {
		return c.JSON(fiber.Map{"error": "Projects not found"})
	}
	return c.JSON(project)
}

func GetProjectsByWorkerID(c *fiber.Ctx) error {

	id := c.Params("id")

	var projects []models.Project
	if err := config.DB.Where("worker_id = ?", id).Find(&projects).Error; err != nil {

		return c.JSON(fiber.Map{"error": "Projects not found"})
	}

	return c.JSON(projects)
}

func UpdateProject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var project models.Project

	project.Id = uint(id)

	if err := c.BodyParser(&project); err != nil {
		return err
	}

	config.DB.Model(&project).Updates(project)

	return c.JSON(project)
}

func DeleteProject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var project models.Project

	project.Id = uint(id)

	config.DB.Delete(&project)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
