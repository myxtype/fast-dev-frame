package sqltypes

import (
	"database/sql/driver"
	"frame/pkg/crypto/md5"
	"frame/pkg/randstr"
	"frame/pkg/sql/format"
)

type Password struct {
	Hash string `json:"h"`
	Salt string `json:"s"`
}

func NewPassword(password string) Password {
	salt := randstr.Hex(8)
	hash := md5.MustEncryptString(salt + password)
	return Password{
		Hash: hash,
		Salt: salt,
	}
}

func (t Password) GormDataType() string {
	return "json"
}

func (t *Password) Scan(src interface{}) error {
	return format.Scan(t, src)
}

func (t Password) Value() (driver.Value, error) {
	return format.Value(t)
}

func (t Password) Check(password string) bool {
	hash, _ := md5.EncryptString(t.Salt + password)
	return hash == t.Hash
}
