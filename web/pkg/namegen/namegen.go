package namegen

import (
	"crypto/md5"
	crypto "crypto/rand"
	"encoding/hex"
	"io"
	"math/big"
	"strconv"
)

func GetMD5Hash() string {
	h := md5.New()
	r := newCryptoRand()
	io.WriteString(h, strconv.FormatInt(r, 10))
	return hex.EncodeToString(h.Sum(nil))
}

func newCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}
