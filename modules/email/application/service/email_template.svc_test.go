package service

import (
	"testing"

	appDTO "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
	domSchema "github.com/d3ta-go/ddd-mod-email/modules/email/domain/schema/email_template"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
)

func newEmailTemplateService(t *testing.T) (*EmailTemplateService, *handler.Handler, error) {
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

	s, err := NewEmailTemplateService(h)
	if err != nil {
		return nil, nil, err
	}
	return s, h, nil
}

func TestEmailTemplateService_ListAll(t *testing.T) {
	svc, h, err := newEmailTemplateService(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateService: %s", err.Error())
		return
	}

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/email/templates/list-all"
	i.RequestInfo.RequestAction = "GET"

	resp, err := svc.ListAll(i)
	if err != nil {
		t.Errorf("Error.EmailTemplateService.ListAll: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestEmailTemplateService_FindByCode(t *testing.T) {
	svc, h, err := newEmailTemplateService(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateService: %s", err.Error())
		return
	}

	req := new(appDTO.ETFindByCodeReqDTO)
	req.Code = "test.code"

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/email/template/*"
	i.RequestInfo.RequestAction = "GET"

	resp, err := svc.FindByCode(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateService.FindByCode: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestEmailTemplateService_Create(t *testing.T) {
	svc, h, err := newEmailTemplateService(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateService: %s", err.Error())
		return
	}

	req := new(appDTO.ETCreateReqDTO)
	req.Code = "test.code.03"
	req.Name = "Template Name 03"
	req.IsActive = true
	req.EmailFormat = "TEXT"
	req.Template = &domSchema.ETCreateVersion{
		SubjectTpl: "Subject Template",
		BodyTpl:    `{{define "T"}}Body Template{{end}}`,
	}

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/email/template"
	i.RequestInfo.RequestAction = "POST"

	resp, err := svc.Create(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateService.Create: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestEmailTemplateService_Update(t *testing.T) {
	svc, h, err := newEmailTemplateService(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateService: %s", err.Error())
		return
	}

	req := new(appDTO.ETUpdateReqDTO)
	req.Keys = new(appDTO.ETUpdateKeysDTO)
	req.Keys.Code = "test.code.03"

	req.Data = new(appDTO.ETUpdateDataDTO)
	req.Data.Name = "Template Name 03 Updated"
	req.Data.IsActive = true
	req.Data.EmailFormat = "TEXT"
	req.Data.Template = &domSchema.ETUpdateVersion{
		SubjectTpl: "Subject Template Updated",
		BodyTpl:    `{{define "T"}}Body Template Updated{{end}}`,
	}

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/email/template/update/*"
	i.RequestInfo.RequestAction = "PUT"

	resp, err := svc.Update(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateService.Update: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestEmailTemplateService_SetActive(t *testing.T) {
	svc, h, err := newEmailTemplateService(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateService: %s", err.Error())
		return
	}

	req := new(appDTO.ETSetActiveReqDTO)
	req.Keys = new(appDTO.ETSetActiveKeysDTO)
	req.Keys.Code = "test.code.03"

	req.Data = new(appDTO.ETSetActiveDataDTO)
	req.Data.IsActive = true

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/email/template/set-active/*"
	i.RequestInfo.RequestAction = "PUT"

	resp, err := svc.SetActive(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateService.SetActive: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestEmailTemplateService_Delete(t *testing.T) {
	svc, h, err := newEmailTemplateService(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateService: %s", err.Error())
		return
	}

	req := new(appDTO.ETDeleteReqDTO)
	req.Code = "test.code.03"

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/email/template/*"
	i.RequestInfo.RequestAction = "DELETE"

	resp, err := svc.Delete(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateService.Delete: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}
