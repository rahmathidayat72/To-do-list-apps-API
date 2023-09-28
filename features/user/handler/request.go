package handler

import "rahmat/to-do-list-app/features/user"

type UserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RequestToCore(input UserRequest) user.CoreUser {
	return user.CoreUser{
		Name:        input.Name, 
		Email:       input.Email,
		Password:    input.Password,
		Address: 	 input.Address,
		PhoneNumber: input.PhoneNumber,
	}

}


