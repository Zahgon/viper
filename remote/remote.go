package remote

import (
	"io"

	crypt "github.com/sagikazarmark/crypt/config"

	"github.com/spf13/viper"
)

type remoteConfigProvider struct{}

func (rc remoteConfigProvider) Get(rp viper.RemoteProvider) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (rc remoteConfigProvider) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (rc remoteConfigProvider) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getConfigManager(rp viper.RemoteProvider) (crypt.ConfigManager, error) {
	_ = "STUB: not implemented"
	return *new(crypt.ConfigManager), nil
}

func init() {
	viper.RemoteConfig = &remoteConfigProvider{}
}
