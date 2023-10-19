package publicProxy

import (
	"context"
	"fmt"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	zhttp "github.com/zitadel/oidc/v2/pkg/http"
	"strings"
)

const V = 2

type Config struct {
	V         int
	Identity  string
	Address   string
	HostMatch string
	Oauth     *OauthConfig
}

type OauthConfig struct {
	RedirectHost     string
	RedirectPort     int
	RedirectHttpOnly bool
	HashKey          string `cf:"+secret"`
	Providers        []*OauthProviderConfig
}

func (oc *OauthConfig) GetProvider(name string) *OauthProviderConfig {
	for _, provider := range oc.Providers {
		if provider.Name == name {
			return provider
		}
	}
	return nil
}

type OauthProviderConfig struct {
	Name         string
	ClientId     string
	ClientSecret string `cf:"+secret"`
}

func DefaultConfig() *Config {
	return &Config{
		Identity: "public",
		Address:  "0.0.0.0:8080",
	}
}

func (c *Config) Load(path string) error {
	if err := cf.BindYaml(c, path, cf.DefaultOptions()); err != nil {
		return errors.Wrapf(err, "error loading frontend config '%v'", path)
	}
	if c.V != V {
		return errors.Errorf("invalid configuration version '%d'; expected '%d'", c.V, V)
	}
	return nil
}

func configureOauthHandlers(ctx context.Context, cfg *Config, tls bool) error {
	if cfg.Oauth == nil {
		logrus.Info("no oauth configuration; skipping oauth handler startup")
		return nil
	}
	if err := configureGoogleOauth(cfg.Oauth, tls); err != nil {
		return err
	}
	if err := configureGithubOauth(cfg.Oauth, tls); err != nil {
		return err
	}
	zhttp.StartServer(ctx, fmt.Sprintf("%s:%d", strings.Split(cfg.Address, ":")[0], cfg.Oauth.RedirectPort))
	return nil
}