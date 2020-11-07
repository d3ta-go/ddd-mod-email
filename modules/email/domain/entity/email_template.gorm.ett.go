package entity

// EmailTemplateEntity represent EmailTemplate Entity
type EmailTemplateEntity struct {
	ID uint64 `json:"ID" gorm:"primary_key;column:id"`

	UUID             string `json:"uuid" gorm:"column:uuid;size:255;unique;not null"`
	Code             string `json:"code" gorm:"column:code;size:100;unique;not null"`
	Name             string `json:"name" gorm:"column:name;size:255;not null"`
	IsActive         bool   `json:"isActive" gorm:"column:is_active;index"`
	EmailFormat      string `json:"emailFormat" gorm:"column:email_format;size:10;not null"`
	DefaultVersionID uint64 `json:"defaultVersionID" gorm:"column:default_version_id;not null"`

	BaseEntity
}

// TableName get real database table name
func (t *EmailTemplateEntity) TableName() string {
	return "eml_email_templates"
}
