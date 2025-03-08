package email

import (
	"smtpserver/pkg/models"

	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateEmail(email models.Email) (*models.Email, error) {
	err := r.DB.Table("emails").Create(&email).Error
	if err != nil {
		return nil, err
	}
	return &email, nil
}

// UpdateUser implements Repository.
func (r *repo) DeleteEmails() error {
	return r.DB.Model(&models.Email{}).Delete(&models.Email{}).Where("created_at < NOW() - INTERVAL '12 hours'").Error
}
