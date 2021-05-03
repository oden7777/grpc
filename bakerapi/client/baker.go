package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oden7777/grpc/bakerapi/api/gen/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const port = 50051

func Bake() error {
	address := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrapf(err, "コネクションエラー")
	}
	defer conn.Close()

	client := api.NewPancakeBakerServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()
	md := metadata.New(map[string]string{"authorization": "Bearer hi/mi/tsu"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	bakeRequest := &api.BakeRequest{Menu: api.Pancake_MIX_BERRY}
	reply, err := client.Bake(ctx, bakeRequest)
	if err != nil {
		return errors.Wrapf(err, "リクエスト失敗")
	}
	log.Printf("リクエスト結果： %s", reply.GetPancake())
	return nil
}

func Report() error {
	address := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrapf(err, "コネクションエラー")
	}
	defer conn.Close()

	client := api.NewPancakeBakerServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()
	md := metadata.New(map[string]string{"authorization": "Bearer hi/mi/tsu"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	reportRequest := &api.ReportRequest{}
	reply, err := client.Report(ctx, reportRequest)
	if err != nil {
		return errors.Wrapf(err, "リクエスト失敗")
	}
	log.Printf("リクエスト結果： %s", reply.GetReport())
	return nil
}
