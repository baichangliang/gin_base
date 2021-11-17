package manager

type LoginValidators struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required,min=4,max=20"`
}
