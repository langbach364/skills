package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

type DBInfo struct {
	DB *sql.DB
}

func check_err(err error) {
	if err != nil {
		println(err)
		log.Fatal(err)
	}
}

func check_error_connect_database(dbInfo *DBInfo) (bool, string) {
	err := dbInfo.DB.Ping()
	return err != nil, "Lỗi kết nối database"
}
func enable_middleware_cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Cors := cors.New(cors.Options{
			AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Language", "Content-Type"},
			AllowedMethods:   []string{"POST"},
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			Debug:            true,
		})
		Cors.ServeHTTP(w, r, next.ServeHTTP)
	})
}

func Connect_owner() (*DBInfo, error) {
	connStr := "root:@ztegc4df9f4e@tcp(172.21.0.3:3306)/SHOP"
	db, err := sql.Open("mysql", connStr)
	check_err(err)
	return &DBInfo{DB: db}, nil
}

func Router_login(router *http.ServeMux) {
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var data data_user
				err = json.Unmarshal(body, &data)
				check_err(err)

				check, status := check_login(data.Username, data.Email, data.Password)
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			fmt.Println("Method is not used")
		}
	})
}

func Router_register(router *http.ServeMux) {
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			body, err := io.ReadAll(r.Body)
			check_err(err)
			var data data_user
			err = json.Unmarshal(body, &data)
			check_err(err)

			check, status := sign_up(data.Username, data.Email, data.Password)
			response := map[string]interface{}{
				"success": check,
				"status":  status,
			}
			json.NewEncoder(w).Encode(&response)
		default:
			fmt.Println("Method is not used")
		}
	})
}

func Router_Email(router *http.ServeMux) {
	router.HandleFunc("/send_code", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			body, err := io.ReadAll(r.Body)
			check_err(err)

			var data sender
			err = json.Unmarshal(body, &data)
			check_err(err)

			check, status := Send_Email(data.Email_sender, data.Password_sender, data.Email_recevier)
			response := map[string]interface{}{
				"success": check,
				"status":  status,
			}
			json.NewEncoder(w).Encode(&response)
		default:
			fmt.Println("Method is not used")
		}
	})
}

func Router_verify_email(router *http.ServeMux) {
	router.HandleFunc("/verify_code", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			body, err := io.ReadAll(r.Body)
			check_err(err)

			var data check_code
			err = json.Unmarshal(body, &data)
			check_err(err)

			check, status := verify_email(data.Email, data.Code)
			response := map[string]interface{}{
				"success": check,
				"status":  status,
			}
			json.NewEncoder(w).Encode(&response)
		default:
			fmt.Println("Method is not used")
		}
	})
}

func Router_encode_data(router *http.ServeMux) {
	router.HandleFunc("/encode_data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			body, err := io.ReadAll(r.Body)
			check_err(err)

			var data encode_passwd
			err = json.Unmarshal(body, &data)
			check_err(err)

			response := map[string]interface{}{
				"success": true,
				"status":  "success",
				"data":    encode_data(data.Email, data.Password, 2),
			}
			json.NewEncoder(w).Encode(&response)

		default:
			fmt.Println("Method is not used")
		}
	})
}

func Add_Router(router *http.ServeMux) {
	Router_login(router)
	Router_register(router)
	Router_Email(router)
	Router_verify_email(router)
	Router_encode_data(router)
}

func muxtiplexer_router(router *http.ServeMux) {
	Add_Router(router)

	dbInfo_owner, err := Connect_owner()
	check, status := check_error_connect_database(dbInfo_owner)

	if check {
		fmt.Println(status)
		return
	}
	check_err(err)
	router.HandleFunc("/select", select_Handler(dbInfo_owner))
	router.HandleFunc("/delete", delete_Handler(dbInfo_owner))
	router.HandleFunc("/insert", insert_Handler(dbInfo_owner))
	router.HandleFunc("/update", update_Handler(dbInfo_owner))
}

func Create_server() {
	router := http.NewServeMux()
	muxtiplexer_router(router)

	server := http.Server{
		Addr:    ":5050",
		Handler: enable_middleware_cors(router),
	}
	server.ListenAndServe()
}
