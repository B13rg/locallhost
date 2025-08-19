# server

```go
import "github.com/b13rg/locallhost/pkg/server"
```

## Index

- [func Serve\(port int\)](<#Serve>)


<a name="Serve"></a>
## func [Serve](<https://github.com:b13rg/locallhost/blob/main/pkg/server/server.go#L35>)

```go
func Serve(port int)
```

Start serving on specified port. If port is \<= 0, defaults to 8080.