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
    res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest",nil)

    res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest",map[string]interface{}{
         "name":  "jason",
         "score": 100,
    })

    Rebuild url to http://127.0.0.1:8000?id=10&title=HttpRequest&name=jason&score=100
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

## Debug

```
  req := HttpRequest.NewRequest()
  req.Debug(true)       //Default false
```

```
  Print in standard output：

  [HttpRequest]
  -------------------------------------------------------------------
  Request: GET http://127.0.0.1:8000?name=iceview&age=19&score=100
  Headers: map[Content-Type:application/x-www-form-urlencoded]
  Cookies: map[]
  Timeout: 30s
  BodyMap: map[age:19 score:100]
  -------------------------------------------------------------------
```


## Request

```
  Get(url string, nil)

  Get(url string, body map[string]interface{})
```

```
  Post(url string, body map[string]interface{})
```

```
  Delete(url string, nil)

  Delete(url string, body map[string]interface{})
```

```
  Put(url string, body map[string]interface{})
```


## Public function

```
  SetHeaders(header map[string]string)

  SetCookies(header map[string]string)

  SetTimeout(d time.Duration)
```


## Response

```
  Response() *http.Response

  StatusCode() int

  Body() ([]byte, error)

  Time() string
```
