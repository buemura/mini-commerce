package router

import (
	"github.com/buemura/event-driven-commerce/mc-api-gateway/internal/infra/http/controller"
	"github.com/labstack/echo/v4"
)

func setupOrderRouters(app *echo.Echo) {
	app.GET("/api/orders", controller.GetManyOrders)
	app.GET("/api/orders/:id", controller.GetOrder)
	app.POST("/api/orders", controller.CreateOrder)
}
