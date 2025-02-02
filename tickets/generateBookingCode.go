package tickets

import (
	"crypto/rand"
	"math/big"
)

func GenerateTicketCode() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 7)

	for i := range code {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		code[i] = letters[num.Int64()]
	}

	for {
		if storeCode(string(code)) {
			break
		} else {
			storeCode(string(code))
		}
	}

	return string(code)

}
