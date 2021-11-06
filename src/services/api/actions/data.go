package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/monit/src/protobuf/storage"
)

func (ht *HTTPController) Heatmap(ctx *fiber.Ctx) error {
	_, err := ht.storageClient.WriteHeatmap(ctx.Context(), &storage.HeatmapMessage{
		StreamID:  "",
		UniqueID:  "",
		Platform:  "",
		App:       "",
		Version:   "",
		Os:        "",
		Browser:   "",
		Region:    "",
		Host:      "",
		UserAgent: "",
	})
	return err
}

func (ht *HTTPController) Event(ctx *fiber.Ctx) error {
	return nil
}
