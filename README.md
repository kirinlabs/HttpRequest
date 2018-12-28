# HttpRequest
A simple `HTTP Request` package for golang. `GET` `POST` `DELETE` `PUT`



# Installation
```
go get github.com/kirinlabs/HttpRequest
```

# Example
```go
package main

import (
    "github.com/kirinlabs/HttpRequest"
)

func main() {

    // Generate request object
    req := HttpRequest.NewRequest()

    // The default value is false
    req.DisableKeepAlives(false)

    // Ignore Https certificate validation
    req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})

    // Set headers
    req.SetHeaders(map[string]string{
    	"Content-Type": "application/x-www-form-urlencoded",
    })

    // Set cookies
    req.SetCookies(map[string]string{
    	"name":      "json",
    })

    // Set timeout
    req.SetTimeout(5)  //default 30s

}
```

# Object Example

```go
package main

import (
    "github.com/kirinlabs/HttpRequest"
)

func main() {
    // Object-oriented operation mode
    req := HttpRequest.NewRequest().Debug(true).SetHeaders(map[string]string{
           "Content-Type": "application/x-www-form-urlencoded",
    }).SetTimeout(5)
}
```

## Get
```go
  // Query parameter
  res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest",nil)

  // Multi parameter,url will be rebuild
  res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest",map[string]interface{}{
       "name":  "jason",
       "score": 100,
  })
  // Will rebuild url to `http://127.0.0.1:8000?id=10&title=HttpRequest&name=jason&score=100`

  body, err := res.Body()
  if err != nil {
     log.Println(err)
     return
  }

  return string(body)

```


## Post
```go
  // Send post request
  res, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
      	"id":    10,
      	"title": "HttpRequest",
      })

  body, err := res.Body()
  if err != nil {
     log.Println(err)
     return
  }

  return string(body)
```



## Debug
```go
  req := HttpRequest.NewRequest()

  // Default false
  req.Debug(true)


  // Print in standard output：
  [HttpRequest]
  -------------------------------------------------------------------
  Request: GET http://127.0.0.1:8000?name=iceview&age=19&score=100
  Headers: map[Content-Type:application/x-www-form-urlencoded]
  Cookies: map[]
  Timeout: 30s
  BodyMap: map[age:19 score:100]
  -------------------------------------------------------------------
```


## Send and print json
```go

  // Set header
  req.SetHeaders(map[string]string{
      "Content-Type": "application/json",
  })


  // Print json format
  body, err := res.Json()
  if err != nil {
     log.Println(err)
     return
  }

  print body

```

## Public function

```go

  // Request

  NewRequest()

  Debug(flag bool)

  SetHeaders(header map[string]string)

  SetCookies(header map[string]string)

  SetTimeout(d time.Duration)

  DisableKeepAlives(flag bool)

  SetTLSClient(v *tls.Config)


  // Response

  Response() *http.Response

  StatusCode() int

  Body() ([]byte, error)

  Time() string

  Json() (string,error)

  Url() string

```
