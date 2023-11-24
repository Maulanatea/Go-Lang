package req

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `jason:"password" validate:"required"`
}
