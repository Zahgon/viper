package viper

import (
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/afero"
	"github.com/spf13/pflag"
)

var v *Viper

func init() {
	v = New()
}

type DecoderConfigOption func(*mapstructure.DecoderConfig)

func DecodeHook(hook mapstructure.DecodeHookFunc) DecoderConfigOption {
	_ = "STUB: not implemented"
	return *new(DecoderConfigOption)
}

type Viper struct {
	keyDelim string

	configPaths []string

	fs afero.Fs

	finder Finder

	remoteProviders []*defaultRemoteProvider

	configName        string
	configFile        string
	configType        string
	configPermissions os.FileMode
	envPrefix         string

	automaticEnvApplied bool
	envKeyReplacer      StringReplacer
	allowEmptyEnv       bool

	parents        []string
	config         map[string]any
	override       map[string]any
	defaults       map[string]any
	kvstore        map[string]any
	pflags         map[string]FlagValue
	env            map[string][]string
	aliases        map[string]string
	typeByDefValue bool

	onConfigChange func(fsnotify.Event)

	logger *slog.Logger

	encoderRegistry EncoderRegistry
	decoderRegistry DecoderRegistry

	decodeHook mapstructure.DecodeHookFunc

	experimentalFinder     bool
	experimentalBindStruct bool
}

func New() *Viper { _ = "STUB: not implemented"; return nil }

type Option interface {
	apply(v *Viper)
}

type optionFunc func(v *Viper)

func (fn optionFunc) apply(v *Viper) { _ = "STUB: not implemented"; return }

func KeyDelimiter(d string) Option { _ = "STUB: not implemented"; return *new(Option) }

type StringReplacer interface {
	Replace(s string) string
}

