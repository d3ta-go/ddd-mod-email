package data

import (
	domEntity "github.com/d3ta-go/ddd-mod-email/modules/email/domain/entity"
	"github.com/d3ta-go/system/system/utils"
)

// EmailTemplate01 data (TEXT)
func EmailTemplate01() domEntity.EmailTemplateEntity {
	return domEntity.EmailTemplateEntity{
		UUID:        utils.GenerateUUID(),
		Code:        "activate-registration-plaintext",
		Name:        "Activate Registration Email (PlainText)",
		IsActive:    true,
		EmailFormat: "TEXT",
	}
}

// EmailTemplate01Version data
func EmailTemplate01Version() domEntity.EmailTemplateVersionEntity {
	return domEntity.EmailTemplateVersionEntity{
		Version:    utils.GenSemVersion(""),
		SubjectTpl: "Activate Registration",
		BodyTpl: `{{define "T"}}Dear {{index . "Header.Name"}},

Please click on the url bellow to complete the verification process for account "{{index . "Body.UserAccount"}}":

{{index . "Body.ActivationURL"}}

If you didn't attempt to verify your email address with our service, delete this email.

Cheers,

{{index . "Footer.Name"}}
{{end}}`,
	}
}
