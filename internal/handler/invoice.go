package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/illenko/pdf-generator/internal/logger"
	"github.com/illenko/pdf-generator/internal/model"
	"github.com/illenko/pdf-generator/internal/service/invoice"
	"log/slog"
	"net/http"
)

type InvoiceHandler interface {
	GenerateInvoice(c *gin.Context)
}

type invoiceHandler struct {
	log     *slog.Logger
	service invoice.Service
}

func NewPdf(log *slog.Logger, service invoice.Service) InvoiceHandler {
	return invoiceHandler{
		log:     log,
		service: service,
	}
}

type ResponseError struct {
	ID    uuid.UUID `json:"id"`
	Error string    `json:"error"`
}

// GenerateInvoice
//
//	@Summary	Generates invoice pdf file
//	@Tags			invoice
//	@Accept			json
//	@Produce		octet-stream
//
// @Param data body model.InvoiceRequest true "Request"
//
//	@Success		200
//	@Failure		400 {object} handler.ResponseError
//	@Failure		500 {object} handler.ResponseError
//	@Router			/invoices [post]
func (h invoiceHandler) GenerateInvoice(c *gin.Context) {
	requestID := uuid.New()
	ctx := logger.AppendCtx(context.Background(), slog.String("requestID", requestID.String()))

	h.log.InfoContext(ctx, "Processing generation request")

	var req model.InvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.WarnContext(ctx, "Invalid request body")
		c.JSON(http.StatusBadRequest, ResponseError{ID: requestID, Error: "Invalid Request Body"})
		return
	}

	pdf, err := h.service.Generate(ctx, req)

	if err != nil {
		h.log.ErrorContext(ctx, "Internal server error")
		c.JSON(http.StatusInternalServerError, ResponseError{ID: requestID, Error: "Internal Server Error"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=file.pdf")
	c.Data(http.StatusOK, "application/pdf", pdf)
}
