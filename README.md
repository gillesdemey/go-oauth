OAuth 1.0 HMAC-SHA1 signing
========

## Install

`go get github.com/gillesdemey/go-oauth`

## Usage

```go

import(
  "github.com/gillesdemey/go-oauth"
)

params := map[string]string {
    "message"                 : "Hello Ladies + Gentlemen, a signed OAuth request!",
    
    // any additional headers for your request
    
    "oauth_nonce"             : "randomGeneratedString",
    "oauth_timestamp"         : "1318622958",
  
    "oauth_consumer_key"      : "myConsumerKey",
    "oauth_token"             : "myOAuthToken",
  }
  
  signature := Sign("HMAC-SHA1", "POST", "https://api.twitter.com/1/statuses/update.json", params, "consumerSecret", "tokenSecret")
  
  // will return your signed signature (eg. "tnnArxj06cWHq44gCs1OSKk/jLY=")
```

## Test

```
$: go test
PASS
ok    github.com/gillesdemey/go-oauth 0.008s
```

## Todo

Write some additional tests
