package actions

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mssola/user_agent"
	"github.com/zikwall/monit/src/protobuf/storage"
)

type HeatmapJSON struct {
	StreamID string `json:"stream_id"`
	UniqueID string `json:"unique_id"`
	Platform string `json:"platform"`
	App      string `json:"app"`
	Version  string `json:"version"`
	OS       string `json:"-"`
	Browser  string `json:"-"`
	Country  string `json:"-"`
	Region   string `json:"-"`
	Host     string `json:"host"`
}

func extractHeatmapRequestData(ctx *fiber.Ctx) ([]HeatmapJSON, string, error) {
	userAgentString := ctx.Get("User-Agent")
	if userAgentString == "" {
		return nil, "", errors.New("empty user agent")
	}

	var data []HeatmapJSON
	if err := ctx.BodyParser(&data); err != nil {
		return nil, "", err
	}

	return data, userAgentString, nil
}

func withUserAgent(header string, heatmap *HeatmapJSON) {
	if header == "" {
		return
	}
	ua := user_agent.New(header)
	browser, version := ua.Browser()
	if version != "" {
		browser = fmt.Sprintf("%s/%s", browser, version)
	}
	heatmap.Browser = browser
	heatmap.Platform = ua.Platform()
	heatmap.OS = ua.OS()
}

func (ht *HTTPController) Heatmap(ctx *fiber.Ctx) error {
	data, ua, err := extractHeatmapRequestData(ctx)
	if err != nil {
		return err
	}

	ip := fmt.Sprintf("%v", ctx.Locals("ip"))
	country, region, mxErr := ht.maxmind.Lookup(ip)

	for i := range data {
		withUserAgent(ua, &data[i])
		if mxErr == nil {
			data[i].Country = country
			data[i].Region = region
		}

		_, err := ht.storageClient.WriteHeatmap(ctx.Context(), &storage.HeatmapMessage{
			StreamID:  data[i].StreamID,
			UniqueID:  data[i].UniqueID,
			Platform:  data[i].Platform,
			App:       data[i].App,
			Version:   data[i].Version,
			OS:        data[i].OS,
			Browser:   data[i].Browser,
			Country:   data[i].Country,
			Region:    data[i].Region,
			Host:      data[i].Host,
			UserAgent: ua,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (ht *HTTPController) Event(_ *fiber.Ctx) error {
	return nil
}
