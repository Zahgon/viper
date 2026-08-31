package viper

import (
	"github.com/spf13/afero"
)

func WithFinder(f Finder) Option { _ = "STUB: not implemented"; return *new(Option) }

type Finder interface {
	Find(fsys afero.Fs) ([]string, error)
}

func Finders(finders ...Finder) Finder { _ = "STUB: not implemented"; return *new(Finder) }

type combinedFinder struct {
	finders []Finder
}

func (c *combinedFinder) Find(fsys afero.Fs) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
