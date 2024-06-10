package main

import (
	"fmt"
	"github.com/uthoplatforms/utho-go/utho"
	"os"
)

func main() {
	key := os.Getenv("API_KEY")
	client, err := utho.NewClient(key)
	if err != nil {
		panic(err)
	}

	res, err := client.Kubernetes().ReadTargetgroup("749299", "1234373")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", res.ID)
}
