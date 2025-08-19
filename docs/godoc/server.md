# server

```go
import "github.com/b13rg/locallhost/pkg/server"
```

## Index

- [func IndexTemplateString\(\) string](<#IndexTemplateString>)
- [func Serve\(port int\)](<#Serve>)
- [type RequestResponse](<#RequestResponse>)
  - [func ExtractRequestData\(req \*http.Request\) \*RequestResponse](<#ExtractRequestData>)


<a name="IndexTemplateString"></a>
## func [IndexTemplateString](<https://github.com:b13rg/locallhost/blob/main/pkg/server/index-tmpl.go#L24>)

```go
func IndexTemplateString() string
```

Returns the html template string for the index page.

<a name="Serve"></a>
## func [Serve](<https://github.com:b13rg/locallhost/blob/main/pkg/server/server.go#L76>)

```go
func Serve(port int)
```

Start serving on specified port.

<a name="RequestResponse"></a>
## type [RequestResponse](<https://github.com:b13rg/locallhost/blob/main/pkg/server/index-tmpl.go#L6-L19>)

Contains the data extracted from the request and returned to the user.

```go
type RequestResponse struct {
    // IP address of the remote client
    RemoteAddr string `json:"remoteAddress"`
    // Port of the remote client
    RemotePort string `json:"remotePort"`
    // Host header
    Host string `json:"host"`
    // Request method
    Method string `json:"method"`
    // Request protocol
    Proto string `json:"protocol"`
    // Request headers
    Header http.Header `json:"header"`
}
```

<a name="ExtractRequestData"></a>
### func [ExtractRequestData](<https://github.com:b13rg/locallhost/blob/main/pkg/server/server.go#L48>)

```go
func ExtractRequestData(req *http.Request) *RequestResponse
```

