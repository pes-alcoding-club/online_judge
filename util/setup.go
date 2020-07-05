package util

import (
	"github.com/gofiber/fiber"
	"github.com/pes-alcoding-club/online_judge/routes"
)

// SetupRoutes creates routes for provided path strings.
func SetupRoutes(app *fiber.App) {
	app.Post("/submission", routes.PostSubmission)
}
