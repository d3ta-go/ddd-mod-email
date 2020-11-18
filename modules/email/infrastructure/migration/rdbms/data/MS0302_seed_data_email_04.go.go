package data

import (
	domEntity "github.com/d3ta-go/ddd-mod-email/modules/email/domain/entity"
	"github.com/d3ta-go/system/system/utils"
)

// EmailTemplate04 data (TEXT)
func EmailTemplate04() domEntity.EmailTemplateEntity {
	return domEntity.EmailTemplateEntity{
		UUID:        utils.GenerateUUID(),
		Code:        "account-activation-html",
		Name:        "Account Activation Email (HTML)",
		IsActive:    true,
		EmailFormat: "HTML",
	}
}

// EmailTemplate04Version data
func EmailTemplate04Version() domEntity.EmailTemplateVersionEntity {
	return domEntity.EmailTemplateVersionEntity{
		Version:    utils.GenSemVersion(""),
		SubjectTpl: "Account Activation",
		BodyTpl: `{{define "T"}}<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
	</head>
	<body>
		<p>
			Dear {{index . "Header.Name"}},
		</p>
		<p>
			Conglatulation! your account has been activated.
		</p>
		<p>
			If you didn't attempt to verify your email address with our service, delete this email.
		</p>
		<p>
			Cheers,
		</p>
		<p>
			{{index . "Footer.Name"}}
		</p>
	</body>
</html>{{end}}`,
	}
}
