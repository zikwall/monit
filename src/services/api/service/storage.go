package service

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/zikwall/monit/src/protobuf/storage"
)

type StorageClientImp struct {
	conn   *grpc.ClientConn
	client storage.StorageClient
}

func newStorageClient(ctx context.Context, listenAddress string) (*StorageClientImp, error) {
	impl := &StorageClientImp{}
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(ctx, listenAddress, opts...)
	if err != nil {
		return nil, err
	}

	impl.client = storage.NewStorageClient(conn)
	impl.conn = conn
	return impl, nil
}

func (sc *StorageClientImp) Client() storage.StorageClient {
	return sc.client
}

func (sc *StorageClientImp) Drop() error {
	return sc.conn.Close()
}

func (sc *StorageClientImp) DropMsg() string {
	return "close storage client connection"
}
