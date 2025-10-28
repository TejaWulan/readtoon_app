package controllers

import (
	"net/http"

	"readtoon_app/interactors"

	"github.com/labstack/echo/v4"
)

type AccountController struct {
	Interactor *interactors.AccountInteractor
}

// POST /accounts
func (ctrl *AccountController) CreateAccountProfile(c echo.Context) error {
	type Request struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CreatedBy string `json:"created_by"`
	}

	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	account, err := ctrl.Interactor.CreateAccountProfile(req.Name, req.Email, req.Password, req.CreatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, account)
}

// PUT /accounts/:uid/name
func (ctrl *AccountController) ChangeName(c echo.Context) error {
	type Request struct {
		NewName   string `json:"new_name"`
		UpdatedBy string `json:"updated_by"`
	}
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	uid := c.Param("uid")
	err := ctrl.Interactor.ChangeName(uid, req.NewName, req.UpdatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "name updated"})
}

// PUT /accounts/:uid/email
func (ctrl *AccountController) ChangeEmail(c echo.Context) error {
	type Request struct {
		NewEmail  string `json:"new_email"`
		UpdatedBy string `json:"updated_by"`
	}
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	uid := c.Param("uid")
	err := ctrl.Interactor.ChangeEmail(uid, req.NewEmail, req.UpdatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "email updated"})
}

// PUT /accounts/:uid/password
func (ctrl *AccountController) ChangePassword(c echo.Context) error {
	type Request struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
		UpdatedBy   string `json:"updated_by"`
	}
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	uid := c.Param("uid")
	err := ctrl.Interactor.ChangePassword(uid, req.OldPassword, req.NewPassword, req.UpdatedBy)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "password updated"})
}

// PUT /accounts/:uid/avatar
func (ctrl *AccountController) UpdateAvatarPhoto(c echo.Context) error {
	type Request struct {
		AvatarPhoto string `json:"avatar_photo"`
		UpdatedBy   string `json:"updated_by"`
	}
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	uid := c.Param("uid")
	err := ctrl.Interactor.UpdateAvatarPhoto(uid, req.AvatarPhoto, req.UpdatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "avatar updated"})
}
