package regutil

import (
	"encoding/base32"
	"regexp"
	"strings"

	base58 "github.com/jbenet/go-base58"
)

// DockerizeHash does base58 to base32 conversion
func DockerizeHash(base58Hash string) string {
	re := regexp.MustCompile(`(/ipfs/)?(.*)`)
	matches := re.FindStringSubmatch(base58Hash)
	base58Hash = matches[len(matches)-1]
	decodedB58 := base58.Decode(base58Hash)
	b32str := base32.StdEncoding.EncodeToString(decodedB58)

	end := len(b32str)
	if end > 0 {
		end = end - 1
	}

	// remove padding
	return strings.ToLower(b32str[0:end])
}

// IpfsifyHash does base32 to base58 conversion
func IpfsifyHash(base32Hash string) string {
	decodedB32, err := base32.StdEncoding.DecodeString(strings.ToUpper(base32Hash) + "=")
	if err != nil {
		return ""
	}

	return base58.Encode(decodedB32)
}
