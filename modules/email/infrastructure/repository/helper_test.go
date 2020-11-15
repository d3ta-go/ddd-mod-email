package repository

import (
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
	return c, v, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.DefaultIdentity, identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}
	i.Claims.Username = "test.d3tago"
	i.RequestInfo.Host = "127.0.0.1:2020"

	return i
}