func EnvKeyReplacer(r StringReplacer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDecodeHook(h mapstructure.DecodeHookFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NewWithOptions(opts ...Option) *Viper { _ = "STUB: not implemented"; return nil }

func SetOptions(opts ...Option) { _ = "STUB: not implemented"; return }

func Reset() { _ = "STUB: not implemented"; return }

var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"}

func OnConfigChange(run func(in fsnotify.Event)) { _ = "STUB: not implemented"; return }

func (v *Viper) OnConfigChange(run func(in fsnotify.Event)) { _ = "STUB: not implemented"; return }

func WatchConfig() { _ = "STUB: not implemented"; return }

func (v *Viper) WatchConfig() { _ = "STUB: not implemented"; return }

func SetConfigFile(in string) { _ = "STUB: not implemented"; return }

func (v *Viper) SetConfigFile(in string) { _ = "STUB: not implemented"; return }

func SetEnvPrefix(in string) { _ = "STUB: not implemented"; return }

func (v *Viper) SetEnvPrefix(in string) { _ = "STUB: not implemented"; return }

func GetEnvPrefix() string { _ = "STUB: not implemented"; return "" }

func (v *Viper) GetEnvPrefix() string { _ = "STUB: not implemented"; return "" }

func (v *Viper) mergeWithEnvPrefix(in string) string { _ = "STUB: not implemented"; return "" }

func AllowEmptyEnv(allowEmptyEnv bool) { _ = "STUB: not implemented"; return }

func (v *Viper) AllowEmptyEnv(allowEmptyEnv bool) { _ = "STUB: not implemented"; return }

func (v *Viper) getEnv(key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func ConfigFileUsed() string { _ = "STUB: not implemented"; return "" }

func (v *Viper) ConfigFileUsed() string { _ = "STUB: not implemented"; return "" }

func AddConfigPath(in string) { _ = "STUB: not implemented"; return }

func (v *Viper) AddConfigPath(in string) { _ = "STUB: not implemented"; return }

func (v *Viper) searchMap(source map[string]any, path []string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *Viper) searchIndexableWithPathPrefixes(source any, path []string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *Viper) searchSliceWithPathPrefixes(
	sourceSlice []any,
	prefixKey string,
	pathIndex int,
	path []string,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *Viper) searchMapWithPathPrefixes(
	sourceMap map[string]any,
	prefixKey string,
	pathIndex int,
	path []string,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *Viper) isPathShadowedInDeepMap(path []string, m map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

func (v *Viper) isPathShadowedInFlatMap(path []string, mi any) string {
	_ = "STUB: not implemented"
	return ""
}

func (v *Viper) isPathShadowedInAutoEnv(path []string) string { _ = "STUB: not implemented"; return "" }

func SetTypeByDefaultValue(enable bool) { _ = "STUB: not implemented"; return }

func (v *Viper) SetTypeByDefaultValue(enable bool) { _ = "STUB: not implemented"; return }

func GetViper() *Viper { _ = "STUB: not implemented"; return nil }

func Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

func (v *Viper) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

func Sub(key string) *Viper { _ = "STUB: not implemented"; return nil }

func (v *Viper) Sub(key string) *Viper { _ = "STUB: not implemented"; return nil }

func GetString(key string) string { _ = "STUB: not implemented"; return "" }

func (v *Viper) GetString(key string) string { _ = "STUB: not implemented"; return "" }

func GetBool(key string) bool { _ = "STUB: not implemented"; return false }

func (v *Viper) GetBool(key string) bool { _ = "STUB: not implemented"; return false }

func GetInt(key string) int { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetInt(key string) int { _ = "STUB: not implemented"; return 0 }

func GetInt32(key string) int32 { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetInt32(key string) int32 { _ = "STUB: not implemented"; return 0 }

func GetInt64(key string) int64 { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetInt64(key string) int64 { _ = "STUB: not implemented"; return 0 }

func GetUint8(key string) uint8 { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetUint8(key string) uint8 { _ = "STUB: not implemented"; return 0 }

func GetUint(key string) uint { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetUint(key string) uint { _ = "STUB: not implemented"; return 0 }

func GetUint16(key string) uint16 { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetUint16(key string) uint16 { _ = "STUB: not implemented"; return 0 }

func GetUint32(key string) uint32 { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetUint32(key string) uint32 { _ = "STUB: not implemented"; return 0 }

func GetUint64(key string) uint64 { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetUint64(key string) uint64 { _ = "STUB: not implemented"; return 0 }

func GetFloat64(key string) float64 { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetFloat64(key string) float64 { _ = "STUB: not implemented"; return 0 }

func GetTime(key string) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (v *Viper) GetTime(key string) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func GetDuration(key string) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (v *Viper) GetDuration(key string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func GetIntSlice(key string) []int { _ = "STUB: not implemented"; return nil }

func (v *Viper) GetIntSlice(key string) []int { _ = "STUB: not implemented"; return nil }

func GetStringSlice(key string) []string { _ = "STUB: not implemented"; return nil }

func (v *Viper) GetStringSlice(key string) []string { _ = "STUB: not implemented"; return nil }

func GetStringMap(key string) map[string]any { _ = "STUB: not implemented"; return nil }

func (v *Viper) GetStringMap(key string) map[string]any { _ = "STUB: not implemented"; return nil }

func GetStringMapString(key string) map[string]string { _ = "STUB: not implemented"; return nil }

func (v *Viper) GetStringMapString(key string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func GetStringMapStringSlice(key string) map[string][]string { _ = "STUB: not implemented"; return nil }

func (v *Viper) GetStringMapStringSlice(key string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func GetSizeInBytes(key string) uint { _ = "STUB: not implemented"; return 0 }

func (v *Viper) GetSizeInBytes(key string) uint { _ = "STUB: not implemented"; return 0 }

func UnmarshalKey(key string, rawVal any, opts ...DecoderConfigOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) UnmarshalKey(key string, rawVal any, opts ...DecoderConfigOption) error {
	_ = "STUB: not implemented"
	return nil
}

func Unmarshal(rawVal any, opts ...DecoderConfigOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) Unmarshal(rawVal any, opts ...DecoderConfigOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) decodeStructKeys(input any, opts ...DecoderConfigOption) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Viper) defaultDecoderConfig(output any, opts ...DecoderConfigOption) *mapstructure.DecoderConfig {
	_ = "STUB: not implemented"
	return nil
}

func stringToWeakSliceHookFunc(sep string) mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}

func decode(input any, config *mapstructure.DecoderConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func UnmarshalExact(rawVal any, opts ...DecoderConfigOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) UnmarshalExact(rawVal any, opts ...DecoderConfigOption) error {
	_ = "STUB: not implemented"
	return nil
}

func BindPFlags(flags *pflag.FlagSet) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) BindPFlags(flags *pflag.FlagSet) error { _ = "STUB: not implemented"; return nil }

func BindPFlag(key string, flag *pflag.Flag) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) BindPFlag(key string, flag *pflag.Flag) error {
	_ = "STUB: not implemented"
	return nil
}

func BindFlagValues(flags FlagValueSet) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) BindFlagValues(flags FlagValueSet) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func BindFlagValue(key string, flag FlagValue) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) BindFlagValue(key string, flag FlagValue) error {
	_ = "STUB: not implemented"
	return nil
}

func BindEnv(input ...string) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) BindEnv(input ...string) error { _ = "STUB: not implemented"; return nil }

func MustBindEnv(input ...string) { _ = "STUB: not implemented"; return }

func (v *Viper) MustBindEnv(input ...string) { _ = "STUB: not implemented"; return }

func (v *Viper) find(lcaseKey string, flagDefault bool) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func readAsCSV(val string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func stringToStringConv(val string) any { _ = "STUB: not implemented"; return *new(any) }

func stringToIntConv(val string) any { _ = "STUB: not implemented"; return *new(any) }

func IsSet(key string) bool { _ = "STUB: not implemented"; return false }

func (v *Viper) IsSet(key string) bool { _ = "STUB: not implemented"; return false }

func AutomaticEnv() { _ = "STUB: not implemented"; return }

func (v *Viper) AutomaticEnv() { _ = "STUB: not implemented"; return }

func SetEnvKeyReplacer(r *strings.Replacer) { _ = "STUB: not implemented"; return }

func (v *Viper) SetEnvKeyReplacer(r *strings.Replacer) { _ = "STUB: not implemented"; return }

func RegisterAlias(alias, key string) { _ = "STUB: not implemented"; return }

func (v *Viper) RegisterAlias(alias, key string) { _ = "STUB: not implemented"; return }

func (v *Viper) registerAlias(alias, key string) { _ = "STUB: not implemented"; return }

func (v *Viper) realKey(key string) string { _ = "STUB: not implemented"; return "" }

func InConfig(key string) bool { _ = "STUB: not implemented"; return false }

func (v *Viper) InConfig(key string) bool { _ = "STUB: not implemented"; return false }

func SetDefault(key string, value any) { _ = "STUB: not implemented"; return }

func (v *Viper) SetDefault(key string, value any) { _ = "STUB: not implemented"; return }

func Set(key string, value any) { _ = "STUB: not implemented"; return }

func (v *Viper) Set(key string, value any) { _ = "STUB: not implemented"; return }

func ReadInConfig() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) ReadInConfig() error { _ = "STUB: not implemented"; return nil }

func MergeInConfig() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) MergeInConfig() error { _ = "STUB: not implemented"; return nil }

func ReadConfig(in io.Reader) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) ReadConfig(in io.Reader) error { _ = "STUB: not implemented"; return nil }

func MergeConfig(in io.Reader) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) MergeConfig(in io.Reader) error { _ = "STUB: not implemented"; return nil }

func MergeConfigMap(cfg map[string]any) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) MergeConfigMap(cfg map[string]any) error { _ = "STUB: not implemented"; return nil }

func WriteConfig() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) WriteConfig() error { _ = "STUB: not implemented"; return nil }

func SafeWriteConfig() error { _ = "STUB: not implemented"; return nil }

func (v *Viper) SafeWriteConfig() error { _ = "STUB: not implemented"; return nil }

func WriteConfigAs(filename string) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) WriteConfigAs(filename string) error { _ = "STUB: not implemented"; return nil }

func WriteConfigTo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) WriteConfigTo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func SafeWriteConfigAs(filename string) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) SafeWriteConfigAs(filename string) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) writeConfig(filename string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) unmarshalReader(in io.Reader, c map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) marshalWriter(w io.Writer, configType string) error {
	_ = "STUB: not implemented"
	return nil
}

func keyExists(k string, m map[string]any) string { _ = "STUB: not implemented"; return "" }

func castToMapStringInterface(
	src map[any]any,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func castMapStringSliceToMapInterface(src map[string][]string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func castMapStringToMapInterface(src map[string]string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func castMapFlagToMapInterface(src map[string]FlagValue) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func mergeMaps(src, tgt map[string]any, itgt map[any]any) { _ = "STUB: not implemented"; return }

func AllKeys() []string { _ = "STUB: not implemented"; return nil }

func (v *Viper) AllKeys() []string { _ = "STUB: not implemented"; return nil }

func (v *Viper) flattenAndMergeMap(shadow map[string]bool, m map[string]any, prefix string) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) mergeFlatMap(shadow map[string]bool, m map[string]any) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func AllSettings() map[string]any { _ = "STUB: not implemented"; return nil }

func (v *Viper) AllSettings() map[string]any { _ = "STUB: not implemented"; return nil }

func (v *Viper) getSettings(keys []string) map[string]any { _ = "STUB: not implemented"; return nil }

func SetFs(fs afero.Fs) { _ = "STUB: not implemented"; return }

func (v *Viper) SetFs(fs afero.Fs) { _ = "STUB: not implemented"; return }

func SetConfigName(in string) { _ = "STUB: not implemented"; return }

func (v *Viper) SetConfigName(in string) { _ = "STUB: not implemented"; return }

func SetConfigType(in string) { _ = "STUB: not implemented"; return }

func (v *Viper) SetConfigType(in string) { _ = "STUB: not implemented"; return }

func SetConfigPermissions(perm os.FileMode) { _ = "STUB: not implemented"; return }

func (v *Viper) SetConfigPermissions(perm os.FileMode) { _ = "STUB: not implemented"; return }

func (v *Viper) getConfigType() string { _ = "STUB: not implemented"; return "" }

func (v *Viper) getConfigFile() (string, error) { _ = "STUB: not implemented"; return "", nil }

func Debug() { _ = "STUB: not implemented"; return }

func DebugTo(w io.Writer) { _ = "STUB: not implemented"; return }

func (v *Viper) Debug() { _ = "STUB: not implemented"; return }

func (v *Viper) DebugTo(w io.Writer) { _ = "STUB: not implemented"; return }
