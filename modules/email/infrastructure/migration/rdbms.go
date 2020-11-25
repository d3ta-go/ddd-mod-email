package migration

import (
	migRunner "github.com/d3ta-go/ddd-mod-email/modules/email/infrastructure/migration/rdbms"
	"github.com/d3ta-go/system/system/handler"
	migRDBMS "github.com/d3ta-go/system/system/migration/rdbms"
)

// NewRDBMSMigration create new RDBMSMigration
func NewRDBMSMigration(h *handler.Handler) (*RDBMSMigration, error) {
	var err error

	mig := new(RDBMSMigration)
	mig.handler = h

	cfg, err := h.GetDefaultConfig()
	if err != nil {
		return nil, err
	}
	if mig.migrator, err = migRDBMS.NewBaseGormMigrator(h, cfg.Databases.EmailDB.ConnectionName); err != nil {
		return nil, err
	}

	return mig, nil
}

// RDBMSMigration represent RDBMSMigration
type RDBMSMigration struct {
	handler  *handler.Handler
	migrator *migRDBMS.BaseGormMigrator
}

// Run run migration
func (m *RDBMSMigration) Run() error {
	if err := m._runMigrates(); err != nil {
		return err
	}
	if err := m._runSeeds(); err != nil {
		return err
	}
	return nil
}

// RollBack rollback migration
func (m *RDBMSMigration) RollBack() error {
	if err := m._rollBackSeeds(); err != nil {
		return err
	}

	if err := m._rollBackMigrates(); err != nil {
		return err
	}
	return nil
}

func (m *RDBMSMigration) _runMigrates() error {
	migrate20201118001InitTable, err := migRunner.NewMigrate20201118001InitTable(m.handler)
	if err != nil {
		return err
	}

	cfg, err := m.handler.GetDefaultConfig()
	if err != nil {
		return err
	}
	if err := m.migrator.RunMigrates(m.handler, cfg.Databases.EmailDB.ConnectionName,
		migrate20201118001InitTable,
	); err != nil {
		return err
	}
	return nil
}

func (m *RDBMSMigration) _rollBackMigrates() error {
	migrate20201118001InitTable, err := migRunner.NewMigrate20201118001InitTable(m.handler)
	if err != nil {
		return err
	}

	cfg, err := m.handler.GetDefaultConfig()
	if err != nil {
		return err
	}
	if err := m.migrator.RollBackMigrates(m.handler, cfg.Databases.EmailDB.ConnectionName,
		migrate20201118001InitTable,
	); err != nil {
		return err
	}
	return nil
}

func (m *RDBMSMigration) _runSeeds() error {
	seed20201118001InitTable, err := migRunner.NewSeed20201118001InitTable(m.handler)
	if err != nil {
		return err
	}

	cfg, err := m.handler.GetDefaultConfig()
	if err != nil {
		return err
	}
	if err := m.migrator.RunSeeds(m.handler, cfg.Databases.EmailDB.ConnectionName,
		seed20201118001InitTable,
	); err != nil {
		return err
	}

	// iam
	if err := m._runIdentitySeeds(); err != nil {
		return err
	}
	return nil
}

func (m *RDBMSMigration) _rollBackSeeds() error {
	seed20201118001InitTable, err := migRunner.NewSeed20201118001InitTable(m.handler)
	if err != nil {
		return err
	}

	cfg, err := m.handler.GetDefaultConfig()
	if err != nil {
		return err
	}
	if err := m.migrator.RollBackSeeds(m.handler, cfg.Databases.EmailDB.ConnectionName,
		seed20201118001InitTable,
	); err != nil {
		return err
	}

	// iam
	if err := m._rollBackIdentitySeeds(); err != nil {
		return err
	}
	return nil
}

func (m *RDBMSMigration) _runIdentitySeeds() error {
	seed20201118002InitCasbin, err := migRunner.NewSeed20201118002InitCasbin(m.handler)
	if err != nil {
		return err
	}
	seed20201125001InitCasbinGraphQL, err := migRunner.NewSeed20201125001InitCasbinGraphQL(m.handler)
	if err != nil {
		return err
	}

	cfg, err := m.handler.GetDefaultConfig()
	if err != nil {
		return err
	}
	if err := m.migrator.RunSeeds(m.handler, cfg.Databases.IdentityDB.ConnectionName,
		seed20201118002InitCasbin,
		seed20201125001InitCasbinGraphQL); err != nil {
		return err
	}
	return nil
}

func (m *RDBMSMigration) _rollBackIdentitySeeds() error {
	seed20201118002InitCasbin, err := migRunner.NewSeed20201118002InitCasbin(m.handler)
	if err != nil {
		return err
	}
	seed20201125001InitCasbinGraphQL, err := migRunner.NewSeed20201125001InitCasbinGraphQL(m.handler)
	if err != nil {
		return err
	}

	cfg, err := m.handler.GetDefaultConfig()
	if err != nil {
		return err
	}
	if err := m.migrator.RollBackSeeds(m.handler, cfg.Databases.IdentityDB.ConnectionName,
		seed20201118002InitCasbin,
		seed20201125001InitCasbinGraphQL); err != nil {
		return err
	}
	return nil
}
