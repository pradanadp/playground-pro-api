package service

import (
	"errors"
	"fmt"

	"github.com/playground-pro-project/playground-pro-api/features/user"
	"github.com/playground-pro-project/playground-pro-api/utils/helper"
)

type userService struct {
	userData user.UserData
}

func New(repo user.UserData) user.UserService {
	return &userService{
		userData: repo,
	}
}

// CreateUser implements user.UserService.
func (us *userService) CreateUser(user user.UserCore) (string, error) {
	if user.Fullname == "" {
		return "", errors.New("error, name is required")
	}
	if user.Email == "" {
		return "", errors.New("error, mail is required")
	} else if _, err := helper.ValidateMailAddress(user.Email); !err {
		return "", errors.New("error, invalid email format")
	}
	if user.Phone == "" {
		return "", errors.New("error, phone is required")
	}
	if user.Password == "" {
		return "", errors.New("error, password is required")
	}

	err := helper.ValidatePassword(user.Password)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	userID, err := us.userData.Create(user)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return userID, nil
}

// DeleteUserByID implements user.UserService.
func (us *userService) DeleteByID(userID string) error {
	err := us.userData.DeleteByID(userID)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}

// GetUserByID implements user.UserService.
func (us *userService) GetByID(userID string) (user.UserCore, error) {
	userEntity, err := us.userData.GetByID(userID)
	if err != nil {
		return user.UserCore{}, fmt.Errorf("error: %w", err)
	}

	return userCore, nil
}

// Login implements user.UserService.
func (us *userService) Login(email string, password string) (user.UserCore, string, error) {
	if email == "" {
		return user.UserCore{}, "", errors.New("email is required")
	}
	if password == "" {
		return user.UserCore{}, "", errors.New("password is required")
	}

	loggedInUser, accessToken, err := us.userData.Login(email, password)
	if err != nil {
		return user.UserCore{}, "", fmt.Errorf("%w", err)
	}

	return loggedInUser, accessToken, nil
}

// UpdateUserByID implements user.UserService.
func (us *userService) UpdateByID(userID string, updatedUser user.UserCore) error {
	if updatedUser.Password != "" {
		err := helper.ValidatePassword(updatedUser.Password)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	if updatedUser.Email != "" {
		_, err := helper.ValidateMailAddress(updatedUser.Email)
		if !err {
			return errors.New("error: invalid email format")
		}
	}

	err := us.userData.UpdateByID(userID, updatedUser)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
