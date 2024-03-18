package core

import (
	"bytes"
	"context"
	"html/template"
	"log/slog"
)

type TemplateService interface {
	Process(ctx context.Context, templateFileName string, data interface{}) (html string, err error)
}

type templateService struct {
	log *slog.Logger
}

func NewTemplateService(log *slog.Logger) TemplateService {
	return templateService{log: log}
}

func (s templateService) Process(ctx context.Context, templateFileName string, data interface{}) (html string, err error) {

	s.log.InfoContext(ctx, "Parsing html template")

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		s.log.ErrorContext(ctx, "Error while parsing template")
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		s.log.ErrorContext(ctx, "Error while applying data to template")
		return "", err

	}

	return buf.String(), nil
}
