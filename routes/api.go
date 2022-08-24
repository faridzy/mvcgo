/*
 * MVCGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package routes

import (
	"mvcgo/app/http/controller"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	//Custom Middleware & Middleware
	route := app.Group("/api") // /api
	// //Initila Route for Controller
	ProductController := controller.ProductController{}
	ProductController.Route(route)

}
