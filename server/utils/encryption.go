package utils

import "golang.org/x/crypto/bcrypt"

// BcryptHash 使用 bcrypt 对密码加密
func BcryptHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// BcryptCheck 验证密码
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
