package appstrings

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

//Random random type
type Random struct {
}

//NewRandom factory for Random
func NewRandom() *Random {
	rand.Seed(time.Now().UnixNano())
	return &Random{}
}

//Generate generates a random string
func (r *Random) Generate(length int) string {
	rnd := rand.Int63()
	rbytes := md5.Sum([]byte(fmt.Sprintf("%d", rnd)))
	s := fmt.Sprintf("%x", rbytes)
	if length != 0 {
		return s[0:length]
	}
	return s
}
