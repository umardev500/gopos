package router

import (
	"github.com/gofiber/fiber/v2"
	"gitub.com/umardev500/gopos/pkg/contract"
)

type Router interface {
	Setup()
}

type router struct {
	app        *fiber.App
	containers []contract.Container
}

func NewRouter(app *fiber.App, containers []contract.Container) Router {
	return &router{
		app:        app,
		containers: containers,
	}
}

func (r *router) Setup() {
	api := r.app.Group("/api")
	for _, container := range r.containers {
		container.HandleApi(api)
	}

	web := r.app.Group("/web")
	for _, container := range r.containers {
		container.HandleWeb(web)
	}
}
