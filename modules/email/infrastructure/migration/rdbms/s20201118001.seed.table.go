package rdbms

import (
	"fmt"

	domEntity "github.com/d3ta-go/ddd-mod-email/modules/email/domain/entity"
	"github.com/d3ta-go/ddd-mod-email/modules/email/infrastructure/migration/rdbms/data"
	"github.com/d3ta-go/system/system/handler"
	migRDBMS "github.com/d3ta-go/system/system/migration/rdbms"
	"gorm.io/gorm"
)

// Seed20201118001InitTable type
type Seed20201118001InitTable struct {
	migRDBMS.BaseGormMigratorRunner
}

// NewSeed20201118001InitTable constructor
func NewSeed20201118001InitTable(h *handler.Handler) (migRDBMS.IGormMigratorRunner, error) {
	gmr := new(Seed20201118001InitTable)
	gmr.SetHandler(h)
	gmr.SetID("Seed20201118001InitTable")
	return gmr, nil
}

// GetID get Seed20201118001InitTable ID
func (dmr *Seed20201118001InitTable) GetID() string {
	return fmt.Sprintf("%T", dmr)
}

// Run run Seed20201118001InitTable
func (dmr *Seed20201118001InitTable) Run(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr._seeds(); err != nil {
			return err
		}
	}
	return nil
}

// RollBack rollback Seed20201118001InitTable
func (dmr *Seed20201118001InitTable) RollBack(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr._unSeeds(); err != nil {
			return err
		}
	}
	return nil
}

func (dmr *Seed20201118001InitTable) _seeds() error {
	if dmr.GetGorm().Migrator().HasTable(&domEntity.EmailTemplateEntity{}) &&
		dmr.GetGorm().Migrator().HasTable(&domEntity.EmailTemplateVersionEntity{}) {

		// data01
		eTpl := data.EmailTemplate01()
		eTpl.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eTpl).Error; err != nil {
			return err
		}

		eTplVer := data.EmailTemplate01Version()
		eTplVer.EmailTemplateID = eTpl.ID
		eTplVer.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eTplVer).Error; err != nil {
			return err
		}

		eTpl.DefaultVersionID = eTplVer.ID
		if err := dmr.GetGorm().Save(&eTpl).Error; err != nil {
			return err
		}

		// data02
		eTpl = data.EmailTemplate02()
		eTpl.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eTpl).Error; err != nil {
			return err
		}

		eTplVer = data.EmailTemplate02Version()
		eTplVer.EmailTemplateID = eTpl.ID
		eTplVer.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eTplVer).Error; err != nil {
			return err
		}

		eTpl.DefaultVersionID = eTplVer.ID
		if err := dmr.GetGorm().Save(&eTpl).Error; err != nil {
			return err
		}

		// data03
		eTpl = data.EmailTemplate03()
		eTpl.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eTpl).Error; err != nil {
			return err
		}

		eTplVer = data.EmailTemplate03Version()
		eTplVer.EmailTemplateID = eTpl.ID
		eTplVer.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eTplVer).Error; err != nil {
			return err
		}

		eTpl.DefaultVersionID = eTplVer.ID
		if err := dmr.GetGorm().Save(&eTpl).Error; err != nil {
			return err
		}

		// data04
		eTpl = data.EmailTemplate04()
		eTpl.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eTpl).Error; err != nil {
			return err
		}

		eTplVer = data.EmailTemplate04Version()
		eTplVer.EmailTemplateID = eTpl.ID
		eTplVer.CreatedBy = "system.d3tago@installation"
		if err := dmr.GetGorm().Create(&eTplVer).Error; err != nil {
			return err
		}

		eTpl.DefaultVersionID = eTplVer.ID
		if err := dmr.GetGorm().Save(&eTpl).Error; err != nil {
			return err
		}
		//
	}
	return nil
}

func (dmr *Seed20201118001InitTable) _unSeeds() error {
	if dmr.GetGorm().Migrator().HasTable(&domEntity.EmailTemplateEntity{}) &&
		dmr.GetGorm().Migrator().HasTable(&domEntity.EmailTemplateVersionEntity{}) {

		// data01
		eTpl := data.EmailTemplate01()
		var ett01 domEntity.EmailTemplateEntity
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateEntity{Code: eTpl.Code}).First(&ett01).Error; err != nil {
			return err
		}

		eTplVer := data.EmailTemplate01Version()
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateVersionEntity{EmailTemplateID: ett01.ID}).Delete(&eTplVer).Error; err != nil {
			return err
		}
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateEntity{Code: eTpl.Code}).Delete(&eTpl).Error; err != nil {
			return err
		}

		// data02
		eTpl = data.EmailTemplate02()
		var ett02 domEntity.EmailTemplateEntity
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateEntity{Code: eTpl.Code}).First(&ett02).Error; err != nil {
			return err
		}

		eTplVer = data.EmailTemplate02Version()
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateVersionEntity{EmailTemplateID: ett02.ID}).Delete(&eTplVer).Error; err != nil {
			return err
		}
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateEntity{Code: eTpl.Code}).Delete(&eTpl).Error; err != nil {
			return err
		}

		// data03
		eTpl = data.EmailTemplate03()
		var ett03 domEntity.EmailTemplateEntity
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateEntity{Code: eTpl.Code}).First(&ett03).Error; err != nil {
			return err
		}

		eTplVer = data.EmailTemplate03Version()
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateVersionEntity{EmailTemplateID: ett03.ID}).Delete(&eTplVer).Error; err != nil {
			return err
		}
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateEntity{Code: eTpl.Code}).Delete(&eTpl).Error; err != nil {
			return err
		}

		// data04
		eTpl = data.EmailTemplate04()
		var ett04 domEntity.EmailTemplateEntity
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateEntity{Code: eTpl.Code}).First(&ett04).Error; err != nil {
			return err
		}

		eTplVer = data.EmailTemplate04Version()
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateVersionEntity{EmailTemplateID: ett04.ID}).Delete(&eTplVer).Error; err != nil {
			return err
		}
		if err := dmr.GetGorm().Unscoped().Where(&domEntity.EmailTemplateEntity{Code: eTpl.Code}).Delete(&eTpl).Error; err != nil {
			return err
		}

	}
	return nil
}
