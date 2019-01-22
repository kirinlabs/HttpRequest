HttpRequest
=======
A simple `HTTP Request` package for golang. `GET` `POST` `DELETE` `PUT` `Upload`



### Installation
go get github.com/kirinlabs/HttpRequest


### How do we use HttpRequest?

Create request object
```go
req := HttpRequest.NewRequest()
```

Keep Alives
```go
req.DisableKeepAlives(false)
```

Ignore Https certificate validation
```go
req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
```

Set headers
```go
req.SetHeaders(map[string]string{
    "Content-Type": "application/x-www-form-urlencoded",
})
```

Set cookies
```go
req.SetCookies(map[string]string{
    "name":"json",
})
```

Set timeout
```go
req.SetTimeout(5)  //default 30s
```

Object-oriented operation mode
```go
req := HttpRequest.NewRequest().Debug(true).SetHeaders(map[string]string{
    "Content-Type": "application/x-www-form-urlencoded",
}).SetTimeout(5)
res,err := HttpRequest.NewRequest().Get("http://127.0.0.1:8000?id=10&title=HttpRequest",nil)
```

### GET

Query parameter
```go
res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest",nil)
```


Multi parameter,url will be rebuild to `http://127.0.0.1:8000?id=10&title=HttpRequest&name=jason&score=100`
```go
res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest",map[string]interface{}{
    "name":  "jason",
    "score": 100,
})
body, err := res.Body()
if err != nil {
    return
}
return string(body)
```


### POST

```go
res, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "HttpRequest",
})
body, err := res.Body()
if err != nil {
    return
}
return string(body)
```


### Upload
Params: url, filename, fileinput

```go
res, err := req.Upload("http://127.0.0.1:8000/upload", "/root/demo.txt","uploadFile")
body, err := res.Body()
if err != nil {
    return
}
return string(body)
```


### Debug
Default false

```go
req.Debug(true)
```

Print in standard output：
```go
[HttpRequest]
-------------------------------------------------------------------
Request: GET http://127.0.0.1:8000?name=iceview&age=19&score=100
Headers: map[Content-Type:application/x-www-form-urlencoded]
Cookies: map[]
Timeout: 30s
BodyMap: map[age:19 score:100]
-------------------------------------------------------------------
```


## Json
Post JSON request

Set header
```go
 req.SetHeaders(map[string]string{"Content-Type": "application/json"})
```

Post request
```go
res, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "HttpRequest",
})
```

Print JSON
```go
body, err := res.Json()
if err != nil {
   return
}
```

### Response

Response() *http.Response
```go
res, err := req.Post("http://127.0.0.1:8000/",nil) //res is a http.Response object
```

StatusCode() int
```go
res.StatusCode()
```

Body() ([]byte, error)
```go
body, err := res.Body()
log.Println(string(body))
```

Time() string
```go
res.Time()  //ms
```

Json() (string,error)
```go
body, err := res.Json() //Format the json return value
log.Println(body)
```

Url() string
```go
res.Url()  //return the requested url
```
