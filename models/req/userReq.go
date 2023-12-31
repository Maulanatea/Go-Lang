package req

type UserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email,min=3,max=32"`
	Password string `json:"password" validate:"required,min=3,max=32"`
}

type UserUpdate struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email,min=3,max=32"`
}
