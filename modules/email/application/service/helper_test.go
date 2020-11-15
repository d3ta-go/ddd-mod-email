package service

import (
	"fmt"
	"testing"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
	"github.com/spf13/viper"
)

func newConfig(t *testing.T) (*config.Config, *viper.Viper, error) {
	c, v, err := config.NewConfig("../../../../conf")
	if err != nil {
		return nil, nil, err
	}
	if !c.CanRunTest() {
		panic(fmt.Sprintf("Cannot Run Test on env `%s`, allowed: %v", c.Environment.Stage, c.Environment.RunTestEnvironment))
	}
	c.IAM.Casbin.ModelPath = "../../../../conf/casbin/casbin_rbac_rest_model.conf"
	return c, v, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.DefaultIdentity, identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}
	cfg, err := h.GetDefaultConfig()
	if err != nil {
		t.Errorf("GetConfig: %s", err.Error())
	}
	if err := i.SetCasbinEnforcer(cfg.IAM.Casbin.ModelPath); err != nil {
		t.Errorf("SetCasbinEnforcer: %s", err.Error())
	}

	i.Claims.Username = "test.d3tago"
	i.Claims.AuthorityID = "group:admin"
	i.RequestInfo.Host = "127.0.0.1:2020"

	return i
}
