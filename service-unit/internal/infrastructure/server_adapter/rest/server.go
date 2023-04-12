package rest

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hanapedia/the-bench/service-unit/internal/domain/core"
)

// must implement core.ServerAdapter
type RestServerAdapter struct {
	addr   string
	server *fiber.App
}

func NewRestServerAdapter() RestServerAdapter {
	return RestServerAdapter{addr: ":8080", server: fiber.New()}
}

func (rsa RestServerAdapter) Serve() error {
	return rsa.server.Listen(rsa.addr)
}

type RestResponse struct {
	Message string `json:"message"`
}

func (rsa *RestServerAdapter) Register(handler *core.Handler) error {
	var err error = nil
	switch handler.Protocol {
	case "read":
		rsa.server.Get(handler.Name, func(c *fiber.Ctx) error {
			for _, task := range handler.TaskSets {
				_, err := task.ServiceAdapter.Call()
				if err != nil {
					return err
				}
			}
			return c.Status(fiber.StatusOK).JSON(RestResponse{Message: fmt.Sprintf("Successfully ran %s", handler.ID)})
		})
	case "write":
		rsa.server.Post(handler.Name, func(c *fiber.Ctx) error {
			for _, task := range handler.TaskSets {
				_, err := task.ServiceAdapter.Call()
				if err != nil {
					return err
				}
			}
			return c.Status(fiber.StatusOK).JSON(RestResponse{Message: fmt.Sprintf("Successfully ran %s", handler.ID)})
		})
	default:
		err = errors.New("Handler has no matching action")
	}
	return err
}

func StartServer(addr string) {
	app := setupRouter()

	err := app.Listen(addr)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}