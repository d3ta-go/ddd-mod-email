package email

import (
	"encoding/json"

	domSchemaEmail "github.com/d3ta-go/ddd-mod-email/modules/email/domain/schema/email"
)

// SendEmailReqDTO type
type SendEmailReqDTO struct {
	TemplateCode   string                 `json:"templateCode"`
	From           *MailAddressDTO        `json:"from"`
	To             *MailAddressDTO        `json:"to"`
	CC             []*MailAddressDTO      `json:"cc"`
	BCC            []*MailAddressDTO      `json:"bcc"`
	TemplateData   map[string]interface{} `json:"templateData"`
	ProcessingType string                 `json:"processingType"`
	// domSchemaEmail.SendEmailRequest
}

// ConvertCC2Domain convert to domSchema
func (r *SendEmailReqDTO) ConvertCC2Domain() []*domSchemaEmail.MailAddress {
	var ms []*domSchemaEmail.MailAddress
	for _, v := range r.CC {
		ms = append(ms, &domSchemaEmail.MailAddress{Email: v.Email, Name: v.Name})
	}
	return ms
}

// ConvertBCC2Domain convert to domSchema
func (r *SendEmailReqDTO) ConvertBCC2Domain() []*domSchemaEmail.MailAddress {
	var ms []*domSchemaEmail.MailAddress
	for _, v := range r.BCC {
		ms = append(ms, &domSchemaEmail.MailAddress{Email: v.Email, Name: v.Name})
	}
	return ms
}

// SendEmailResDTO type
type SendEmailResDTO struct {
	domSchemaEmail.SendEmailResponse
}

// MailAddressDTO type
type MailAddressDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ToJSON covert to JSON
func (r *SendEmailReqDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
