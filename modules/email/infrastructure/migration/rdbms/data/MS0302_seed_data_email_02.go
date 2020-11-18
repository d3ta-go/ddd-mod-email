package data

import (
	domEntity "github.com/d3ta-go/ddd-mod-email/modules/email/domain/entity"
	"github.com/d3ta-go/system/system/utils"
)

// EmailTemplate02 data (HTML)
func EmailTemplate02() domEntity.EmailTemplateEntity {
	return domEntity.EmailTemplateEntity{
		UUID:        utils.GenerateUUID(),
		Code:        "activate-registration-html",
		Name:        "Activate Registration Email (HTML)",
		IsActive:    true,
		EmailFormat: "HTML",
	}
}

// EmailTemplate02Version data
func EmailTemplate02Version() domEntity.EmailTemplateVersionEntity {
	return domEntity.EmailTemplateVersionEntity{
		Version:    utils.GenSemVersion(""),
		SubjectTpl: "Activate Registration",
		BodyTpl: `{{define "T"}}<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
	</head>
	<body>
		<p>
			Dear {{index . "Header.Name"}},
		</p>
		<p>
			Please click on the url bellow to complete the verification process for account "<strong>{{index . "Body.UserAccount"}}</strong>":
		</p>
		<p>
			<a href="{{index . "Body.ActivationURL"}}">Activation URL</a>
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
