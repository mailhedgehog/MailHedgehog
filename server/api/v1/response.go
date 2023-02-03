package v1

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Response struct {
	Message string
	Data    interface{}
	Meta    interface{}
	Status  int
}

func (response *Response) Send(ctx *fiber.Ctx) error {
	respMap := make(fiber.Map)

	status := http.StatusOK
	if response.Status > 0 {
		status = response.Status
	}

	if response.Message != "" {
		respMap["message"] = response.Message
	}

	if response.Data != nil {
		respMap["data"] = response.Data
	}

	if response.Data != nil {
		respMap["meta"] = response.Meta
	}

	return ctx.Status(status).JSON(respMap)
}
