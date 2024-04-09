package admin

type AdminRequestBody struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  *string `json:"password"`
	CreatedBy *int    `json:"createdBy"`
	UpdatedBy *int    `json:"updatedBy"`
}
