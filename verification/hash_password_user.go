package verification

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func HashPasswordUser(password string) string {

	salt := md5.Sum([]byte(os.Getenv("SALT")))
	saltHash := hex.EncodeToString(salt[:])
	hashPasswordSalt := password + saltHash

	hash := md5.Sum([]byte(hashPasswordSalt))
	return hex.EncodeToString(hash[:])
}
