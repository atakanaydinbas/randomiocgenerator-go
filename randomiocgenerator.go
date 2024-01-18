package generator

import (
	"fmt"
	"math/rand"
)

// GenerateRandomIPv4 generates a random IPv4 address.
func GenerateRandomIPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
	)
}

// GenerateRandomDomain generates a random domain name.
func GenerateRandomDomain() string {
	charSet := "abcdefghijklmnopqrstuvwxyz0123456789"
	randomDomain := make([]byte, 10)
	for i := range randomDomain {
		randomDomain[i] = charSet[rand.Intn(len(charSet))]
	}
	return fmt.Sprintf("%s.com", string(randomDomain))
}
