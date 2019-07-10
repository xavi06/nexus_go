# nexus_go

Nexus Repositories Golang API

#### Usage
```go
import (
    "fmt"
    "github.com/xavi06/nexus_go"
)

baseURL := "http://10.30.12.11:8081/service/rest"
id := "ZG9ja2VyLWZjY3M6Mjk2ZDhiZGViMmE2YWRiMzI4OTBjMTcwZTIzYzBjNWY"
api := gonexus.NewAPI(baseURL, "admin", "123456")
data, err := api.APIComponent(baseURL, id)
if err != nil {
    // error handler
}
fmt.Println(data)
```
