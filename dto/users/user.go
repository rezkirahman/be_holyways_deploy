package usersdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Image    string `json:"image" form:"image"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Phone    string `json:"phone" form:"phone"`
	Image    string `json:"image" form:"image"`
}
