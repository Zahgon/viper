package viper

import (
	"io"
)

var SupportedRemoteProviders = []string{"etcd", "etcd3", "consul", "firestore", "nats"}

func resetRemote() { _ = "STUB: not implemented"; return }

type remoteConfigFactory interface {
	Get(rp RemoteProvider) (io.Reader, error)
	Watch(rp RemoteProvider) (io.Reader, error)
	WatchChannel(rp RemoteProvider) (<-chan *RemoteResponse, chan bool)
}

type RemoteResponse struct {
	Value []byte
	Error error
}

var RemoteConfig remoteConfigFactory

type UnsupportedRemoteProviderError string

func (str UnsupportedRemoteProviderError) Error() string { _ = "STUB: not implemented"; return "" }

type RemoteConfigError string

func (rce RemoteConfigError) Error() string { _ = "STUB: not implemented"; return "" }

type defaultRemoteProvider struct {
	provider      string
	endpoint      string
	path          string
	secretKeyring string
}

func (rp defaultRemoteProvider) Provider() string { _ = "STUB: not implemented"; return "" }

func (rp defaultRemoteProvider) Endpoint() string { _ = "STUB: not implemented"; return "" }

func (rp defaultRemoteProvider) Path() string { _ = "STUB: not implemented"; return "" }

func (rp defaultRemoteProvider) SecretKeyring() string { _ = "STUB: not implemented"; return "" }

type RemoteProvider interface {
	Provider() string
	Endpoint() string
	Path() string
	SecretKeyring() string
}

func AddRemoteProvider(provider, endpoint, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) AddRemoteProvider(provider, endpoint, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) providerPathExists(p *defaultRemoteProvider) bool {
	_ = "STUB: not implemented"
	return false
}

func ReadRemoteConfig() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) ReadRemoteConfig() error { _ = "STUB: not implemented"; return nil }

func WatchRemoteConfig() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) WatchRemoteConfig() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) WatchRemoteConfigOnChannel() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) getKeyValueConfig() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) getRemoteConfig(provider RemoteProvider) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Viper) watchKeyValueConfigOnChannel() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) watchKeyValueConfig() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) watchRemoteConfig(provider RemoteProvider) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
