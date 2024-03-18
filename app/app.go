package app

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/pdf-generator/internal/handler"
	"github.com/illenko/pdf-generator/internal/logger"
	"github.com/illenko/pdf-generator/internal/server"
	"github.com/illenko/pdf-generator/internal/service/core"
	"github.com/illenko/pdf-generator/internal/service/invoice"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

func Run() {
	fx.New(
		fx.Provide(
			logger.New,
			invoice.NewService,
			core.NewTemplateService,
			core.NewPdfService,
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
