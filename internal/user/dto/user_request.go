package user

type UserRequestBody struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedBy *int   `json:"createdBy"`
}
