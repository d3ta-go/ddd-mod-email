package repository

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/d3ta-go/ddd-mod-email/modules/email/domain/repository"
	schema "github.com/d3ta-go/ddd-mod-email/modules/email/domain/schema/email_template"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/d3ta-go/system/system/utils"
	"github.com/spf13/viper"
)

func newEmailTemplateRepo(t *testing.T) (repository.IEmailTemplateRepo, *handler.Handler, error) {
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

	r, err := NewEmailTemplateRepo(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func TestEMailTemplateRepo_ListAll(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	i := newIdentity(h, t)

	res, err := repo.ListAll(i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.ListAll: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.ListAll: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_FindByCode(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.infra-layer.repo.find-by-code")

	req := &schema.ETFindByCodeRequest{
		Code: testData["et-code"],
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.FindByCode(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.FindByCode: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.FindByCode: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_Create(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.infra-layer.repo.create")

	unique := utils.GenerateUUID()
	isActive, _ := strconv.ParseBool(testData["et-is-active"])

	req := &schema.ETCreateRequest{
		Code:        fmt.Sprintf(testData["et-code"], unique),
		Name:        fmt.Sprintf(testData["et-name"], unique),
		IsActive:    isActive,
		EmailFormat: testData["et-email-format"],
		Template: &schema.ETCreateVersion{
			SubjectTpl: testData["et-tpl-subject"],
			BodyTpl:    testData["et-tpl-body"],
		},
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Create(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.Create: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		// save to test-data
		// save result for next test
		viper.Set("test-data.email.email-template.infra-layer.repo.update.et-code", req.Code)
		viper.Set("test-data.email.email-template.infra-layer.repo.update.et-name", req.Name)

		viper.Set("test-data.email.email-template.infra-layer.repo.set-active.et-code", req.Code)
		viper.Set("test-data.email.email-template.infra-layer.repo.delete.et-code", req.Code)
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("Resp.EmailTemplateRepo.Create: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_Update(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.infra-layer.repo.update")

	isActive, _ := strconv.ParseBool(testData["et-is-active"])
	req := &schema.ETUpdateRequest{
		Keys: &schema.ETUpdateKeys{
			Code: testData["et-code"],
		},
		Data: &schema.ETUpdateData{
			Name:        testData["et-name"] + " Updated",
			IsActive:    isActive,
			EmailFormat: testData["et-email-format"],
			Template: &schema.ETUpdateVersion{
				SubjectTpl: testData["et-tpl-subject"],
				BodyTpl:    testData["et-tpl-body"],
			},
		},
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Update(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.Update: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.Update: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_SetActive(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.infra-layer.repo.set-active")

	isActive, _ := strconv.ParseBool(testData["et-is-active"])
	req := &schema.ETSetActiveRequest{
		Keys: &schema.ETSetActiveKeys{
			Code: testData["et-code"],
		},
		Data: &schema.ETSetActiveData{
			IsActive: isActive,
		},
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.SetActive(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.SetActive: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.SetActive: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_Delete(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.infra-layer.repo.delete")

	req := &schema.ETDeleteRequest{
		Code: testData["et-code"],
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Delete(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.Delete: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.Delete: %s", string(respJSON))
	}
}
