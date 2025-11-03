package controllers

import (
	"net/http"

	"readtoon_app/database"
	"readtoon_app/interactors"

	"github.com/labstack/echo/v4"
)

type AccountController struct {
	Interactor *interactors.AccountInteractor
}

// POST /accounts

func CreateAccountProfileController(c echo.Context) error {
	var body map[string]any
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  7400,
			"message": "invalid_body_format",
		})
	}

	status, message, result := interactors.CreateAccountProfile(database.DB, body)
	return c.JSON(http.StatusOK, map[string]any{
		"status":  status,
		"message": message,
		"data":    result,
	})
}

func ChangeNameController(c echo.Context) error {
	var body map[string]any
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  7400,
			"message": "invalid_body_format",
		})
	}

	status, message, result := interactors.ChangeName(database.DB, body)
	return c.JSON(http.StatusOK, map[string]any{
		"status":  status,
		"message": message,
		"data":    result,
	})
}

func ChangeEmailController(c echo.Context) error {
	var body map[string]any
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  7400,
			"message": "invalid_body_format",
		})
	}

	status, message, result := interactors.ChangeEmail(database.DB, body)
	return c.JSON(http.StatusOK, map[string]any{
		"status":  status,
		"message": message,
		"data":    result,
	})
}

func ChangePasswordController(c echo.Context) error {
	var body map[string]any
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  7400,
			"message": "invalid_body_format",
		})
	}

	status, message, result := interactors.ChangePassword(database.DB, body)
	return c.JSON(http.StatusOK, map[string]any{
		"status":  status,
		"message": message,
		"data":    result,
	})
}

func UpdateAvatarPhotoController(c echo.Context) error {
	var body map[string]any
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  7400,
			"message": "invalid_body_format",
		})
	}

	status, message, result := interactors.UpdateAvatarPhoto(database.DB, body)
	return c.JSON(http.StatusOK, map[string]any{
		"status":  status,
		"message": message,
		"data":    result,
	})
}
