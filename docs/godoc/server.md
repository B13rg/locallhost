# server

```go
import "github.com/b13rg/locallhost/pkg/server"
```

## Index

- [func Serve\(port int\)](<#Serve>)
- [type RequestResponse](<#RequestResponse>)
  - [func ExtractRequestData\(req \*http.Request\) \*RequestResponse](<#ExtractRequestData>)


<a name="Serve"></a>
## func [Serve](<https://github.com:b13rg/locallhost/blob/main/pkg/server/server.go#L74>)

```go
func Serve(port int)
```

Start serving on specified port.

<a name="RequestResponse"></a>
## type [RequestResponse](<https://github.com:b13rg/locallhost/blob/main/pkg/server/index-tmpl.go#L6-L17>)

Contains the data extracted from the request and returned to the user.

```go
type RequestResponse struct {
    // IP address of the remote client
    RemoteAddr string
    // Port of the remote client
    RemotePort string
    // Request method
    Method string
    // Request protocol
    Proto string
    // Request headers
    Header http.Header
}
```

<a name="ExtractRequestData"></a>
### func [ExtractRequestData](<https://github.com:b13rg/locallhost/blob/main/pkg/server/server.go#L49>)

```go
func ExtractRequestData(req *http.Request) *RequestResponse
```

