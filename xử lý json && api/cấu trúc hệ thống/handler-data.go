package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func sendPostRequest(url string, payload interface{}) (map[string]interface{}, error) {
	jsonPayload, err := json.Marshal(payload)
	check_err(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	check_err(err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	check_err(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	check_err(err)

	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	check_err(err)

	return responseData, nil
}

func login_account(username string, email string, password string) (bool, string) {
	payload := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}

	responseData, err := sendPostRequest("http://127.0.0.1:5050/login", payload)
	check_err(err)

	fmt.Printf("Response Data: %v\n", responseData)

	success, ok := responseData["success"].(bool)
	if !ok {
		log.Fatal("Chuyển đổi dữ liệu thất bại")
	}

	return success, "Đã gửi dữ liệu"
}

func register_account(username string, email string, password string) (bool, string) {
	payload := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}

	responseData, err := sendPostRequest("http://127.0.0.1:5050/register", payload)
	check_err(err)

	fmt.Printf("Response Data: %v\n", responseData)

	success, ok := responseData["success"].(bool)
	if !ok {
		log.Fatal("Chuyển đổi dữ liệu thất bại")
	}

	return success, "Đã gửi dữ liệu"
}

func select_product(input string) ([]string, string) {
	payload := map[string]string{
		"query": "SELECT category_name, product_name FROM Products",
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/select", payload)
	check_err(err)

	input = Translate_VN_to_EN(input)
	responseData = to_lowercase_map(responseData)

	if responseData["success"] != true {
		return nil, "Truy vấn thất bại"
	}

	Input := split_string(input)

	products := search_products(Input, responseData)
	return products, "Truy vấn thành công"
}

func ratting_products(evaluate RattingPayload) (bool, string) {
	query := fmt.Sprintf(
		"INSERT INTO Ratting (product_id, user_id, start, comment) VALUES (%d, %d, %f, '%s')",
		evaluate.Product_id, evaluate.User_id, evaluate.Start, evaluate.Comment,
	)
	payload := map[string]interface{}{
		"query": query,
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/insert", payload)
	check_err(err)

	if responseData["success"] != true {
		return false, "Thêm dữ liệu thất bại kiểm tra lại thuộc tính"
	}
	return true, "Thêm thanh công"
}

func send_email(sender string, password string, receiver string) (bool, string) {
	payload := map[string]string{
		"sender":   sender,
		"password": password,
		"receiver": receiver,
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/send_code", payload)
	check_err(err)

	if responseData["success"] != true {
		return false, "Gửi email thất bại"
	}
	return true, "Gửi email thành công"
}

func verify_email(email string, code string) (bool, string) {
	payload := map[string]string{
		"email": email,
		"code":  code,
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/verify_code", payload)
	check_err(err)

	if responseData["success"] != true {
		return false, "Xác thực email thất bại"
	}
	return true, "Xác thực email thành công"
}

func change_password(email string, newPassword string) (bool, string) {
	payload_passwd := map[string]string{
		"email":    email,
		"password": newPassword,
	}

	Password, err := sendPostRequest("http://127.0.0.1:5050/encode_data", payload_passwd)
	check_err(err)

	query := fmt.Sprintf("UPDATE Users SET password = '%s' WHERE email = '%s'", Password["data"], email)

	payload := map[string]string{
		"query": query,
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/update", payload)
	check_err(err)

	if responseData["success"] != true {
		return false, "Thay đổi mật khẩu thất bại"
	}
	return true, "Thay đổi mật khẩu thành công"
}

func get_data(name_list string, attribute string) (map[string]interface{}, string) {
	query := fmt.Sprintf("SELECT %s FROM %s", attribute, name_list)

	payload := map[string]string{
		"query": query,
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/select", payload)
	check_err(err)

	if responseData["success"] != true {
		return nil, "Truy vấn thất bại"
	}
	for key, value := range responseData {
		responseData[key] = convert_type(value)
	}

	return responseData, "Truy vấn thành công"
}

func delete_data(name_list string, attribute string, primary_key interface{}) (bool, string) {
	primary_key = convert_string(primary_key)
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = '%s'", name_list, attribute, primary_key)

	payload := map[string]string{
		"query": query,
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/delete", payload)
	check_err(err)
	if responseData["success"] != true {
		return false, "Xóa dữ liệu thất bại"
	}
	return true, "Xóa dữ liệu thành công"
}

func insert_data(name_list string, data map[string]interface{}) (bool, string) {
	atttribute := get_keys_string(data)
	values := get_values_string(data)
	fmt.Println(values)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", name_list, atttribute, values)
	payload := map[string]string{
		"query": query,
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/insert", payload)
	check_err(err)
	if responseData["success"] != true {
		return false, "Thêm dữ liệu thất bại"
	}
	return true, "Thêm dữ liệu thành công"
}

func save_image_path_to_database(product_id string, image_path string) (bool, string) {
	query := fmt.Sprintf("UPDATE Products SET image = '%s' WHERE product_id = '%s'", image_path, product_id)

	payload := map[string]string{
		"query": query,
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/update", payload)
	check_err(err)

	if responseData["success"] != true {
		return false, "Cập nhật đường dẫn ảnh thất bại"
	}
	return true, "Cập nhật đường dẫn ảnh thành công"
}

func update_data(name_list string, data map[string]interface{}, condition string) (bool, string) {
	
	Condition := data[condition]
	setString := Systax_update_database(data, []string{condition})
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = '%v'", name_list, setString, condition, Condition)

	payload := map[string]string{
		"query": query,
	}
	responseData, err := sendPostRequest("http://127.0.0.1:5050/update", payload)
	check_err(err)
	if responseData["success"] != true {
		return false, "Cập nhật dữ liệu thất bại"
	}
	return true, "Cập nhật dữ liệu thành công"
}
