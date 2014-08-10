package oauth

import (
  // "fmt"
  // "time"
  "testing"
  // "strconv"
  // "github.com/nu7hatch/gouuid"
)

func TestSignature(t *testing.T) {

  params := map[string]string {
    "status"                  : "Hello Ladies + Gentlemen, a signed OAuth request!",
    "include_entities"        : "true",
    "oauth_nonce"             : "kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg",
    "oauth_timestamp"         : "1318622958",

    "oauth_consumer_key"      : "xvz1evFS4wEEPTGEFPHBog",
    "oauth_token"             : "370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb",

    "oauth_version"           : "1.0",
  }

  signature := Sign("HMAC-SHA1", "POST", "https://api.twitter.com/1/statuses/update.json", params, "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw", "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE")

  if signature != "tnnArxj06cWHq44gCs1OSKk/jLY=" {
    t.Error("invalid signature generated")
  }

}

// func TestAuthorizationHeader(t *testing.T) {

//   u, _ := uuid.NewV4()

//   params := map[string]string {
//     "oauth_nonce"             : u.String(),
//     "oauth_timestamp"         : strconv.FormatInt(time.Now().Unix(), 10),

//     "oauth_consumer_key"      : "MyConsumerKey",
//     "oauth_consumer_secret"   : "SuperSecret",

//     "oauth_token_secret"      : "TokenSecret",
//     "oauth_token"             : "Token",
//   }

//   signature := Sign("HMAC-SHA1", "GET", "https://api/yourwebservice.com/", params, params["oauth_consumer_secret"], params["oauth_token_secret"])

//   fmt.Printf("OAuth oauth_consumer_key=\"%s\", oauth_nonce=\"%s\", oauth_signature=\"%s\", oauth_signature_method=\"%s\", oauth_timestamp=\"%s\", oauth_version=\"%s\"\n", params["oauth_consumer_key"], params["oauth_nonce"], signature, "HMAC-SHA1", params["oauth_timestamp"], "1.0")

// }