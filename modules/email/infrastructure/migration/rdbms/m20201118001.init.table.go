package rdbms

import (
	"fmt"

	domEntity "github.com/d3ta-go/ddd-mod-email/modules/email/domain/entity"
	"github.com/d3ta-go/system/system/handler"
	migRDBMS "github.com/d3ta-go/system/system/migration/rdbms"
	"gorm.io/gorm"
)

// Migrate20201118001InitTable type
type Migrate20201118001InitTable struct {
	migRDBMS.BaseGormMigratorRunner
}

// NewMigrate20201118001InitTable constructor
func NewMigrate20201118001InitTable(h *handler.Handler) (migRDBMS.IGormMigratorRunner, error) {
	gmr := new(Migrate20201118001InitTable)
	gmr.SetHandler(h)
	gmr.SetID("Migrate20201118001InitTable")
	return gmr, nil
}

// GetID get Migrate20201118001InitTable ID
func (dmr *Migrate20201118001InitTable) GetID() string {
	return fmt.Sprintf("%T", dmr)
}

// Run run Migrate20201118001InitTable
func (dmr *Migrate20201118001InitTable) Run(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr.GetGorm().AutoMigrate(
			&domEntity.EmailEntity{},
			&domEntity.EmailTemplateEntity{},
			&domEntity.EmailTemplateVersionEntity{},
		); err != nil {
			return err
		}
	}
	return nil
}

// RollBack rollback Migrate20201118001InitTable
func (dmr *Migrate20201118001InitTable) RollBack(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if dmr.GetGorm().Migrator().HasTable(&domEntity.EmailEntity{}) {
			if err := dmr.GetGorm().Migrator().DropTable(&domEntity.EmailEntity{}); err != nil {
				return err
			}
		}
		if dmr.GetGorm().Migrator().HasTable(&domEntity.EmailTemplateVersionEntity{}) {
			if err := dmr.GetGorm().Migrator().DropTable(&domEntity.EmailTemplateVersionEntity{}); err != nil {
				return err
			}
		}
		if dmr.GetGorm().Migrator().HasTable(&domEntity.EmailTemplateEntity{}) {
			if err := dmr.GetGorm().Migrator().DropTable(&domEntity.EmailTemplateEntity{}); err != nil {
				return err
			}
		}
	}
	return nil
}
