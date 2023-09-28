package data

import (
	"errors"
	"rahmat/to-do-list-app/features/user"
	"rahmat/to-do-list-app/helper"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewDataUser(db *gorm.DB) user.DataUserInterface {
	return &UserQuery{
		db: db,
	}
}

// Insert implements user.DataUserInterface.
func (r *UserQuery) Insert(insert user.CoreUser) error {
	// mapping dari struct core to struct gorm model
	// userInput := User{

	// 	Name:        insert.Name,
	// 	Email:       insert.Email,
	// 	Password:    insert.Password,
	// 	Address:     insert.Address,
	// 	PhoneNumber: insert.PhoneNumber,
	// }
	userInput := CoreToModel(insert)
	userInput.Password = helper.HashPassword(userInput.Password)

	//simpan ke DB
	tx := r.db.Create(&userInput)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("failed to insert, row affected is 0")
	}

	return nil

}

// SelectAll implements user.DataUserInterface.
func (r *UserQuery) SelectAll() ([]user.CoreUser, error) {
	// panic("unimplemented")
	var dataUser []User
	tx := r.db.Find(&dataUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// fmt.Println(tx.Error)
	// lakukan mepping dari struck gorm model ke struct code (entity)
	var userCore []user.CoreUser

	for _, v := range dataUser {
		// var user = user.CoreUser{
		// 	ID:          v.ID,
		// 	Name:        v.Name,
		// 	Email:       v.Email,
		// 	Password:    v.Password,
		// 	Address:     v.Address,
		// 	PhoneNumber: v.PhoneNumber,
		// 	CreatedAt:   v.CreatedAt,
		// 	UpdateAt:    v.UpdatedAt,
		// }
		var user = ModelToCore(v)
		userCore = append(userCore, user)

	}
	// fmt.Println(userCore)
	return userCore, nil
}

// Update implements user.DataUserInterface.
func (r *UserQuery) Update(insert user.CoreUser, id uint) error {
	// panic("unimplemented")
	var updateUser User
	tx := r.db.First(&updateUser, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return errors.New("User id not found")
		}
		return tx.Error
	}

	//mengecek ada data yang terupdate atau tidak
	if tx.RowsAffected == 0 {
		return errors.New("user id not found")
	}

	updateUser.Name = insert.Name
	updateUser.Email = insert.Email
	updateUser.Address = insert.Address
	updateUser.PhoneNumber = insert.PhoneNumber

	// Menggunakan helper.HashPassword untuk mengamankan password baru
	if insert.Password != "" {
		updateUser.Password = helper.HashPassword(insert.Password)
	}
	userGorm := CoreToModel(insert)
	userGorm.Password = helper.HashPassword(userGorm.Password)

	tx = r.db.Model(&User{}).Where("id=?", id).Updates(updateUser)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectById implements user.DataUserInterface.
func (r *UserQuery) SelectById(id uint) (user.CoreUser, error) {
	// panic("unimplemented")
	var userData User
	tx := r.db.Find(&userData, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {

			return user.CoreUser{}, errors.New("User id not found")
		}
		return user.CoreUser{}, tx.Error
	}
	//mapping data dari model ke core
	userCore := ModelToCore(userData)
	return userCore, nil
}

// Delete implements user.DataUserInterface.
func (r *UserQuery) Delete(id uint) error {
	// panic("unimplemented")

	var deleteUser = User{}
	tx := r.db.Delete(&deleteUser, id)
	if tx.Error != nil {
		errors.New("Failed delete user")
	}
	if tx.RowsAffected == 0 {
		return errors.New("user id not found")
	}
	return nil
}
