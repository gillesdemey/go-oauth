package lib

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"strings"
)

type pair struct {
	key   string
	value string
}

type pairs []pair

func GenerateBase(signMethod string, httpMethod string, baseUri string, params map[string]string) string {

	paramsString := generateParamsString(signMethod, params)

	httpMethod = strings.ToUpper(httpMethod)

	base := httpMethod + "&" + RFC3986(baseUri) + "&" + RFC3986(paramsString)

	return base
}

func generateParamsString(signMethod string, params map[string]string) string {

	pairs := make(pairs, 0)

	for k, v := range params {

		if v != "" {
			pairs = append(pairs, pair{key: k, value: v})
		}

	}

	oauth_params := map[string]string{
		"oauth_signature_method": signMethod,
		"oauth_version":          "1.0",
	}

	for k, v := range oauth_params {
		pairs = append(pairs, pair{key: k, value: v})
	}

	values := url.Values{}

	for _, pair := range pairs {
		values.Add(pair.key, pair.value)
	}

	paramsString := fixURLEncoding(values.Encode())

	return paramsString
}

func SHA1(message string, key string) string {

	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(message))
	hash := h.Sum(nil)

	base64signature := base64.StdEncoding.EncodeToString(hash)

	return base64signature
}

func RFC3986(str string) string {
	return fixURLEncoding(url.QueryEscape(str))
}

func fixURLEncoding(msg string) string {
	return strings.Replace(msg, "+", "%20", -1)
}
