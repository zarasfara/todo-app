package entities

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
