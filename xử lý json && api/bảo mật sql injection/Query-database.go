package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func insert_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				check, status := Structure_query(query.Query, "insert")

				fmt.Println("Cấu trúc câu truy vấn: ", query.Query)
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				if check {
					dbInfo.DB.Exec(query.Query)
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}

	}
}

func select_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				check, status := Structure_query(query.Query, "select")

				if !check {
					response := map[string]interface{}{
						"success": check,
						"status":  status,
					}
					json.NewEncoder(w).Encode(response)
					return
				}

				rows, err := dbInfo.DB.Query(query.Query)
				check_err(err)
				defer rows.Close()

				var data []map[string]interface{}

				columns, err := rows.Columns()
				check_err(err)

				for rows.Next() {
					values := make([]interface{}, len(columns))
					valuePtrs := make([]interface{}, len(columns))
					for i := range values {
						valuePtrs[i] = &values[i]
					}

					err := rows.Scan(valuePtrs...)
					check_err(err)

					row := make(map[string]interface{})
					for i, column := range columns {
						row[column] = values[i]
					}
					data = append(data, row)
				}

				fmt.Println("Dữ liệu truy vấn", data)
				fmt.Println("Cấu trc câu truy vấn: ", query.Query)
				response := map[string]interface{}{
					"success": true,
					"data":    data,
				}
				json.NewEncoder(w).Encode(response)
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}
	}
}

func delete_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				check, status := Structure_query(query.Query, "delete")
				fmt.Println("Cấu trúc câu truy vấn: ", query.Query)

				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				if check {
					dbInfo.DB.Exec(query.Query)
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}
	}
}

func update_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				fmt.Println("Cấu trúc câu truy vấn: ", query.Query)
				check, status := Structure_query(query.Query, "update")
				response := map[string]interface{}{
					"success": check,
					"status":  status,
				}
				if check {
					dbInfo.DB.Exec(query.Query)
				}
				json.NewEncoder(w).Encode(&response)
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}
	}
}
