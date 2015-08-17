package main

import (
	"fmt"
	"os"
	"path/filepath"
	// "runtime"
	// "strings"

	"github.com/cloud66/cloud66"
)

var (
	client       cloud66.Client
	debugMode    bool   = false
	VERSION      string = "dev"
	tokenFile    string = "cloud66.json"
	tokenDir     string = "/Users/babak/"
	fayeEndpoint string = "https://sockets.cloud66.com:443/push"
)

func main() {
	// debugMode = os.Getenv("CLOUD66_DEBUG")
	getClient()

	var stacks []cloud66.Stack

	stacks, err := client.StackList()

	fmt.Printf("%s, %s", stacks, err)
}

func getClient() {
	// is there a token file?
	_, err := os.Stat(filepath.Join(tokenDir, tokenFile))
	if err != nil {
		fmt.Println("No previous authentication found.")
		cloud66.Authorize(tokenDir, tokenFile)
		os.Exit(1)
	} else {
		client = cloud66.GetClient(tokenDir, tokenFile, "")
		client.Debug = debugMode
	}
}
