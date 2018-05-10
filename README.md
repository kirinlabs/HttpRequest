# HttpRequest
A simple HTTP Request package for golang.


## Installation

```
go get github.com/iceview/HttpRequest
```

## Example

```
package main

import (
	"github.com/iceview/HttpRequest"
	"fmt"
)

func main() {
	req := HttpRequest.NewRequest()

    req.SetHeaders(map[string]string{
    	"Content-Type": "application/x-www-form-urlencoded",
    })

    req.SetCookies(map[string]string{
    	"name":      "iceview",
    	"sessionid": "NFLJASDLFIWRLKVLZXJLZVJALFFNASLLXNMVNALWSDFJF3WO",
    })

    req.SetTimeout(5)  //default 30s

    /*
    res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest")
    */

    res, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    	"id":    10,
    	"title": "HttpRequest",
    })

    if err != nil {
    	log.Println(err)
    	return
    }

    body, err := res.Body()
    if err != nil {
    	log.Println(err)
    	return
    }

    fmt.Println(string(body))
}
```

## Request

```
Get(url string)

Post(url string,body map[string]interface{})

Delete(url string)

Put(url string,body map[string]interface{})
```


## Public function

```
SetHeaders(header map[string]string)

SetCookies(header map[string]string)

SetTimeout(d time.Duration)
```


## Response

```
Response()

StatusCode()

Body()
```
