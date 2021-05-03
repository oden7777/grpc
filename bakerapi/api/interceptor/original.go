package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func MyCustomInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		fmt.Printf("MyCustomInterceptor: request is [%v]\n", req)
		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		fmt.Printf("MyCustomInterceptor: response is [%v]\n", res)
		return res, nil
	}
}
