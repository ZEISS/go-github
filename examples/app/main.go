package main

import (
	"context"
	"log"
	"os"

	hooks "github.com/zeiss/fiber-hooks/v3"
	"github.com/zeiss/fiber-hooks/v3/github"
	"github.com/zeiss/go-github/apps"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/spf13/cobra"
	"github.com/zeiss/pkg/logx"
)

// Flags ...
type Flags struct {
	Addr   string
	Secret string
}

var (
	cfg   = apps.NewConfig()
	flags = &Flags{}
)

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, _ []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&flags.Addr, "addr", ":8080", "addr")
	rootCmd.PersistentFlags().StringVar(&flags.Secret, "secret", "", "secret")

	rootCmd.SilenceUsage = true
}

func run(_ context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	_, err := logx.RedirectStdLog(logx.LogSink)
	if err != nil {
		return err
	}

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New())

	dispatcher := apps.NewDispatcher(cfg)
	decoder := github.NewDecoder()

	hook := hooks.New(hooks.Config{
		SigningSecret: flags.Secret,
		Decoder:       decoder,
		Dispatcher:    dispatcher,
	})

	app.Post("/webhooks", hook)

	if err := app.Listen(flags.Addr); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
