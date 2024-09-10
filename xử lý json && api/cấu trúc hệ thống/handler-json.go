package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Login(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var object LoginPayload
				err = json.Unmarshal(body, &object)
				check_err(err)
				check, status := login_account(object.Username, object.Email, object.Password)
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}

func Register(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var object RegisterPayload
				err = json.Unmarshal(body, &object)
				check_err(err)
				check, status := register_account(object.Username, object.Email, object.Password)
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}

func Search_products(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var search SearchPayload
				err = json.Unmarshal(body, &search)
				check_err(err)

				products, status := select_product(search.Search)
				response := map[string]interface{}{
					"product": products,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}

func Ratting_product(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var ratting RattingPayload
				err = json.Unmarshal(body, &ratting)
				check_err(err)

				check, status := ratting_products(ratting)
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		}
	}
}

func Send_code(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var send_code EmailPayload
				err = json.Unmarshal(body, &send_code)
				check_err(err)

				check, status := send_email("hiencute3321@gmail.com", "ppmdpxecepmjvdhu", send_code.Email)
				reponseData := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&reponseData)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}

func Verify_code(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var cor_em CorrectEmail
				err = json.Unmarshal(body, &cor_em)
				check_err(err)

				check, status := verify_email(cor_em.Email, cor_em.Code)
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}

func Change_password(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var change_pass ChangePassword
				err = json.Unmarshal(body, &change_pass)
				check_err(err)

				check, status := change_password(change_pass.Email, change_pass.NewPassword)
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}

func Get_data(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var data GetData
				err = json.Unmarshal(body, &data)
				check_err(err)

				list, status := get_data(data.NameList, data.Attribute)

				response := map[string]interface{}{
					"data":   list,
					"status": status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}

func Delete_data(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var data DeleteData
				err = json.Unmarshal(body, &data)
				check_err(err)

				check, status := delete_data(data.NameList, data.Attribute, data.PrimaryKey)

				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}

func Insert_data(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var data InsertData
				err = json.Unmarshal(body, &data)
				check_err(err)

				check, status := insert_data(data.NameList, data.Data)
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}

func Upload_image(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Bắt đầu xử lý upload ảnh")
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				log.Println("Đang xử lý yêu cầu POST")
				r.ParseMultipartForm(10 << 20)
				file, handler, err := r.FormFile("image")
				if err != nil {
					log.Printf("Lỗi khi đọc file: %v", err)
					http.Error(w, "Lỗi khi đọc file", http.StatusBadRequest)
					return
				}
				log.Printf("Đã đọc file: %s", handler.Filename)

				defer file.Close()

				if !is_allowed_file_type(handler.Filename) {
					log.Printf("Loại file không được chấp nhận: %s", handler.Filename)
					http.Error(w, "Chỉ chấp nhận file PNG hoặc JPG", http.StatusBadRequest)
					return
				}
				log.Println("Loại file hợp lệ")

				uploadDir := "./image-products/"
				if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
					log.Printf("Thư mục upload không tồn tại, đang tạo: %s", uploadDir)
					os.Mkdir(uploadDir, 0755)
				}
				filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename))
				filepath := filepath.Join(uploadDir, filename)
				log.Printf("Đường dẫn file: %s", filepath)

				dst, err := os.Create(filepath)
				if err != nil {
					log.Printf("Lỗi khi tạo file: %v", err)
					http.Error(w, "Lỗi khi tạo file", http.StatusInternalServerError)
					return
				}
				log.Printf("Đã tạo file tại: %s", filepath)
				defer dst.Close()

				if _, err = io.Copy(dst, file); err != nil {
					log.Printf("Lỗi khi lưu file: %v", err)
					http.Error(w, "Lỗi khi lưu file", http.StatusInternalServerError)
					return
				}
				log.Println("Đã lưu file thành công")

				productID := r.FormValue("product_id")
				log.Printf("Product ID: %s", productID)

				success, message := save_image_path_to_database(productID, filepath)
				if !success {
					log.Printf("Lỗi khi lưu đường dẫn ảnh vào database: %s", message)
					http.Error(w, message, http.StatusInternalServerError)
					return
				}
				log.Println("Đã lưu đường dẫn ảnh vào database")

				response := map[string]string{
					"filepath": filepath,
					"message":  message,
				}
				log.Printf("Phản hồi: %v", response)
				json.NewEncoder(w).Encode(response)
			}
		default:
			log.Printf("Phương thức không được hỗ trợ: %s", r.Method)
			fmt.Println("Method is not used")
		}
		log.Println("Kết thúc xử lý upload ảnh")
	}
}

func Update_data(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var data UpdateData
				err = json.Unmarshal(body, &data)
				check_err(err)

				check, status := update_data(data.NameList, data.Data, data.Condition)
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	}
}
