package service

import (
	"testing"

	appDTOEmail "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email"
	"github.com/spf13/viper"

	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
)

func newEmailService(t *testing.T) (*EmailService, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, v, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetDefaultConfig(c)
	h.SetViper("config", v)

	// viper for test-data
	viperTest := viper.New()
	viperTest.SetConfigType("yaml")
	viperTest.SetConfigName("test-data")
	viperTest.AddConfigPath("../../../../conf/data")
	viperTest.ReadInConfig()
	h.SetViper("test-data", viperTest)

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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email.app-layer.service.send")

	req := new(appDTOEmail.SendEmailReqDTO)
	req.TemplateCode = testData["email-template-code"]
	// req.TemplateCode = "account-activation-html"
	req.From = &appDTOEmail.MailAddressDTO{
		Email: testData["from-email"], Name: testData["from-name"],
	}
	req.To = &appDTOEmail.MailAddressDTO{
		Email: testData["to-email"], Name: testData["to-name"],
	}
	req.CC = []*appDTOEmail.MailAddressDTO{
		{Email: testData["cc-email-01"], Name: testData["cc-name-01"]},
		{Email: testData["cc-email-02"], Name: testData["cc-name-02"]}}
	req.BCC = []*appDTOEmail.MailAddressDTO{
		{Email: testData["bcc-email-01"], Name: testData["bcc-name-01"]},
		{Email: testData["bcc-email-02"], Name: testData["bcc-name-02"]}}

	testDataET := viper.GetStringMapString("test-data.email.email.app-layer.service.send.email-template-data")
	req.TemplateData = map[string]interface{}{
		"Header.Name":        testDataET["header-name"],
		"Body.UserAccount":   testDataET["body-user-account"],
		"Body.ActivationURL": testDataET["body-activation-url"],
		"Footer.Name":        testDataET["footer-name"],
	}
	req.ProcessingType = testData["processing-type"]

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
