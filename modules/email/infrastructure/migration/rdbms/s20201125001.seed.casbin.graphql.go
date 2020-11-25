package rdbms

import (
	"fmt"

	"github.com/d3ta-go/system/system/handler"
	migRDBMS "github.com/d3ta-go/system/system/migration/rdbms"
	"gorm.io/gorm"
)

// CasbinRule
var cPsGraphQL = []IamCasbinRule{
	// role:default - email
	{PType: "p", V0: "role:default", V1: "/api/v1/email/templates/list-all", V2: "GET"},
	{PType: "p", V0: "role:default", V1: "/api/v1/email/template/*", V2: "GET"},

	// role:default - email - grapqhql api
	{PType: "p", V0: "role:default", V1: "graphql.email.template.list-all", V2: "READ"},
	{PType: "p", V0: "role:default", V1: "graphql.email.template.find-by-code", V2: "READ"},

	// role:admin - email - grapqhql api
	{PType: "p", V0: "role:admin", V1: "graphql.email.send", V2: "EXECUTE"},
	{PType: "p", V0: "role:admin", V1: "graphql.email.template.list-all", V2: "READ"},
	{PType: "p", V0: "role:admin", V1: "graphql.email.template.find-by-code", V2: "READ"},
	{PType: "p", V0: "role:admin", V1: "graphql.email.template.create", V2: "CREATE"},
	{PType: "p", V0: "role:admin", V1: "graphql.email.template.update", V2: "UPDATE"},
	{PType: "p", V0: "role:admin", V1: "graphql.email.template.set-active", V2: "UPDATE"},
	{PType: "p", V0: "role:admin", V1: "graphql.email.template.delete", V2: "DELETE"},
}

var vGsGraphQL = []IamCasbinRule{
	// group -> role (for flexibility)
	{PType: "g", V0: "group:system", V1: "role:system"},
	{PType: "g", V0: "group:default", V1: "role:default"},
	{PType: "g", V0: "group:admin", V1: "role:admin"},
}

// Seed20201125001InitCasbinGraphQL type
type Seed20201125001InitCasbinGraphQL struct {
	migRDBMS.BaseGormMigratorRunner
}

// NewSeed20201125001InitCasbinGraphQL constructor
func NewSeed20201125001InitCasbinGraphQL(h *handler.Handler) (migRDBMS.IGormMigratorRunner, error) {
	gmr := new(Seed20201125001InitCasbinGraphQL)
	gmr.SetHandler(h)
	gmr.SetID("Seed20201125001InitCasbinGraphQL")
	return gmr, nil
}

// GetID get Seed20201125001InitCasbinGraphQL ID
func (dmr *Seed20201125001InitCasbinGraphQL) GetID() string {
	return fmt.Sprintf("%T", dmr)
}

// Run run Seed20201125001InitCasbinGraphQL
func (dmr *Seed20201125001InitCasbinGraphQL) Run(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr.GetGorm().AutoMigrate(&IamCasbinRule{}); err != nil {
			return err
		}
		if err := dmr._seeds(); err != nil {
			return err
		}
	}
	return nil
}

// RollBack rollback Seed20201125001InitCasbinGraphQL
func (dmr *Seed20201125001InitCasbinGraphQL) RollBack(h *handler.Handler, dbGorm *gorm.DB) error {
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

func (dmr *Seed20201125001InitCasbinGraphQL) _seeds() error {
	if dmr.GetGorm().Migrator().HasTable(&IamCasbinRule{}) {
		if err := dmr.GetGorm().Create(&cPsGraphQL).Error; err != nil {
			return err
		}

		for _, v := range vGsGraphQL {
			var ett IamCasbinRule
			result := dmr.GetGorm().Unscoped().Where(v).First(&ett)
			if result.RowsAffected < 1 {
				if err := dmr.GetGorm().Create(&v).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (dmr *Seed20201125001InitCasbinGraphQL) _unSeeds() error {
	if dmr.GetGorm().Migrator().HasTable(&IamCasbinRule{}) {

		for _, v := range cPsGraphQL {
			if err := dmr.GetGorm().Unscoped().Where(&v).Delete(&IamCasbinRule{}).Error; err != nil {
				return err
			}
		}

		for _, v := range vGsGraphQL {
			var ett IamCasbinRule
			result := dmr.GetGorm().Unscoped().Where(&IamCasbinRule{PType: "p", V0: v.V1}).First(&ett)
			if result.RowsAffected < 1 {
				if err := dmr.GetGorm().Where(&v).Delete(&IamCasbinRule{}).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}
