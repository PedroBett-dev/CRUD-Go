package response

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"nam"`
	Age   int8   `json:"age"`
	Email string `json:"email"`
}
