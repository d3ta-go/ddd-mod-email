package application

import (
	appSvc "github.com/d3ta-go/ddd-mod-email/modules/email/application/service"
	"github.com/d3ta-go/system/system/handler"
)

// NewEmailApp new EmailApp
func NewEmailApp(h *handler.Handler) (*EmailApp, error) {
	var err error

	app := new(EmailApp)
	app.handler = h

	if app.EmailSvc, err = appSvc.NewEmailService(h); err != nil {
		return nil, err
	}
	if app.EmailTemplateSvc, err = appSvc.NewEmailTemplateService(h); err != nil {
		return nil, err
	}

	return app, nil
}

// EmailApp represent DDD Module: Email (Application Layer)
type EmailApp struct {
	handler          *handler.Handler
	EmailSvc         *appSvc.EmailService
	EmailTemplateSvc *appSvc.EmailTemplateService
}
