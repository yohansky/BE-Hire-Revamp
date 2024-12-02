package controllers

import (
	"be-hire-revamp/src/config"
	"be-hire-revamp/src/models"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.User{}, page))
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	config.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var user models.User

	user.Id = uint(id)

	if err := config.DB.Preload("Role").Find(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "ID tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.JSON(fiber.Map{
			"Message": "Id tidak ditemukan",
		})
	}

	var user models.User

	user.Id = uint(id)

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	config.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var user models.User

	user.Id = uint(id)

	config.DB.Delete(&user)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
