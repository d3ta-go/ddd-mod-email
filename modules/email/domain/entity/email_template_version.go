package entity

// EmailTemplateVersionEntity represent EmailTemplateVersion Entity
type EmailTemplateVersionEntity struct {
	ID uint64 `json:"ID" gorm:"primary_key;column:id"`

	Version    string `json:"version" gorm:"column:version;size:100;index:idx_version,unique;not null"`
	SubjectTpl string `json:"subjectTpl" gorm:"column:subject_tpl;size:255;not null"`
	BodyTpl    string `json:"bodyTpl" gorm:"column:body_tpl"`

	EmailTemplateID uint64              `json:"emailTemplateID" gorm:"column:email_template_id;index:idx_version,unique;not null"`
	EmailTemplate   EmailTemplateEntity `json:"emailTemplate" gorm:"ForeignKey:EmailTemplateID;AssociationForeignKey:ID;"`

	BaseEntity
}

// TableName get real database table name
func (t *EmailTemplateVersionEntity) TableName() string {
	return "eml_email_template_versions"
}
