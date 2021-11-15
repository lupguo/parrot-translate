package trans

import (
	"crypto/sha256"
	"fmt"
)

// HashText 取s的hash值返回
func HashText(s string) (hashID string, err error) {
	sha := sha256.New()
	_, err = sha.Write([]byte(s))
	if err != nil {
		return "HashID-500", err
	}
	return fmt.Sprintf("%x", sha.Sum(nil)), nil
}
