package service

import (
	"context"
	clickhousebuffer "github.com/zikwall/clickhouse-buffer"
	"github.com/zikwall/monit/src/pkg/drop"
	"github.com/zikwall/monit/src/services/storage/clickhouse"
)

type Impl struct {
	Buffer *clickhouse.BufferAdapter
	*drop.Impl
}

type Options struct {
	Clickhouse *clickhousebuffer.ClickhouseCfg
}

func New(ctx context.Context, options *Options) (*Impl, error) {
	s := &Impl{}
	s.Impl = drop.NewContext(ctx)

	buffer, err := clickhouse.NewBuffer(ctx, options.Clickhouse)
	if err != nil {
		return nil, err
	}

	s.Buffer = buffer
	s.AddDroppers(
		s.Buffer,
	)
	return s, nil
}
