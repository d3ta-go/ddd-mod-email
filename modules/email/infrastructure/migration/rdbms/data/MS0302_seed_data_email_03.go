package data

import (
	domEntity "github.com/d3ta-go/ddd-mod-email/modules/email/domain/entity"
	"github.com/d3ta-go/system/system/utils"
)

// EmailTemplate03 data (TEXT)
func EmailTemplate03() domEntity.EmailTemplateEntity {
	return domEntity.EmailTemplateEntity{
		UUID:        utils.GenerateUUID(),
		Code:        "account-activation-plaintext",
		Name:        "Account Activation Email (TEXT)",
		IsActive:    true,
		EmailFormat: "TEXT",
	}
}

// EmailTemplate03Version data
func EmailTemplate03Version() domEntity.EmailTemplateVersionEntity {
	return domEntity.EmailTemplateVersionEntity{
		Version:    utils.GenSemVersion(""),
		SubjectTpl: "Account Activation",
		BodyTpl: `{{define "T"}}Dear {{index . "Header.Name"}},

Conglatulation! your account has been activated.

If you didn't attempt to verify your email address with our service, delete this email.

Cheers,

{{index . "Footer.Name"}}
{{end}}`,
	}
}
