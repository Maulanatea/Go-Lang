package req

type BookCreateReq struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}
