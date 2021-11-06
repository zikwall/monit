package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/monit/src/protobuf/storage"
)

type HTTPController struct {
	storageClient storage.StorageClient
}
type Response fiber.Map

func NewHTTPController(storage storage.StorageClient) *HTTPController {
	ht := &HTTPController{storageClient: storage}
	return ht
}
