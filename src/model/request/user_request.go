package request

type UserRequest struct {
	Name     string `json:"name" biding:"required, min=4 max=50"`
	Age      int8   `json:"age" biding:""`
	Email    string `json:"email" biding:"required"`
	Password string `json:"password" biding:"requiried"`
}
