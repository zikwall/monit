package service

import (
	"context"

	"github.com/zikwall/monit/src/pkg/drop"
)

type Impl struct {
	StorageClient *StorageClientImp
	Maxmind       *Maxmind
	*drop.Impl
}

type ImplOptions struct {
	StorageAddress string
	MaxmindOptions *MaxmindOptions
}

func New(ctx context.Context, options *ImplOptions) (*Impl, error) {
	s := &Impl{}
	s.Impl = drop.NewContext(ctx)

	storageClientImpl, err := newStorageClient(ctx, options.StorageAddress)
	if err != nil {
		return nil, err
	}

	maxmind, err := newMaxmind(options.MaxmindOptions)
	if err != nil {
		return nil, err
	}

	s.StorageClient = storageClientImpl
	s.Maxmind = maxmind
	s.AddDroppers(
		s.StorageClient,
		s.Maxmind,
	)

	return s, nil
}
