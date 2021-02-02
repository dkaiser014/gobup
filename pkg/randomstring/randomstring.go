package randomstring

import (
	"math/rand"
	"strings"
	"time"
)

// RandomString returns a random string
func RandomString() string {
	rand.Seed(time.Now().Unix())

	// Lowercase and Uppercase
	var output strings.Builder
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%&/()=?¿¡"
	length := 20

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	return output.String()
}
