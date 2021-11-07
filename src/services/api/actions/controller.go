package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/monit/src/protobuf/storage"
	"github.com/zikwall/monit/src/services/api/service"
)

type HTTPController struct {
	storageClient storage.StorageClient
	maxmind       *service.Maxmind
}
type Response fiber.Map

func NewHTTPController(storage storage.StorageClient, maxmind *service.Maxmind) *HTTPController {
	ht := &HTTPController{storageClient: storage, maxmind: maxmind}
	return ht
}
