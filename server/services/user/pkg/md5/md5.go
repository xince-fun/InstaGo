package md5

import (
	"crypto/md5"
	"fmt"
	"github.com/google/wire"

	"github.com/xince-fun/InstaGo/server/services/user/conf"
	"strings"
)

var EncryptManagerSet = wire.NewSet(
	ProvideSalt,
	NewEncryptManager,
)

func ProvideSalt() Salt {
	return Salt(conf.GlobalServerConf.DBConfig.Salt)
}

type Salt string

type EncryptManager struct {
	salt Salt
}

func NewEncryptManager(salt Salt) *EncryptManager {
	return &EncryptManager{
		salt: salt,
	}
}

func (e *EncryptManager) EncryptPassword(code string) string {
	return e.md5Crypt(code, e.salt)
}

func (e *EncryptManager) md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
