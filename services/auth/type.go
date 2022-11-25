package auth

type SignupReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"string"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type BaseRes struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var SuccessMessage = "Success"
var ErrInternal = "Internal error"
