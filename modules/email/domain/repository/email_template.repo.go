package repository

import (
	domSchemaET "github.com/d3ta-go/ddd-mod-email/modules/email/domain/schema/email_template"
	"github.com/d3ta-go/system/system/identity"
)

// IEmailTemplateRepo interface
type IEmailTemplateRepo interface {
	ListAll(i identity.Identity) (*domSchemaET.ETListAllResponse, error)
	FindByCode(req *domSchemaET.ETFindByCodeRequest, i identity.Identity) (*domSchemaET.ETFindByCodeResponse, error)
	Create(req *domSchemaET.ETCreateRequest, i identity.Identity) (*domSchemaET.ETCreateResponse, error)
	Update(req *domSchemaET.ETUpdateRequest, i identity.Identity) (*domSchemaET.ETUpdateResponse, error)
	SetActive(req *domSchemaET.ETSetActiveRequest, i identity.Identity) (*domSchemaET.ETSetActiveResponse, error)
	Delete(req *domSchemaET.ETDeleteRequest, i identity.Identity) (*domSchemaET.ETDeleteResponse, error)
}
