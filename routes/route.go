package routes

import (
	"readtoon_app/config"
	controllers "readtoon_app/controller"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	load := config.LoadConfig()

	// Account
	e.POST(load.API_VERSION+"/account/register", controllers.CreateAccountProfileController)
	e.POST(load.API_VERSION+"/account/updatename", controllers.ChangeNameController)
	e.POST(load.API_VERSION+"/account/updateemail", controllers.ChangeEmailController)
	e.POST(load.API_VERSION+"/account/updateemail", controllers.ChangePasswordController)
	e.POST(load.API_VERSION+"/account/updateemail", controllers.UpdateAvatarPhotoController)

}
