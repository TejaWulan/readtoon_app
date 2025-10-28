package interactors

import (
	"errors"
	"time"

	"readtoon_app/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountInteractor struct {
	DB *gorm.DB
}

// CreateAccountProfile creates a new user profile
func (i *AccountInteractor) CreateAccountProfile(name, email, password, createdBy string) (*models.AccountProfile, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account := &models.AccountProfile{
		UID:         uuid.NewString(),
		Name:        name,
		Email:       email,
		Password:    string(hashed),
		CreatedBy:   createdBy,
		CreatedDate: time.Now(),
	}

	if err := i.DB.Create(account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

// ChangeName updates user's name
func (i *AccountInteractor) ChangeName(uid, newName, updatedBy string) error {
	return i.DB.Model(&models.AccountProfile{}).
		Where("uid = ?", uid).
		Updates(map[string]interface{}{
			"name":         newName,
			"updated_by":   updatedBy,
			"updated_date": time.Now(),
		}).Error
}

// ChangeEmail updates user's email
func (i *AccountInteractor) ChangeEmail(uid, newEmail, updatedBy string) error {
	return i.DB.Model(&models.AccountProfile{}).
		Where("uid = ?", uid).
		Updates(map[string]interface{}{
			"email":        newEmail,
			"updated_by":   updatedBy,
			"updated_date": time.Now(),
		}).Error
}

// ChangePassword updates user's password (hashed)
func (i *AccountInteractor) ChangePassword(uid, oldPassword, newPassword, updatedBy string) error {
	var account models.AccountProfile
	if err := i.DB.First(&account, "uid = ?", uid).Error; err != nil {
		return err
	}

	// Verify old password
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(oldPassword)); err != nil {
		return errors.New("invalid old password")
	}

	// Hash new password
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return i.DB.Model(&models.AccountProfile{}).
		Where("uid = ?", uid).
		Updates(map[string]interface{}{
			"password":     string(hashed),
			"updated_by":   updatedBy,
			"updated_date": time.Now(),
		}).Error
}

// UpdateAvatarPhoto changes the user's avatar photo URL
func (i *AccountInteractor) UpdateAvatarPhoto(uid, photoPath, updatedBy string) error {
	return i.DB.Model(&models.AccountProfile{}).
		Where("uid = ?", uid).
		Updates(map[string]interface{}{
			"avatar_photo": photoPath,
			"updated_by":   updatedBy,
			"updated_date": time.Now(),
		}).Error
}
