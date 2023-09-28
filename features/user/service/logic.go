package service

import (
	"errors"
	"rahmat/to-do-list-app/features/user"
)

func NewServiceUser(repo user.DataUserInterface) user.ServiceUserInterface {
	return &userService{
		userData: repo,
	}
}

type userService struct {
	userData user.DataUserInterface
}

// CreateUser implements user.ServiceUserInterface.
func (s *userService) CreateUser(insert user.CoreUser) error {
	// falidasi unruk data yang wajib di isi

	if insert.Name == "" || insert.Email == "" || insert.Password == "" {
		return errors.New("validation error. name/email/password required")
	}

	err := s.userData.Insert(insert)
	if err != nil {
		return errors.New("Error inset data")
	}
	return err
}

// GetAll implements user.ServiceUserInterface.
func (s *userService) GetAll() ([]user.CoreUser, error) {
	result, err := s.userData.SelectAll()
	return result, err
}

// Update implements user.ServiceUserInterface.
func (s *userService) Update(insert user.CoreUser, id uint) error {
	if id == 0 {
		return errors.New("validation error. invalid ID")
	}

	err := s.userData.Update(insert, id)
	return err
}
