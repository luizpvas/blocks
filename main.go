package main

import (
	"fmt"
	"io/ioutil"

	"github.com/luizpvas/blocks/config"
	"github.com/luizpvas/blocks/endpoints"
)

func main() {
	data, err := ioutil.ReadFile("blocks.yaml")
	if err != nil {
		fmt.Printf("Could find `blocks.yaml` in the current directory.\n")
		return
	}

	appconfig, err := config.ParseAppConfig(data)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("You have registered %d resources.\n", len(appconfig.Resources))

	panic(endpoints.StartHTTPServer(appconfig))
}
