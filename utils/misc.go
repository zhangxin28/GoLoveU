package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"reflect"
	"strings"
)

// IsEmpty shows whether a valus is empty
func IsEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

// Contains indicates whether the target contains the search
func Contains(search interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == search {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(search)).IsValid() {
			return true
		}
	}
	return false
}

// ContainsIgnoreCase represents the target contains the search, ignore case
func ContainsIgnoreCase(search string, target []string) bool {
	if len(search) == 0 {
		return false
	}

	if len(target) == 0 {
		return false
	}
	search = strings.ToLower(search)
	for i := 0; i < len(target); i++ {
		if strings.ToLower(target[i]) == search {
			return true
		}
	}
	return false
}

// EncodePassword returns the bcrypt hash of the password at the given
// cost. If the cost given is less than MinCost, the cost will be set to
// DefaultCost, instead. Use CompareHashAndPassword, as defined in this package,
// to compare the returned hashed password with its cleartext version.
func EncodePassword(rawPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

// ValidatePassword compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns nil on success, or an error on failure.
func ValidatePassword(encodePassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(inputPassword))
	return err == nil
}

// WriteString 写入内容
func WriteString(path string, content string, append bool) error {
	flag := os.O_RDWR | os.O_CREATE
	if append {
		flag = flag | os.O_APPEND
	}
	file, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

// AppendLine 追加行
func AppendLine(path, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	content = strings.Join([]string{content, "\n"}, "")
	_, err = file.WriteString(content)
	return err
}
