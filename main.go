package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/pes-alcoding-club/online_judge/util"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "oj-api-v1", log.LstdFlags)

	// TODO: Fine tune params.
	app := fiber.New(&fiber.Settings{
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	})
	app.Use(middleware.Recover())

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Fatal(app.Listen(8080))
		wg.Done()
	}()
	defer wg.Wait()

	// setting up routes.
	util.SetupRoutes(app)

	// attempting graceful shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate signal, attempting graceful shutdown...", sig)

	// make sure all current connections are closed and then shut down.
	// this WILL NOT terminate keep-alive connections. Fine tune read and
	// write timeouts in server settings.
	app.Shutdown()
}
