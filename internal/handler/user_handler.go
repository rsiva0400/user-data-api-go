package handler

import (
	"strconv"
	"time"
	"userdata-api/internal/models"
	"userdata-api/internal/service"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (userHandler *UserHandler) CreateUser(ctx *fiber.Ctx) error {

	var req models.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	// 🧩 Validate fields
	if err := validate.Struct(&req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errs := make(map[string]string)
		for _, vErr := range validationErrors {
			errs[vErr.Field()] = vErr.Error()
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_error": errs,
		})
	}
	// ✅ Create user via service
	dob, _ := time.Parse("2006-01-02", req.Dob)
	id, err := userHandler.userService.CreateUser(ctx.UserContext(), req.Name, dob)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"id":      id,
	})
}

func (userHandler *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	userID, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// get user by id
	err = userHandler.userService.DeleteUser(ctx.UserContext(), int32(userID))

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// 3️⃣ Return as JSON response
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (userHandler *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	userID, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var req models.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	// 🧩 Validate fields
	if err := validate.Struct(&req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errs := make(map[string]string)
		for _, vErr := range validationErrors {
			errs[vErr.Field()] = vErr.Error()
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_error": errs,
		})
	}
	// ✅ Create user via service
	dob, _ := time.Parse("2006-01-02", req.Dob)
	user, err := userHandler.userService.UpdateUser(ctx.UserContext(), int32(userID), req.Name, dob)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// 3️⃣ Return as JSON response
	return ctx.JSON(user)

}

func (userHandler *UserHandler) ListUsers(ctx *fiber.Ctx) error {
	// Default values
	limit := 10
	page := 1

	// Read from query params if provided
	if limitQuery := ctx.Query("limit"); limitQuery != "" {
		if val, err := strconv.Atoi(limitQuery); err == nil {
			limit = val
		}
	}

	if pageNumber := ctx.Query("page"); pageNumber != "" {
		if val, err := strconv.Atoi(pageNumber); err == nil {
			page = val
		}
	}
	limit = max(limit, 100)
	offset := (page - 1) * limit

	users, err := userHandler.userService.ListUsers(ctx.UserContext(), int32(limit), int32(offset))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"page":  page,
		"users": users,
	})
}

func (userHandler *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	userID, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// get user by id
	user, err := userHandler.userService.GetUserByID(ctx.UserContext(), int32(userID))

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// 3️⃣ Return as JSON response
	return ctx.JSON(user)
}
