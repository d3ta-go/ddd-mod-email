package repository

import (
	"testing"

	"github.com/d3ta-go/ddd-mod-email/modules/email/domain/repository"
	domSchemaEmail "github.com/d3ta-go/ddd-mod-email/modules/email/domain/schema/email"
	domSchemaET "github.com/d3ta-go/ddd-mod-email/modules/email/domain/schema/email_template"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/spf13/viper"
)

func newEmailRepo(t *testing.T) (repository.IEmailRepo, *handler.Handler, error) {
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

	r, err := NewEmailRepo(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func TestEMailRepo_Send(t *testing.T) {
	repo, h, err := newEmailRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email.infra-layer.repo.send")
	testDataET := viper.GetStringMapString("test-data.email.email.app-layer.service.send.email-template-data")

	req := &domSchemaEmail.SendEmailRequest{
		TemplateCode: testData["email-template-code"],

		From: &domSchemaEmail.MailAddress{
			Email: testData["from-email"], Name: testData["from-name"],
		},
		To: &domSchemaEmail.MailAddress{
			Email: testData["to-email"], Name: testData["to-name"],
		},
		CC: []*domSchemaEmail.MailAddress{
			{Email: testData["cc-email-01"], Name: testData["cc-name-01"]},
			{Email: testData["cc-email-02"], Name: testData["cc-name-02"]}},
		BCC: []*domSchemaEmail.MailAddress{
			{Email: testData["bcc-email-01"], Name: testData["bcc-name-01"]},
			{Email: testData["bcc-email-02"], Name: testData["bcc-name-02"]}},
		TemplateData: map[string]interface{}{
			"Header.Name":        testDataET["header-name"],
			"Body.UserAccount":   testDataET["body-user-account"],
			"Body.ActivationURL": testDataET["body-activation-url"],
			"Footer.Name":        testDataET["footer-name"],
		},
		ProcessingType: testData["processing-type"],
		Template: &domSchemaET.ETFindByCodeData{
			EmailTemplate: domSchemaET.EmailTemplate{
				ID:          1,
				EmailFormat: "HTML",
			},
			DefaultTemplateVersion: domSchemaET.EmailTemplateVersion{
				SubjectTpl: testData["et-tpl-subject"],
				BodyTpl: `{{define "T"}}<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
	</head>
	<body>
		<p>
			Dear {{index . "Header.Name"}}, 
		</p>
		<p>
			Velit consequat pariatur nisi anim. Enim mollit mollit officia voluptate eiusmod tempor dolor proident aliqua officia occaecat dolor.
			<strong>Laboris consequat ex eu ad et tempor tempor ullamco sunt est. </strong>
			Aliquip sunt cillum aute exercitation nisi non. Ad quis adipisicing voluptate laboris est elit duis. 
			Nostrud fugiat enim qui qui incididunt irure. Et elit laborum aliquip ullamco enim veniam aliqua ad veniam do ipsum.
		</p>
		<p>
			<a href="{{index . "Body.URL"}}">Eiusmod fugiat sunt elit deserunt sint labore ut dolor sint enim.</a>
		</p>
		<p>
			Regards, 
		</p>
		<p>
			{{index . "Footer.Signature"}}
		</p>

	</body>
</html>
{{end}}`,
			},
		},
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Send(req, i)
	if err != nil {
		t.Errorf("Error.EmailRepo.Send: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailRepo.Send: %s", string(respJSON))
	}
}
