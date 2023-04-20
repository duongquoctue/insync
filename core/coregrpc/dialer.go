package coregrpc

import (
	"context"
	"time"

	"github.com/duongquoctue/insync/core/logger"
	"google.golang.org/grpc"
	grpcinsecure "google.golang.org/grpc/credentials/insecure"
)

const (
	defaultTimeout = time.Second * 20
)

func Dial(addr string, codes []string) map[string]*grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	dialMap := make(map[string]*grpc.ClientConn)
	for _, c := range codes {
		conn, err := grpc.DialContext(
			ctx,
			addr,
			grpc.WithBlock(),
			grpc.WithReturnConnectionError(),
			grpc.WithTransportCredentials(grpcinsecure.NewCredentials()),
		)
		if err != nil {
			logger.NewLoggerCtx(ctx).Fatalf("grpc dialer error: %v", err)
			continue
		}
		dialMap[c] = conn
	}

	return dialMap
}
