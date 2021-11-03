package models

type LoginRequest struct {
	UserName   string `json:"UserName" from:"UserName" binding:"required"`
	Password   string `json:"Password" from:"Password" binding:"required"`
	RememberMe bool   `json:"RememberMe" from:"RememberMe"`
}
