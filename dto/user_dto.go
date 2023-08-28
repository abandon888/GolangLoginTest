package dto

import "awesomeProject/model"

type UserDto struct {
	Name      string `json:"name" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
}

// ToUserDto 将User转化为UserDto
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
