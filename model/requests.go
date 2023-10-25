package model

// user

type ReqRegister struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type ReqRegisterTeacher struct {
	Token    string `json:"token" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type ReqSearchUser struct {
	Token string `json:"token" validate:"required"`
	Query string `json:"query" validate:"required,min=3"`
}

// session

type ReqLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqChallenge struct {
	Token string `json:"token" validate:"required"`
}
