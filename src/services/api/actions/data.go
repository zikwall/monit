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

func withUserAgent(header string, heatmap HeatmapJSON) HeatmapJSON {
	if header == "" {
		return heatmap
	}
	ua := user_agent.New(header)
	browser, version := ua.Browser()
	if version != "" {
		browser = fmt.Sprintf("%s/%s", browser, version)
	}
	heatmap.Browser = browser
	heatmap.Platform = ua.Platform()
	heatmap.OS = ua.OS()
	return heatmap
}

func (ht *HTTPController) Heatmap(ctx *fiber.Ctx) error {
	data, ua, err := extractHeatmapRequestData(ctx)
	if err != nil {
		return err
	}

	ip := fmt.Sprintf("%v", ctx.Locals("ip"))
	country, region, mxErr := ht.maxmind.Lookup(ip)

	for _, heatmap := range data {
		heatmap = withUserAgent(ua, heatmap)
		if mxErr == nil {
			heatmap.Country = country
			heatmap.Region = region
		}

		_, err := ht.storageClient.WriteHeatmap(ctx.Context(), &storage.HeatmapMessage{
			StreamID:  heatmap.StreamID,
			UniqueID:  heatmap.UniqueID,
			Platform:  heatmap.Platform,
			App:       heatmap.App,
			Version:   heatmap.Version,
			OS:        heatmap.OS,
			Browser:   heatmap.Browser,
			Country:   heatmap.Country,
			Region:    heatmap.Region,
			Host:      heatmap.Host,
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
