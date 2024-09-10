package main

type query struct {
	Query string `json:"query"`
}

type sender struct {
	Email_sender    string `json:"sender"`
	Password_sender string `json:"password"`
	Email_recevier  string `json:"receiver"`
}

type check_code struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type data_user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type encode_passwd struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}