package common

import (
	"crypto/sha1"
	"encoding/base64"
)

func HashString(input string) string {

	hasher := sha1.New()
	bv := []byte(input)
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha

}

