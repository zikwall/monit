package service

import (
	"context"
	"github.com/zikwall/monit/src/pkg/drop"
)

type Impl struct {
	StorageClient *StorageClientImp
	*drop.Impl
}

type ImplOptions struct {
	StorageAddress string
}

func New(ctx context.Context, options *ImplOptions) (*Impl, error) {
	s := &Impl{}
	s.Impl = drop.NewContext(ctx)
	storageClientImpl, err := newStorageClient(ctx, options.StorageAddress)
	if err != nil {
		return nil, err
	}

	s.StorageClient = storageClientImpl

	s.AddDroppers(
		s.StorageClient,
	)

	return s, nil
}
