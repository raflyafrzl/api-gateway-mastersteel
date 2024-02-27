package controller

import (
	"errors"
	"service-user/helpers"
	"service-user/model"

	"service-user/config"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WebResponse struct {
	Code   int
	Status string
	Data   interface{}
}

func Register(c *fiber.Ctx) error {
	var requestBody model.User
	//db := config.GetMongoDatabase().Collection("user")
	dbPostgres := config.InitPostgres().Table("users")
	requestBody.Id = uuid.New().String()

	ctx, cancel := config.NewPostgreContext()
	defer cancel()
	err := c.BodyParser(&requestBody)
	if err != nil {
		panic(err)
	}
	requestBody.Password = helpers.HashPassword([]byte(requestBody.Password))
	err = dbPostgres.WithContext(ctx).Create(requestBody).Error

	if err != nil {

		return c.Status(400).JSON(WebResponse{
			Code:   400,
			Status: "error while creating new user",
			Data:   requestBody.Email,
		})
	}

	return c.Status(201).JSON(WebResponse{
		Code:   201,
		Status: "OK",
		Data:   requestBody.Email,
	})
}

func Login(c *fiber.Ctx) error {
	dbPostgres := config.InitPostgres().Table("users")
	var requestBody model.User
	var result model.User

	c.BodyParser(&requestBody)
	ctx, cancel := config.NewPostgreContext()
	defer cancel()
	err := dbPostgres.WithContext(ctx).Where("email=?", requestBody.Email).Find(&result).Error

	if err != nil {
		return c.Status(401).JSON(WebResponse{
			Code:   401,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}
	checkPassword := helpers.ComparePassword([]byte(result.Password), []byte(requestBody.Password))
	if !checkPassword {
		return c.Status(401).JSON(WebResponse{
			Code:   401,
			Status: "BAD_REQUEST",
			Data:   errors.New("invalid password").Error(),
		})
	}

	access_token := helpers.SignToken(requestBody.Email)

	return c.Status(200).JSON(struct {
		Code        int
		Status      string
		AccessToken string
		Data        interface{}
	}{
		Code:        200,
		Status:      "OK",
		AccessToken: access_token,
		Data:        result,
	})
}

func Auth(c *fiber.Ctx) error {
	return c.JSON("OK")
}
