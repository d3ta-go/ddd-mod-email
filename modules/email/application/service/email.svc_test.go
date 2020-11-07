package service

import (
	"testing"

	appDTOEmail "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email"

	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
)

func newEmailService(t *testing.T) (*EmailService, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetDefaultConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, nil, err
	}

	s, err := NewEmailService(h)
	if err != nil {
		return nil, nil, err
	}

	return s, h, nil
}

func TestEmailService_Send(t *testing.T) {
	svc, h, err := newEmailService(t)
	if err != nil {
		t.Errorf("Error.newEmailService: %s", err.Error())
		return
	}

	req := new(appDTOEmail.SendEmailReqDTO)
	req.TemplateCode = "activate-registration-html"
	// req.TemplateCode = "account-activation-html"
	req.From = &appDTOEmail.MailAddressDTO{Email: "d3tago.from@domain.com", Name: "D3TA Golang"}
	req.To = &appDTOEmail.MailAddressDTO{Email: "d3tago.test@outlook.com", Name: "D3TAgo Test (Outlook)"}
	req.CC = []*appDTOEmail.MailAddressDTO{
		{Email: "d3tago.test@protonmail.com", Name: "D3TAgo Test CC 1 (Protonmail)"},
		{Email: "d3tago.test.cc@tutanota.com", Name: "D3TAgo Test CC 2 (Tutanota)"}}
	req.BCC = []*appDTOEmail.MailAddressDTO{
		{Email: "d3tago.test@tutanota.com", Name: "D3TAgo Test BCC 1 (Tutanota)"},
		{Email: "d3tago.test.bcc@outlook.com", Name: "D3TAgo Test BCC 2 (Outlook)"}}
	req.TemplateData = map[string]interface{}{
		"Header.Name":        "John Doe",
		"Body.UserAccount":   "john.doe",
		"Body.ActivationURL": "https://google.com",
		"Footer.Name":        "Customer Service",
	}
	req.ProcessingType = "SYNC"

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/email/send"
	i.RequestInfo.RequestAction = "POST"

	resp, err := svc.Send(req, i)
	if err != nil {
		t.Errorf("Error.EmailService.Send: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}
