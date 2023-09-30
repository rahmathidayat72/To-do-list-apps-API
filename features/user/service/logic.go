package service

import (
	"errors"
	"rahmat/to-do-list-app/app/middlewares"
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

// Login implements user.ServiceUserInterface.
func (s *userService) Login(email string, password string) (dataLogin user.CoreUser, token string, err error) {
	// panic("unimplemented")

	dataLogin, err = s.userData.Login(email, password)
	if err != nil {
		return user.CoreUser{}, "", err
	}
	//pembutan token
	token, err = middlewares.CreatedToken(int(dataLogin.ID))
	if err != nil {
		return user.CoreUser{}, "", err
	}

	return dataLogin, token, nil
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

// SelectById implements user.ServiceUserInterface.
func (s *userService) SelectById(id uint) (user.CoreUser, error) {
	// mengechek id yang di input sudah benar atau tidak
	if id == 0 {
		return user.CoreUser{}, errors.New("validation error. invalid id")
	}
	result, err := s.userData.SelectById(id)
	return result, err
}

// Delete implements user.ServiceUserInterface.
func (s *userService) Delete(id uint) error {
	// panic("unimplemented")
	if id == 0 {
		return errors.New("validation error. invalid id")
	}
	err := s.userData.Delete(id)
	return err
}
