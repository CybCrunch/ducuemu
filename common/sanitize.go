package common

import (
	"strings"
	"fmt"
)

func NonAscii(s string) bool {

	f := func(r rune) bool {
		return r < 'A' || r > 'z'
	}

	if strings.IndexFunc(s, f) != -1 {
		return true
	} else {
		return false
	}

}

