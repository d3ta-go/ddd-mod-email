package repository

import (
	domSchemaEmail "github.com/d3ta-go/ddd-mod-email/modules/email/domain/schema/email"
	"github.com/d3ta-go/system/system/identity"
)

// IEmailRepo represent EmailRepo interface
type IEmailRepo interface {
	Send(req *domSchemaEmail.SendEmailRequest, i identity.Identity) (*domSchemaEmail.SendEmailResponse, error)
}
