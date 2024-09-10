package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-mail/mail"
)

var codeStore = make(map[string]string)
var codeExpiry = make(map[string]time.Time)

const charset = "234789adfghiSWXYZ"
const tokenLength = 6

var seededRandPkg *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRandPkg.Intn(len(charset))]
	}
	return string(b)
}

func randomToken(length int) string {
	return StringWithCharset(length, charset)
}

func Send_Email(sender string, password string, receiver string) (bool, string) {
	m := mail.NewMessage()
	token := randomToken(tokenLength)

	m.SetHeader("From", sender)
	m.SetHeader("To", receiver)
	m.SetBody("text/plain", fmt.Sprintf("mã xác nhận của bạn là: %s", token))
	m.SetHeader("Subject", "Cảm ơn bạn đã xem")

	d := mail.NewDialer("smtp.gmail.com", 587, sender, password)
	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		log.Printf("\n Sender :%s", sender)
		log.Printf("\n Password:%s", password)
		log.Printf("\n Receiver:%s", receiver)

		return false, "Lỗi kết nối email từ thượng nguồn"
	}

	codeStore[receiver] = token
	codeExpiry[receiver] = time.Now().Add(5 * time.Minute)
	return true, "Đã lưu trữ mã xác nhận vào thượng nguồn"
}

func verify_email(email string, code string) (bool, string) {
	storedCode, ok := codeStore[email]
	if !ok {
		return false, "Lỗi mã lưu trữ mã xác nhận"
	}

	if time.Now().After(codeExpiry[email]) {
		delete(codeStore, email)
		delete(codeExpiry, email)
		return false, "Mã xác nhận đã hết hạn"
	}

	if storedCode != code {
		return false, "Mã xác nhận không đúng"
	}

	delete(codeStore, email)
	delete(codeExpiry, email)

	return true, "Xác nhận thành công"
}
