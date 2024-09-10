package main

import ()

// Hàm này kiểm tra cấu trúc của câu truy vấn
// This is function to check the structure of the query
func Structure_query(query string, Struct string) (bool, string) {
	words := split_words(query)
	if !check_query(words) {
		return false, "Truy vấn sai không hợp lệ"
	}
	if !check_condition_allowed(words) {
		return false, "Truy vấn vi phạm toán tử"
	}
	if !check_sructure_query(words, Struct) {
		return false, "Truy vấn sai cấu trúc"
	}
	return true, "Truy vấn hợp lệ"
}