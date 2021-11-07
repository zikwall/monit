package clickhouse

import (
	"context"

	clickhousebuffer "github.com/zikwall/clickhouse-buffer"
)

type BufferAdapter struct {
	buffer clickhousebuffer.Client
}

func NewBuffer(ctx context.Context, cfg *clickhousebuffer.ClickhouseCfg) (*BufferAdapter, error) {
	b := &BufferAdapter{}
	ch, err := clickhousebuffer.NewClickhouseWithOptions(cfg)
	if err != nil {
		return nil, err
	}

	b.buffer = clickhousebuffer.NewClientWithOptions(ctx, ch,
		clickhousebuffer.DefaultOptions().SetFlushInterval(5000).SetBatchSize(5000),
	)
	return b, nil
}

func (b *BufferAdapter) Client() clickhousebuffer.Client {
	return b.buffer
}

func (b *BufferAdapter) Drop() error {
	b.buffer.Close()
	return nil
}

func (b *BufferAdapter) DropMsg() string {
	return "close clickhouse buffer adapter"
}
