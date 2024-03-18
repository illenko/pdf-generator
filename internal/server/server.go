package server

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/pdf-generator/docs"
	"github.com/illenko/pdf-generator/internal/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(handler handler.InvoiceHandler, healthHandler handler.HealthHandler) *gin.Engine {
	e := gin.Default()

	e.POST("/invoices", handler.GenerateInvoice)
	e.GET("/health/liveness", healthHandler.Liveness)
	e.GET("/health/readiness", healthHandler.Readiness)

	docs.SwaggerInfo.BasePath = "/"
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return e
}
