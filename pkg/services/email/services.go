package email

import "smtpserver/pkg/models"

type Service interface {
	CreateEmail(email models.Email) (*models.Email, error)
	DeleteEmails() error
}

type userSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &userSvc{repo: r}
}

// CreateEmail implements Service.
func (s *userSvc) CreateEmail(email models.Email) (*models.Email, error) {
	return s.repo.CreateEmail(email)
}

func (s *userSvc) DeleteEmails() error {
	return s.repo.DeleteEmails()
}
