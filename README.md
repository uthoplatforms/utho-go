# Utho Go API client

[![GoDoc](https://godoc.org/github.com/uthoplatforms/utho-go?status.svg)](https://godoc.org/github.com/uthoplatforms/utho-go)

Godo is a Go client library for accessing the Utho API V2.

You can view the client API docs here: [https://pkg.go.dev/github.com/uthoplatforms/utho-go](https://pkg.go.dev/github.com/uthoplatforms/utho-go)

You can view Utho API V2 docs here: [https://utho.com/api-docs](https://utho.com/api-docs)

## Install
```sh
go get github.com/uthoplatforms/utho-go
```

## Usage

```go
import "github.com/uthoplatforms/utho-go/utho"
```

Create a new Utho client, then use the exposed services to
access different parts of the Utho API.

### Authentication

Currently, Personal Access Token (PAT) is the only method of
authenticating with the API. You can manage your tokens
at the Utho Control Panel [Applications Page](https://console.utho.com/api).

You can then use your token to create a new client:

```go
package main

import "github.com/uthoplatforms/utho-go/utho"

func main() {
    client, err := utho.NewClient("your-api-token")
}
```

## Examples


To create a new Cloud Instance:

```go
package main

import (
	"fmt"

	"github.com/uthoplatforms/utho-go/utho"
)

func main() {
	client, err := utho.NewClient("your-api-token")
	if err != nil {
		fmt.Printf("Something gone wrong: %s\n\n", err)
	}

	instanceName := "example-name"

	createRequest := utho.CreateCloudInstanceParams{
		Dcslug:       "innoida",
		Image:        "ubuntu-18.10-x86_64",
		Planid:       "10045",
		Billingcycle: "hourly",
		Cloud:        []utho.CloudHostname{utho.CloudHostname{Hostname: instanceName}},
	}

	newInstance, err := client.CloudInstances().Create(createRequest)
	if err != nil {
		fmt.Printf("Something gone wrong: %s\n\n", err)
	}

	fmt.Println(newInstance.ID)
}
```

## Versioning

Each version of the client is tagged and the version is updated accordingly.

To see the list of past versions, run `git tag`.


## Documentation

For details on all the functionality in this library, see the [GoDoc](http://godoc.org/github.com/uthoplatforms/utho-go) documentation OR the [API Docs](https://utho.com/api-docs).
