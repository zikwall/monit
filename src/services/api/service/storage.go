package service

import (
	"github.com/zikwall/monit/src/protobuf/storage"
	"google.golang.org/grpc"
)

type StorageClientImp struct {
	conn   *grpc.ClientConn
	client storage.StorageClient
}

func newStorageClient(listenAddress string) (*StorageClientImp, error) {
	impl := &StorageClientImp{}
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(listenAddress, opts...)
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
