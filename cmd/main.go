package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitub.com/umardev500/gopos/internal/app/container"
	"gitub.com/umardev500/gopos/pkg/database"
	"gitub.com/umardev500/gopos/pkg/router"
	"gitub.com/umardev500/gopos/pkg/validator"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	ch := make(chan error, 1)

	db := database.GetGormInstance()
	v := validator.NewValidator()
	containers := container.NewRegistryContainer(db, v)
	router.NewRouter(app, containers).Setup()

	go func() {
		port := os.Getenv("PORT")
		addr := ":" + port

		log.Info().Msgf("Server running on %s", addr)
		ch <- app.Listen(addr)
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	select {
	case err := <-ch:
		log.Fatal().Err(err).Msg("Failed to start server")

	case <-ctx.Done():
		fmt.Println()
		app.Shutdown()
		log.Fatal().Msg("Graceful shutdown")
	}
}
