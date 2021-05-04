package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/oden7777/grpc/imageapi/api/gen/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const port = 50051

func Upload(filePath string) error {
	address := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrapf(err, "コネクションエラー")
	}
	defer conn.Close()

	client := pb.NewImageUploadServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrapf(err, "ファイル展開失敗")
	}
	defer file.Close()

	stream, err := client.Upload(ctx)
	if err != nil {
		return errors.Wrapf(err, "リクエストの送信失敗")
	}

	err = uploadMetaData(stream, file)
	if err != nil {
		return errors.Wrapf(err, "メタデータのリクエスト処理が失敗")
	}

	err = uploadFile(stream, file)
	if err != nil {
		return errors.Wrapf(err, "イメージののアップロード処理が失敗")
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		return errors.Wrapf(err, "レスポンス取得失敗")
	}

	log.Printf("リクエスト結果： %s", reply.String())
	return nil
}

func uploadMetaData(stream pb.ImageUploadService_UploadClient, file *os.File) error {
	metaDataRequest := &pb.ImageUploadRequest{File: &pb.ImageUploadRequest_FileMeta_{
		FileMeta: &pb.ImageUploadRequest_FileMeta{
			Filename: file.Name(),
		},
	}}
	err := stream.Send(metaDataRequest)
	if err != nil {
		return errors.Wrapf(err, "メタデータのリクエストの送信失敗")
	}
	return nil
}

func uploadFile(stream pb.ImageUploadService_UploadClient, file *os.File) error {
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		//fmt.Printf("sent : %d\n", n)
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrapf(err, "ファイル読み込み失敗")
		}
		fileDateRequest := &pb.ImageUploadRequest{
			File: &pb.ImageUploadRequest_Data{
				Data: buf[:n],
			},
		}
		err = stream.Send(fileDateRequest)
		if err != nil {
			return errors.Wrapf(err, "イメージデータのリクエストの送信失敗")
		}
	}
	return nil
}
