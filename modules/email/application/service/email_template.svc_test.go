package service

import (
	"fmt"
	"strconv"
	"testing"

	appDTO "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
	domSchema "github.com/d3ta-go/ddd-mod-email/modules/email/domain/schema/email_template"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/d3ta-go/system/system/utils"
	"github.com/spf13/viper"
)

func newEmailTemplateService(t *testing.T) (*EmailTemplateService, *handler.Handler, error) {
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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.app-layer.service.find-by-code")

	req := new(appDTO.ETFindByCodeReqDTO)
	req.Code = testData["et-code"]

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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.app-layer.service.create")

	unique := utils.GenerateUUID()
	req := new(appDTO.ETCreateReqDTO)
	req.Code = fmt.Sprintf(testData["et-code"], unique)
	req.Name = fmt.Sprintf(testData["et-name"], unique)
	// req.IsActive = true/false
	req.IsActive, _ = strconv.ParseBool(testData["et-is-active"])
	req.EmailFormat = testData["et-email-format"]
	req.Template = &domSchema.ETCreateVersion{
		SubjectTpl: testData["et-tpl-subject"],
		BodyTpl:    testData["et-tpl-body"],
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
		// save to test-data
		// save result for next test
		viper.Set("test-data.email.email-template.app-layer.service.update.et-code", req.Code)
		viper.Set("test-data.email.email-template.app-layer.service.update.et-name", req.Name)

		viper.Set("test-data.email.email-template.app-layer.service.set-active.et-code", req.Code)
		viper.Set("test-data.email.email-template.app-layer.service.delete.et-code", req.Code)
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestEmailTemplateService_Update(t *testing.T) {
	svc, h, err := newEmailTemplateService(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateService: %s", err.Error())
		return
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.app-layer.service.update")

	req := new(appDTO.ETUpdateReqDTO)
	req.Keys = new(appDTO.ETUpdateKeysDTO)
	req.Keys.Code = testData["et-code"]

	req.Data = new(appDTO.ETUpdateDataDTO)
	req.Data.Name = testData["et-name"] + " Updated"
	// req.Data.IsActive = true/false
	req.Data.IsActive, _ = strconv.ParseBool(testData["et-is-active"])
	req.Data.EmailFormat = testData["et-email-format"]
	req.Data.Template = &domSchema.ETUpdateVersion{
		SubjectTpl: testData["et-tpl-subject"],
		BodyTpl:    testData["et-tpl-body"],
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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.app-layer.service.set-active")

	req := new(appDTO.ETSetActiveReqDTO)
	req.Keys = new(appDTO.ETSetActiveKeysDTO)
	req.Keys.Code = testData["et-code"]

	req.Data = new(appDTO.ETSetActiveDataDTO)
	// req.Data.IsActive = true/false
	req.Data.IsActive, _ = strconv.ParseBool(testData["et-is-active"])

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

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.app-layer.service.delete")

	req := new(appDTO.ETDeleteReqDTO)
	req.Code = testData["et-code"]

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
