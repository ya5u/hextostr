package hextostr

import (
	"encoding/hex"
	"strings"
)

// HexToStr convert hex to ascii
func HexToStr(h string) (string, error) {
	b, err := hex.DecodeString(h)
	if err != nil {
		return "", err
	}
	return string(b), err
}

// StrToHex convert ascii to hex
func StrToHex(s string, length int) string {
	h := hex.EncodeToString([]byte(s))
	bl := length * 2
	if bl > len(h) {
		h = h + strings.Repeat("0", bl-len(h))
	} else if bl > 0 && bl < len(h) {
		h = h[0:bl]
	}
	return h
}
