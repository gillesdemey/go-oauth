/**
 * Reference material:
 * https://dev.twitter.com/docs/auth/creating-signature
 */

package oauth

import (
	"fmt"
	"github.com/gillesdemey/go-oauth/lib"
)

/**
 * The sign function will call the internal HMAC or RSA functions
 */
func Sign(signMethod string, httpMethod string, baseUri string, params map[string]string, consumerSecret string, tokenSecret string) string {

	signature := ""

	switch signMethod {

	default:
		fmt.Println("signing method unsupported.")

	case "HMAC-SHA1":
		base := lib.GenerateBase(signMethod, httpMethod, baseUri, params)
		signature = signHMAC(base, consumerSecret, tokenSecret)
		break

	}

	return signature

}

func signHMAC(base string, consumerSecret string, tokenSecret string) string {
	key := lib.RFC3986(consumerSecret) + "&" + lib.RFC3986(tokenSecret)
	return lib.SHA1(base, key)
}
