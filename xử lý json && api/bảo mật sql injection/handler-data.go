package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func check_account(username string, email string) (string, string, error) {
	db, err := Connect_owner()
	check_err(err)

	storePassword := ""
	var Email string
	if username != "" {
		err = db.DB.QueryRow("SELECT email FROM Users WHERE username = ?", username).Scan(&Email)
		if err != nil {
			return "", "", err
		}
	} else {
		Email = email
	}
	db.DB.QueryRow("SELECT password FROM Users WHERE email = ?", Email).Scan(&storePassword)
	return storePassword, Email, nil
}

func check_login(username string, email string, password string) (bool, string) {
	storePassword, log, err := check_account(username, email)
	if err != nil {
		return false, err.Error()
	}
	if storePassword == "" {
		return false, "Sai mật khẩu"
	}
	pass := encode_data(log, password, 2)
	return storePassword == pass, "Đúng mật khẩu"
}

func check_username(username string) bool {
	db, err := Connect_owner()
	check_err(err)

	count := 0
	err = db.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE username = ?", username).Scan(&count)
	return err == nil && count == 0
}

func check_email(email string) bool {
	db, err := Connect_owner()
	check_err(err)

	count := 0
	err = db.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE email = ?", email).Scan(&count)
	return err == nil && count == 0
}

func sign_up(username string, email string, Password string) (bool, string) {
	fmt.Println("Starting sign_up function")
	pass := encode_data(email, Password, 2)
	db, err := Connect_owner()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return false, "Lỗi kết nối database từ api đăng nhập"
	}

	if !check_email(email) {
		fmt.Println("Email already exists")
		return false, "Email đã tồn tại"
	}

	if !check_username(username) && username != "" {
		fmt.Println("Username already exists")
		return false, "username đã tồn tại"
	}

	if username == "" {
		_, err = db.DB.Exec("INSERT INTO Users (email, password) VALUES (?, ?)", email, pass)
		if err != nil {
			fmt.Println("Error inserting email and password:", err)
			return false, "Lỗi quá trình thêm email và password vào database"
		}
		return true, "Thêm thành công"
	}

	_, err = db.DB.Exec("INSERT INTO Users (username, email, password) VALUES (?, ?, ?)", username, email, pass)
	if err != nil {
		fmt.Println("Error inserting username, email, and password:", err)
		return false, "Lỗi quá trình thêm username và email và password vào database"
	}
	return true, "Thêm thành công"
}
