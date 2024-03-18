package app

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/pdf-generator/internal/handler"
	"github.com/illenko/pdf-generator/internal/logger"
	"github.com/illenko/pdf-generator/internal/server"
	"github.com/illenko/pdf-generator/internal/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

type App struct{}

func (a App) Run() {
	fx.New(
		fx.Provide(
			logger.New,
			service.New,
			handler.NewPdf,
			handler.NewHealth,
			server.New,
		),
		fx.Invoke(func(e *gin.Engine) {
			err := e.Run(":8080")
			if err != nil {
				return
			}
		}),
		fx.WithLogger(func(log *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: log}
		}),
	).Run()
}
