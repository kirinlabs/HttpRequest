HttpRequest
=======
A simple `HTTP Request` package for golang. `GET` `POST` `DELETE` `PUT` `Upload`



### Installation
go get github.com/kirinlabs/HttpRequest


### How do we use HttpRequest?

Create request object
```go
req := HttpRequest.NewRequest()
req := HttpRequest.NewRequest().Debug(true).DisableKeepAlives(false).SetTimeout(5)
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
    "Connection": "keep-alive",
})

req.SetHeaders(map[string]string{
    "Source":"api",
})
```

Set cookies
```go
req.SetCookies(map[string]string{
    "name":"json",
    "token":"",
})

req.SetCookies(map[string]string{
    "age":"19",
})
```

Set timeout
```go
req.SetTimeout(5)  //default 30s
```

Object-oriented operation mode
```go
req := HttpRequest.NewRequest().
	Debug(true).
	SetHeaders(map[string]string{
	    "Content-Type": "application/x-www-form-urlencoded",
	}).SetTimeout(5)
res,err := HttpRequest.NewRequest().Get("http://127.0.0.1")
```

### GET

Query parameter
```go
res, err := req.Get("http://127.0.0.1:8000")
res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest")
res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest",nil)
res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest","address=beijing")

res, err := HttpRequest.Get("http://127.0.0.1:8000")
res, err := HttpRequest.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")
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
res, err := req.Post("http://127.0.0.1:8000")
res, err := req.Post("http://127.0.0.1:8000", "title=github&type=1")
res, err := req.JSON().Post("http://127.0.0.1:8000", "{\"id\":10,\"title\":\"HttpRequest\"}")
res, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "HttpRequest",
})
body, err := res.Body()
if err != nil {
    return
}
return string(body)

res, err := HttpRequest.Post("http://127.0.0.1:8000")
res, err := HttpRequest.JSON().Post("http://127.0.0.1:8000",map[string]interface{}{"title":"github"})
res, err := HttpRequest.Debug(true).SetHeaders(map[string]string{}).JSON().Post("http://127.0.0.1:8000","{\"title\":\"github\"}")
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
ReqBody: map[age:19 score:100]
-------------------------------------------------------------------
```


## Json
Post JSON request

Set header
```go
 req.SetHeaders(map[string]string{"Content-Type": "application/json"})
```
Or
```go
req.JSON().Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "github",
})

req.JSON().Post("http://127.0.0.1:8000", "{\"title\":\"github\",\"id\":10}")
```

Post request
```go
res, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "HttpRequest",
})
```

Print formatted JSON
```go
str, err := res.Export()
if err != nil {
   return
}
```

Unmarshal JSON
```go
var u User
err := res.Json(&u)
if err != nil {
   return err
}

var m map[string]interface{}
err := res.Json(&m)
if err != nil {
   return err
}
```

### Response

Response() *http.Response
```go
res, err := req.Post("http://127.0.0.1:8000/") //res is a http.Response object
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

Print formatted JSON
```go
str, err := res.Export()
if err != nil {
   return
}
```

Unmarshal JSON
```go
var u User
err := res.Json(&u)
if err != nil {
   return err
}

var m map[string]interface{}
err := res.Json(&m)
if err != nil {
   return err
}
```

Url() string
```go
res.Url()  //return the requested url
```

Headers() map[string]string
```go
res.Headers()  //return the response headers
```


### Advanced
GET
```go
import "github.com/kirinlabs/HttpRequest"
   
res,err := HttpRequest.Get("http://127.0.0.1:8000/")
res,err := HttpRequest.Get("http://127.0.0.1:8000/","title=github")
res,err := HttpRequest.Get("http://127.0.0.1:8000/?title=github")
res,err := HttpRequest.Debug(true).JSON().Get("http://127.0.0.1:8000/")
```

POST
```go
import "github.com/kirinlabs/HttpRequest"
   
res,err := HttpRequest.Post("http://127.0.0.1:8000/")
res,err := HttpRequest.SetHeaders(map[string]string{
	"title":"github",
}).Post("http://127.0.0.1:8000/")
res,err := HttpRequest.Debug(true).JSON().Post("http://127.0.0.1:8000/")
```


### Example
```go
import "github.com/kirinlabs/HttpRequest"
   
res,err := HttpRequest.Get("http://127.0.0.1:8000/")
res,err := HttpRequest.Get("http://127.0.0.1:8000/","title=github")
res,err := HttpRequest.Get("http://127.0.0.1:8000/?title=github")
res,err := HttpRequest.Get("http://127.0.0.1:8000/",map[string]interface{}{
	"title":"github",
})
res,err := HttpRequest.Debug(true).JSON().SetHeaders(map[string]string{
	"source":"api",
}).SetCookies(map[string]string{
	"name":"httprequest",
}).Post("http://127.0.0.1:8000/")


//Or
req := HttpRequest.NewRequest()
req := req.Debug(true).SetHeaders()
res,err := req.Debug(true).JSON().SetHeaders(map[string]string{
    "source":"api",
}).SetCookies(map[string]string{
    "name":"httprequest",
}).Post("http://127.0.0.1:8000/")
```