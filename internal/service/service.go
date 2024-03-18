package service

import (
	"context"
	"github.com/illenko/pdf-generator/internal/model"
	"log/slog"
)

type InvoiceService interface {
	Generate(ctx context.Context, request model.InvoiceRequest) string
}

type invoiceService struct {
	log *slog.Logger
}

func New(log *slog.Logger) InvoiceService {
	return invoiceService{log: log}
}

func (s invoiceService) Generate(ctx context.Context, request model.InvoiceRequest) string {
	return ""
}
