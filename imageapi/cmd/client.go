package main

import (
	"os"

	"github.com/oden7777/grpc/imageapi/api/client"
)

func main() {
	// ファイルパス情報
	arg := os.Args[1]
	client.Upload(arg)
}
