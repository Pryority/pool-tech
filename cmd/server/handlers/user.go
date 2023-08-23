// cmd/server/handlers/user.go
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pryority/pool-tech/cmd/server/models"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]models.User)

func RegisterUser(c *fiber.Ctx) error {
	// Parse the request body and validate the data
	var userData models.CreateUserSchema
	if err := c.BodyParser(&userData); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Hash the user's password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create a new user")
	}
	userData.Password = string(hashedPassword)

	// Create the new user in the database
	if _, err := models.CreateUser(db, userData); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create a new user")
	}

	// Return a response to indicate successful registration
	// return c.JSON(fiber.Map{
	// 	"message": "User registered successfully!",
	// })
	return c.Render("partials/register-success.html", fiber.StatusCreated)
}

// Login handles the login request and redirects to the home page.
func Login(c *fiber.Ctx) error {
	var user models.User

	// Parse the request body and bind it to the user model
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Retrieve the user from the in-memory store based on the provided username
	storedUser, found := users[user.Email]
	if !found {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Compare the provided password with the hashed password from the store
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Successful login
	// Handle user login logic here (e.g., set session, JWT, etc.)

	// Redirect to the home page after successful login
	return c.Redirect("/", fiber.StatusFound)
}
