package requestModel

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

type BasicResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
