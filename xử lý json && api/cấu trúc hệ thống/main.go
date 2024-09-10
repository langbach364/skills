package main

import (
	"net/http"

	"github.com/rs/cors"
)

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

func muxtiplexer_router(router *http.ServeMux) {
	router.HandleFunc("/login", Login(router))
	router.HandleFunc("/register", Register(router))
	router.HandleFunc("/send_code", Send_code(router))
	router.HandleFunc("/verify_code", Verify_code(router))
	router.HandleFunc("/change_password", Change_password(router))
	router.HandleFunc("/get_data", Get_data(router))
	router.HandleFunc("/delete_data", Delete_data(router))
	router.HandleFunc("/insert_data", Insert_data(router))
	router.HandleFunc("/upload_image", Upload_image(router))
	router.HandleFunc("/update_data", Update_data(router))
}

func Create_server() {
	router := http.NewServeMux()
	muxtiplexer_router(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: enable_middleware_cors(router),
	}
	server.ListenAndServe()
}

func main() {
	Create_server()
}
	