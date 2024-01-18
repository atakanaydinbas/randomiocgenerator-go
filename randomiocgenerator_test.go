package generator_test

import (
	"net"
	"regexp"
)

func ValidateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}

func ValidateDomain(domain string) bool {
	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)
	return domainRegex.MatchString(domain)
}
