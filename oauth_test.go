package oauth

import (
	"testing"
)

func TestSignature(t *testing.T) {

	params := map[string]string{
		"status":           "Hello Ladies + Gentlemen, a signed OAuth request!",
		"include_entities": "true",
		"oauth_nonce":      "kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg",
		"oauth_timestamp":  "1318622958",

		"oauth_consumer_key": "xvz1evFS4wEEPTGEFPHBog",
		"oauth_token":        "370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb",
	}

	signature := Sign("HMAC-SHA1", "POST", "https://api.twitter.com/1/statuses/update.json", params, "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw", "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE")

	if signature != "tnnArxj06cWHq44gCs1OSKk/jLY=" {
		t.Error("invalid signature generated")
	}

}
