package main

import "github.com/oden7777/grpc/bakerapi/client"

func main() {
	client.Bake()
	client.Report()
}
