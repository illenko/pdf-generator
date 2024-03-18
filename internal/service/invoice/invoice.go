package invoice

import (
	"context"
	"github.com/illenko/pdf-generator/internal/logger"
	"github.com/illenko/pdf-generator/internal/model"
	"github.com/illenko/pdf-generator/internal/service/core"
	"log/slog"
)

type Service interface {
	Generate(ctx context.Context, request model.InvoiceRequest) (pdf []byte, err error)
}

type invoiceService struct {
	log             *slog.Logger
	templateService core.TemplateService
	pdfService      core.PdfService
}

func NewService(log *slog.Logger, templateService core.TemplateService, pdfService core.PdfService) Service {
	return invoiceService{log: log,
		templateService: templateService,
		pdfService:      pdfService,
	}
}

const TemplateFileName = "templates/invoice.html"

func (s invoiceService) Generate(ctx context.Context, request model.InvoiceRequest) (pdf []byte, err error) {

	ctx = logger.AppendCtx(ctx, slog.String("templateFileName", TemplateFileName))
	html, err := s.templateService.Process(ctx, TemplateFileName, request)

	if err != nil {
		s.log.ErrorContext(ctx, "Error while applying template")
		return nil, err
	}

	s.log.InfoContext(ctx, "Template applied")

	pdf, err = s.pdfService.Process(ctx, html)

	if err != nil {
		s.log.ErrorContext(ctx, "Error while converting html template to pdf")
		return nil, err
	}

	return
}
