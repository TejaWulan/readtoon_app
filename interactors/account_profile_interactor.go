package interactors

import (
	"errors"
	"time"

	"readtoon_app/models"
	"readtoon_app/utility"

	"gorm.io/gorm"
)

type AccountInteractor struct {
	DB *gorm.DB
}

func CreateAccountProfile(db *gorm.DB, body map[string]any) (int, string, any) {
	bodyAccess := utility.MappingUnmarshal(body["access"])
	bodyPayload := utility.MappingUnmarshal(body["payload"])

	if bodyAccess == nil || bodyPayload == nil {
		return 7400, "cannot_empty:access_or_payload", nil
	}

	name, _ := bodyPayload["name"].(string)
	email, _ := bodyPayload["email"].(string)
	password, _ := bodyPayload["password"].(string)
	avatar, _ := bodyPayload["avatar_photo"].(string)

	if name == "" || email == "" || password == "" {
		return 7401, "missing_required_fields:name_email_password", nil
	}

	account := models.AccountProfile{
		UID:         utility.GenerateUUID(),
		Name:        name,
		Email:       email,
		Password:    utility.HashPassword(password),
		AvatarPhoto: avatar,
		CreatedDate: time.Now(),
		CreatedBy:   bodyAccess["user_uid"].(string),
	}

	if err := db.Create(&account).Error; err != nil {
		return 7500, "failed_to_create_account", err.Error()
	}

	return 7200, "success_create_account", account
}

func SignInAccount(db *gorm.DB, body map[string]any) (int, string, any) {
	bodyAccess := utility.MappingUnmarshal(body["access"])
	bodyPayload := utility.MappingUnmarshal(body["payload"])

	if bodyAccess == nil || bodyPayload == nil {
		return 7400, "cannot_empty:access_or_payload", nil
	}

	email, _ := bodyPayload["email"].(string)
	password, _ := bodyPayload["password"].(string)

	if email == "" || password == "" {
		return 7401, "missing_required_fields:email_password", nil
	}

	// Cari akun berdasarkan email
	var account models.AccountProfile
	if err := db.Where("email = ?", email).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 7404, "account_not_found", nil
		}
		return 7500, "failed_query_account", err.Error()
	}

	// Verifikasi password
	if !utility.CheckPasswordHash(password, account.Password) {
		return 7403, "invalid_credentials", nil
	}

	return 7200, "success_sign_in", map[string]any{
		"uid":          account.UID,
		"name":         account.Name,
		"email":        account.Email,
		"avatar_photo": account.AvatarPhoto,
		"last_login":   time.Now(),
	}
}

// ChangeName updates user's name

func ChangeName(db *gorm.DB, body map[string]any) (int, string, any) {
	bodyAccess := utility.MappingUnmarshal(body["access"])
	bodyPayload := utility.MappingUnmarshal(body["payload"])
	if bodyAccess == nil || bodyPayload == nil {
		return 7400, "cannot_empty:access_or_payload", nil
	}

	uid := bodyAccess["user_uid"].(string)
	newName, _ := bodyPayload["name"].(string)
	if newName == "" {
		return 7401, "missing_field:name", nil
	}

	if err := db.Model(&models.AccountProfile{}).
		Where("uid = ?", uid).
		Update("name", newName).Error; err != nil {
		return 7500, "failed_to_update_name", err.Error()
	}

	return 7200, "success_change_name", nil
}

func ChangeEmail(db *gorm.DB, body map[string]any) (int, string, any) {
	bodyAccess := utility.MappingUnmarshal(body["access"])
	bodyPayload := utility.MappingUnmarshal(body["payload"])
	if bodyAccess == nil || bodyPayload == nil {
		return 7400, "cannot_empty:access_or_payload", nil
	}

	uid := bodyAccess["user_uid"].(string)
	newEmail, _ := bodyPayload["email"].(string)
	if newEmail == "" {
		return 7401, "missing_field:email", nil
	}

	if err := db.Model(&models.AccountProfile{}).
		Where("uid = ?", uid).
		Update("email", newEmail).Error; err != nil {
		return 7500, "failed_to_update_email", err.Error()
	}

	return 7200, "success_change_email", nil
}

func ChangePassword(db *gorm.DB, body map[string]any) (int, string, any) {
	bodyAccess := utility.MappingUnmarshal(body["access"])
	bodyPayload := utility.MappingUnmarshal(body["payload"])
	if bodyAccess == nil || bodyPayload == nil {
		return 7400, "cannot_empty:access_or_payload", nil
	}

	uid := bodyAccess["user_uid"].(string)
	oldPassword, _ := bodyPayload["old_password"].(string)
	newPassword, _ := bodyPayload["new_password"].(string)
	if oldPassword == "" || newPassword == "" {
		return 7401, "missing_field:old_password_or_new_password", nil
	}

	var account models.AccountProfile
	if err := db.First(&account, "uid = ?", uid).Error; err != nil {
		return 7404, "account_not_found", err.Error()
	}

	if !utility.CheckPasswordHash(oldPassword, account.Password) {
		return 7402, "invalid_old_password", nil
	}

	hashed := utility.HashPassword(newPassword)
	if err := db.Model(&account).Update("password", hashed).Error; err != nil {
		return 7500, "failed_to_update_password", err.Error()
	}

	return 7200, "success_change_password", nil
}

// ✅ UPDATE AVATAR PHOTO
func UpdateAvatarPhoto(db *gorm.DB, body map[string]any) (int, string, any) {
	bodyAccess := utility.MappingUnmarshal(body["access"])
	bodyPayload := utility.MappingUnmarshal(body["payload"])
	if bodyAccess == nil || bodyPayload == nil {
		return 7400, "cannot_empty:access_or_payload", nil
	}

	uid := bodyAccess["user_uid"].(string)
	newPhoto, _ := bodyPayload["avatar_photo"].(string)
	if newPhoto == "" {
		return 7401, "missing_field:avatar_photo", nil
	}

	if err := db.Model(&models.AccountProfile{}).
		Where("uid = ?", uid).
		Update("avatar_photo", newPhoto).Error; err != nil {
		return 7500, "failed_to_update_avatar_photo", err.Error()
	}

	return 7200, "success_update_avatar_photo", nil
}
