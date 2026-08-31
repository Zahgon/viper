package viper

import (
	"github.com/spf13/afero"
)

func ExperimentalFinder() Option { _ = "STUB: not implemented"; return *new(Option) }

func (v *Viper) findConfigFile() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (v *Viper) findConfigFileWithFinder(finder Finder) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (v *Viper) findConfigFileOld() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (v *Viper) searchInPath(in string) (filename string) { _ = "STUB: not implemented"; return "" }

func exists(fs afero.Fs, path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }
