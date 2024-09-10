package main

import ()

type LoginPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailPayload struct {
	Email string `json:"email"`
}

type UsernamePayload struct {
	Username string `json:"username"`
}

type SearchPayload struct {
	Search string `json:"search"`
}

type RattingPayload struct {
	Product_id int     `json:"product_id"`
	User_id    int     `json:"user_id"`
	Start      float64 `json:"start"`
	Comment    string  `json:"comment"`
}

type CorrectEmail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type ChangePassword struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

type GetData struct {
	NameList string `json:"nameList"`
	Attribute string `json:"attribute"`
}

type DeleteData struct {
	NameList string `json:"nameList"`
	Attribute string `json:"attribute"`
	PrimaryKey interface{} `json:"primaryKey"`
}

type InsertData struct {
	NameList string `json:"nameList"`
	Data map[string]interface{} `json:"data"`
}

type UpdateData struct {
	NameList string `json:"nameList"`
	Data map[string]interface{} `json:"data"`
	Condition string `json:"condition"`
}