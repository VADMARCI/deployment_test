package hash

import "github.com/dchest/uniuri"

type Hash struct{}

func (h Hash) GenerateHash(length int) string {
	return uniuri.NewLen(length)
}
