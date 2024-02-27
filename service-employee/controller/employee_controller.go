package controller

import (
	"fmt"
	"net/http"
	"service-employee/config"
	"service-employee/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var user_uri string = "http://localhost:3001/user"

type WebResponse struct {
	Code   int
	Status string
	Data   interface{}
}

func CreateEmployee(c *fiber.Ctx) error {
	var requestBody model.Employee
	dbPostgres := config.InitPostgres().Table("employees")

	ctx, cancel := config.NewPostgreContext()
	defer cancel()
	err := c.BodyParser(&requestBody)
	if err != nil {
		panic(err)
	}
	requestBody.Id = uuid.New().String()
	access_token := c.Get("access_token")
	if len(access_token) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	req, err := http.NewRequest("GET", user_uri+"/auth", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		panic(err)
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", access_token)

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		panic(err)
	}
	defer resp.Body.Close()

	// Print the response
	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Headers:", resp.Header)

	if resp.Status != "200 OK" {
		c.Status(401).SendString("invalid token")
	}

	defer cancel()
	err = dbPostgres.WithContext(ctx).Create(requestBody).Error

	if err != nil {

		return c.Status(400).JSON(WebResponse{
			Code:   400,
			Status: "error while creating new user",
			Data:   requestBody.Name,
		})
	}
	return c.JSON(WebResponse{
		Code:   201,
		Status: "OK",
		Data:   requestBody,
	})
}
