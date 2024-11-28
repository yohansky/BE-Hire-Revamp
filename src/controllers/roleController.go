package controllers

import (
	"be-hire-revamp/src/config"
	"be-hire-revamp/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	config.DB.Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	role := models.Role{
		Name: roleDto["Name"].(string),
	}

	config.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	config.DB.Find(&role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	config.DB.Exec("DELETE FROM role_permission WHERE role_id = ?", id)

	role := models.Role{
		Id: uint(id),
	}

	config.DB.Delete(&role)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
