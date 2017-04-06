# Compose is a Go client library fot the IBM Compose API

This is currently work in progress so be aware of breaking changes. 

## Example usage:

### Providion mongodb

Export your compose token as env variable
```
export COMPOSE_TOKEN="mytoken"
```

```
package main

import (
        "fmt"
        "log"

        "github.com/combor/compose"
)

func main() {
        token, err := compose.GetApiToken()
        if err != nil {
                log.Fatal(err)
        }
        apiURL := compose.GetApiURL()
        c := compose.NewClient(token, apiURL)
        accounts, err := c.GetAccounts()
        if err != nil {
                log.Println(err)
        }
        accountId := accounts[0].Id

        deployment, err := c.CreateDeployment(accountId, "testmongo", "mongodb", "aws:eu-west-1", "", 1, true, true)
        if err != nil {
                log.Println(err)
        }
        fmt.Printf("%s\n", deployment.CreatedAt)

}
```

### Scale existing deployment
```
package main

import (
        "fmt"
        "log"
        "net/url"

        "github.com/combor/compose"
)

func main() {
        token := "mytoken"
        apiURL, err := url.Parse("https://api.compose.io/2016-07")
        if err != nil {
                log.Fatal(err)
        }
        c := compose.NewClient(token, apiURL)

        deployments, err := c.GetDeployments()
        if err != nil {
                log.Fatal(err)
        }
        deploymentId := deployments[0].Id
        scale, err := c.ScaleDeployment(deploymentId, 7)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%s\n", scale.Status)
}
```

### Delete all deployments

```
package main

import (
        "fmt"
        "log"

        "github.com/combor/compose"
)

func main() {
        token, err := compose.GetApiToken()
        if err != nil {
                log.Fatal(err)
        }
        apiURL := compose.GetApiURL()
        c := compose.NewClient(token, apiURL)

        deployments, err := c.GetDeployments()
        if err != nil {
                log.Println(err)
        }
        for _, deployment := range deployments {
                recipe, err := c.DeleteDeployment(deployment.Id)
                if err != nil {
                        log.Println(err)
                }
                fmt.Printf("%#v\n", recipe.Status)
        }

}
``` 
