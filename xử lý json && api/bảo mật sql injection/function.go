package main

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Hàm này kiểm tra xem thuộc tính được truy vấn không
// This is function to check if the attribute is queried
func check_allowed(attribute_No_Allowed map[string]bool, column string) bool {
	return !attribute_No_Allowed[column]
}

// Hàm này chặn các thuộc tính không được truy vấn
// This is function to block attributes that are not allowed to query
func queryNotAllowed() map[string]bool {
	Attricbute := map[string]bool{
		"tokenID": true,
		"code":    true,
	}
	return Attricbute
}

// Hàm này cấp phép toán tử được phép sử dụng trong truy vấn
// This function allows the use of operators in a query
func Operator_Allowed() map[string]bool {
	operator := map[string]bool{
		">": true,
	}
	return operator
}

// Hàm này đặt các ký tự đặc biệt vào danh sách cấm truy vấn
// This function puts special characters into the list of prohibited queries
func add_black_list() map[string]bool {
	black_list := map[string]bool{
		";":     true,
		"(":     true,
		")":     true,
		"--":    true,
		"'":     true,
		"UNION": true,
		"%":     true,
		"_":     true,
		"LIKE":  true,
		"NULL":  true,
		"OR":    true,
	}
	return black_list
}

// Hàm này chia một chuỗi theo từng từ
// This function splits a string into words
func split_words(query string) []string {

	black_list := add_black_list()
	for blackListItem := range black_list {
		query = strings.ReplaceAll(query, blackListItem, "")
	}

	words := strings.Fields(query)

	return words
}

// Hàm này kiểm tra câu truy vấn hợp lệ
// This is function to check the query is valid
func check_query(words []string) bool {
	return len(words) > 2
}

func list_query() map[string]bool {
	words := map[string]bool{
		"SELECT": true,
		"DELETE": true,
		"INSERT": true,
		"UPDATE": true,
	}
	return words
}

func check_structure(word string, Struct string, i int, list map[string]bool) bool {
	if word == Struct && i == 0 && list[word] {
		return true
	} else if list[word] && i != 0 {
		return false
	}
	return true
}

// Hàm này kiểm tra cấu trúc câu truy vấn hợp lệ
// This is function to check the structure of the query is valid
func check_sructure_query(words []string, Struct string) bool {
	list := list_query()
	Struct = strings.ToUpper(Struct)
	for i, word := range words {
		word = strings.ToUpper(word)
		if !check_structure(word, Struct, i, list) {
			return false
		}
	}
	return true
}

// Hàm này kiểm tra xem các thuộc tính và toán tử có được truy vấn có được cho phép hay không
// This function checks if the attributes and operators are allowed to be queried
func check_condition_allowed(words []string) bool {
	for _, word := range words {
		if !check_allowed(queryNotAllowed(), word) {
			return false
		} else if !check_allowed(Operator_Allowed(), word) {
			return false
		}
	}
	return true
}
